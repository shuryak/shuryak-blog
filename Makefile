GOPATH:=$(shell go env GOPATH)

.PHONY: minikube-env
minikube-env:
	@eval $(minikube -p minikube docker-env)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install github.com/go-micro/generator/cmd/protoc-gen-micro@latest

.PHONY: proto
proto:
	@protoc --proto_path=. --micro_out=. --go_out=:. proto/user/*.proto
	@protoc --proto_path=. --micro_out=. --go_out=:. proto/articles/*.proto

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: update
update:
	@go get -u

.PHONY: swag-init
swag-init:
	@go install github.com/swaggo/swag/cmd/swag@latest

.PHONY: swagger
swagger:
	@swag init -o internal/api-gw/docs -g internal/api-gw/swagger/swagger.go

.PHONY: build-user
build-user:
	@go build -o user cmd/user/*.go

.PHONY: build-articles
build-articles:
	@go build -o articles cmd/articles/*.go

.PHONY: build-api-gw
build-api-gw:
	@go build -o api-gw cmd/api-gw/*.go

.PHONY: user-image
user-image:
	docker build -f internal/user/Dockerfile -t ${IMAGE} .

.PHONY: articles-image
articles-image:
	docker build -f internal/articles/Dockerfile -t ${IMAGE} .

.PHONY: api-gw-image
api-gw-image:
	docker build -f internal/api-gw/Dockerfile -t ${IMAGE} .

.PHONY: apply
apply:
	kubectl apply -f k8s/
	kubectl apply -f k8s/jaeger/
	kubectl apply -f k8s/nats/
	kubectl apply -f k8s/user/
	kubectl apply -f k8s/articles/
	kubectl apply -f k8s/dashboard/
	kubectl apply -f k8s/api-gw/

.PHONY: delete
delete:
	kubectl delete svc --all
	kubectl delete deploy --all
	kubectl delete pvc --all
	kubectl delete pv --all
	kubectl delete cm --all

.PHONY: api-gw-port-forward
api-gw-port-forward:
	kubectl port-forward svc/api-gw 8080:8080
