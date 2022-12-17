include ../../.env
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

certs:
	mkdir -p certs
	openssl genpkey -algorithm Ed25519 -out certs/ed25519key.pem
	openssl pkey -in certs/ed25519key.pem -pubout -out certs/ed25519key.pem.pub
.PHONY: certs

remove-certs:
	rm certs/ed25519key.pem
	rm certs/ed25519key.pem.pub
.PHONY: remove-certs
