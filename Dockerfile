FROM golang:alpine AS builder

WORKDIR /app

COPY . .

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o minusx ./cmd/server

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/minusx .

RUN chmod +x minusx

EXPOSE 8081

ENTRYPOINT ["./minusx"]