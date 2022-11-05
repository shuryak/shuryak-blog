include .env
export
PG_URL := postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${DB}?sslmode=disable
export PG_URL

#run:
#	docker-compose up -d postgres db-migrations && go run cmd/auth/main.go
#.PHONY: start

protoc:
	protoc internal/delivery/grpc/auth.proto --go-grpc_out=internal/delivery/grpc/ --go_out=internal/delivery/grpc/
.PHONY: protoc

evans:
	evans internal/delivery/grpc/auth.proto -p 50052
.PHONY: evans
