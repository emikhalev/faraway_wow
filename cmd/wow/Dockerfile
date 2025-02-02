FROM golang:1.22-alpine3.20 as builder

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o wow ./cmd/wow

FROM alpine:latest

WORKDIR /app

EXPOSE 52345

COPY --from=builder /app/wow /app/wow

COPY config.yaml config.yaml

ENTRYPOINT ["/app/wow"]

CMD ["--config", "./config.yaml"]
