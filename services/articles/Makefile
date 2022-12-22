include ../../.env
export
PG_URL := postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${DB}?sslmode=disable
export PG_URL

containers:
	swag init -g internal/controller/http/v1/router.go && docker-compose up postgres db-migrations
.PHONY: containers

run:
	swag init -g internal/controller/http/v1/router.go && docker-compose up -d postgres db-migrations && go run cmd/app/main.go
.PHONY: start

swagger:
	xdg-open http://localhost:8080/swagger/index.html >/dev/null 2>&1
.PHONY: swagger

protoc:
	protoc internal/controller/grpc/articles.proto --go-grpc_out=internal/controller/grpc/ --go_out=internal/controller/grpc/ && \
	protoc internal/controller/grpc/auth.proto --go-grpc_out=internal/controller/grpc/ --go_out=internal/controller/grpc/
.PHONY: protoc

evans:
	evans internal/controller/grpc/articles.proto -p 50051
.PHONY: evans
