package main

import (
	"auth/config"
	"auth/internal"
	grpccontroller "auth/internal/delivery/grpc"
	"auth/internal/entity"
	"auth/pkg/postgres"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"time"
)

func unaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	// TODO: log
	fmt.Println("--> unary interceptor: ", info.FullMethod)
	return handler(ctx, req)
}

func streamInterceptor(
	srv interface{},
	stream grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	// TODO: log
	fmt.Println("--> unary interceptor: ", info.FullMethod)
	return handler(srv, stream)
}

func seedUsers(us internal.UserStore) {
	createUser(us, "mike", "password1", "user")
	createUser(us, "john", "password2", "admin")
}

func createUser(us internal.UserStore, username, password, role string) (*entity.User, error) {
	user, err := entity.NewUser(username, password, role)
	if err != nil {
		return nil, err
	}
	return us.Create(context.Background(), *user)
}

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		// TODO: log
		fmt.Println("config error")
	}

	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		// TODO: log
		fmt.Println("postgres error")
	}
	defer pg.Close()

	us := internal.NewPostgresUserStore(pg)
	seedUsers(us)
	uss := internal.NewPostgresUserSessionStore(pg)
	sm := internal.NewSessionManager(uss, 64, 7*24*time.Hour)

	jwt := internal.NewJWTManager("super-secret", 10*time.Minute)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
		grpc.StreamInterceptor(streamInterceptor),
	)

	list, err := net.Listen("tcp", ":"+cfg.GRPC.Port)
	if err != nil {
		// TODO: log
		fmt.Println("listen fatal error")
	}

	grpccontroller.NewAuthServer(s, us, jwt, sm)

	err = s.Serve(list)
	if err != nil {
		// TODO: log
		fmt.Println("serve fatal error")
	}
}
