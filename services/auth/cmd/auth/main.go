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
	createUser(us, "shuryak", "password", "user")
	createUser(us, "admin", "password", "admin")
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
		log.Fatalf("Config error: %v", err)
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

	jwt, err := internal.NewJWTManager(cfg.Certs.PrivateKeyPEMPath, cfg.Certs.PublicKeyPEMPath, 10*time.Minute)
	if err != nil {
		log.Fatalf("Certs error: %v", err)
	}

	list, err := net.Listen("tcp", ":"+cfg.GRPC.Port)

	s := grpc.NewServer()
	grpccontroller.NewAuthGRPCServer(s, us, jwt, sm)

	if err != nil {
		l.Fatal("listen fatal error")
		return
	}

	l.Info("Start GRPC server at %v", list.Addr())

	err = s.Serve(list)
	if err != nil {
		l.Fatal("serve fatal error")
		return
	}
}
