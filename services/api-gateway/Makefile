include ../../.env
export
PG_URL := postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${DB}?sslmode=disable
export PG_URL

protoc:
	protoc internal/articles/pb/articles.proto --go-grpc_out=internal/articles/ --go_out=internal/articles/ && \
	protoc internal/auth/pb/auth.proto --go-grpc_out=internal/auth/ --go_out=internal/auth/
.PHONY: protoc
