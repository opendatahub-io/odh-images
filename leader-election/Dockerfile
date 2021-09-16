FROM registry.access.redhat.com/ubi8/go-toolset:1.15.14-3 AS builder
WORKDIR /src

COPY go.mod .
COPY go.sum .
COPY main.go .
RUN mkdir -p /tmp/bin \
    && CGO_ENABLED=0 go build -o /tmp/bin/leader-election \
    && chmod +x /tmp/bin/leader-election

#base image
FROM registry.access.redhat.com/ubi8/go-toolset:1.15.13-4
WORKDIR /app
COPY --from=builder /tmp/bin/leader-election /usr/local/bin/leader-election
ENTRYPOINT ["/usr/local/bin/leader-election"]
