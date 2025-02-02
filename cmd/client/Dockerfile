FROM golang:1.22-alpine3.20 as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o client ./cmd/client

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/client /app/client

COPY config.yaml config.yaml

ENTRYPOINT ["/app/client"]

CMD ["--addr", "server:52345"]