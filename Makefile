include .env
export
PG_URL := postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${DB}?sslmode=disable
export PG_URL

debug:
	docker-compose up -d postgres db-migrations && go run cmd/app/main.go
.PHONY: debug
