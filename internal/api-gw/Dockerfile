FROM golang:1.20.2-alpine3.17 AS builder

ENV SERVICE=api-gw

# Set Go env
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/$SERVICE

# Install dependencies
RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev protoc git
RUN apk add protobuf-dev

# Download grpc_health_probe; TODO: delete?
RUN GRPC_HEALTH_PROBE_VERSION=v0.4.15 && \
wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
chmod +x /bin/grpc_health_probe

# Build Go binary
COPY . /go/src/$SERVICE/
RUN --mount=type=cache,mode=0755,target=/go/pkg/mod make init swag-init && go mod download
RUN --mount=type=cache,target=/root/.cache/go-build --mount=type=cache,mode=0755,target=/go/pkg/mod make proto swagger tidy build-api-gw

# Deployment container
FROM scratch

ENV SERVICE=api-gw

COPY --from=builder /etc/ssl/certs /etc/ssl/certs
# TODO: delete?
COPY --from=builder /bin/grpc_health_probe /bin/
COPY --from=builder /go/src/$SERVICE/$SERVICE /$SERVICE
COPY --from=builder /go/src/$SERVICE/internal/$SERVICE/config/config.yml /
EXPOSE 8080
ENTRYPOINT ["/api-gw"]
CMD []
