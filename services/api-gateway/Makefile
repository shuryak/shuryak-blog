include ../../.env
export
PG_URL := postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${DB}?sslmode=disable
export PG_URL

gen:
	swag init -g internal/swagger/swagger.go
.PHONY: gen

swagger:
	xdg-open http://localhost:8080/swagger/index.html >/dev/null 2>&1
.PHONY: swagger

protoc:
	protoc internal/articles/pb/articles.proto --go-grpc_out=internal/articles/ --go_out=internal/articles/ && \
	protoc internal/auth/pb/auth.proto --go-grpc_out=internal/auth/ --go_out=internal/auth/
.PHONY: protoc
