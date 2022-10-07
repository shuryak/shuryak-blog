debug:
	docker-compose up -d postgres db-migrations && go run cmd/app/main.go
.PHONY: debug
