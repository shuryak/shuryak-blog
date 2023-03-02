FROM golang:1.20.1-alpine3.16 AS builder

ENV SERVICE=user

# Set Go env
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/$SERVICE

# Install dependencies
RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev protoc git
RUN apk add protobuf-dev

# Build Go binary
COPY . /go/src/$SERVICE/
RUN make init && go mod download
RUN make proto tidy build-user

# Deployment container
FROM scratch

ENV SERVICE=user

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/$SERVICE/$SERVICE /$SERVICE
COPY --from=builder /go/src/$SERVICE/internal/$SERVICE/config/config.yml /
COPY --from=builder /go/src/$SERVICE/internal/$SERVICE/migrations /migrations
EXPOSE 50051
ENTRYPOINT ["/user"]
CMD []
