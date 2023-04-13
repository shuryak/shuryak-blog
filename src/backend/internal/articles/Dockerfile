FROM golang:1.20.2-alpine3.17 AS builder

ENV SERVICE=articles

# Set Go env
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/$SERVICE

# Install dependencies
RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev protoc git
RUN apk add protobuf-dev

# Build Go binary
COPY . /go/src/$SERVICE/
RUN --mount=type=cache,mode=0755,target=/go/pkg/mod make init && go mod download
RUN --mount=type=cache,target=/root/.cache/go-build --mount=type=cache,mode=0755,target=/go/pkg/mod make proto tidy build-articles

# Deployment container
FROM scratch

ENV SERVICE=articles

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/$SERVICE/$SERVICE /$SERVICE
COPY --from=builder /go/src/$SERVICE/internal/$SERVICE/config/config.yml /
COPY --from=builder /go/src/$SERVICE/internal/$SERVICE/migrations /migrations
EXPOSE 50051
ENTRYPOINT ["/articles"]
CMD []
