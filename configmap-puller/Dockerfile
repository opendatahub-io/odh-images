FROM golang:1.16 as builder

WORKDIR /src

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd/ ./cmd

RUN CGO_ENABLED=0 go build -o /bin/configmap-puller ./cmd/configmap-puller

# actual image:
FROM alpine:latest
WORKDIR /app
COPY --from=builder /bin/configmap-puller /bin/configmap-puller
ENTRYPOINT [ "/bin/configmap-puller" ]
