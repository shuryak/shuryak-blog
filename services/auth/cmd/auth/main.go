package main

import (
	"auth/config"
	"auth/internal"
	grpccontroller "auth/internal/delivery/grpc"
	"auth/internal/entity"
	"auth/pkg/logger"
	"auth/pkg/postgres"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

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

func accessibleRoles() map[string][]string {
	return nil
}

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	l := logger.New(cfg.Log.Level)

	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal("postgres error")
		return
	}
	defer pg.Close()

	us := internal.NewPostgresUserStore(pg)
	seedUsers(us)
	uss := internal.NewPostgresUserSessionStore(pg)
	sm := internal.NewSessionManager(uss, 64, 7*24*time.Hour)

	jwt := internal.NewJWTManager("super-secret", 10*time.Minute)

	interceptor := internal.NewAuthInterceptor(jwt, accessibleRoles(), l)

	s := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor.Unary()),
		grpc.StreamInterceptor(interceptor.Stream()),
	)

	list, err := net.Listen("tcp", ":"+cfg.GRPC.Port)
	if err != nil {
		l.Fatal("listen fatal error")
		return
	}

	grpccontroller.NewAuthServer(s, us, jwt, sm)

	l.Info("Start GRPC server at %s", list.Addr())

	err = s.Serve(list)
	if err != nil {
		l.Fatal("serve fatal error")
		return
	}
}
