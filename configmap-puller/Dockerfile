FROM registry.access.redhat.com/ubi8/go-toolset:1.15.13-4 as builder

WORKDIR /src

COPY go.mod .
COPY go.sum .
COPY vendor/ ./vendor
COPY cmd/ ./cmd

RUN mkdir -p /tmp/bin \
    && CGO_ENABLED=0 go build -o /tmp/bin/configmap-puller ./cmd/configmap-puller \
    && chmod +x /tmp/bin/configmap-puller

# actual image:
FROM registry.access.redhat.com/ubi8/go-toolset:1.15.13-4
WORKDIR /app
COPY --from=builder /tmp/bin/configmap-puller /usr/local/bin/configmap-puller
ENTRYPOINT [ "/usr/local/bin/configmap-puller" ]
