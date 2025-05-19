FROM golang:alpine AS builder

WORKDIR /app
    
COPY go.mod go.sum ./
RUN go mod download
    
COPY . .
    
RUN CGO_ENABLED=0 GOOS=linux go build -o minusx ./cmd/server
    
FROM alpine:latest
    
WORKDIR /app
    
RUN apk --no-cache add curl
    
COPY --from=builder /app/minusx .
    
RUN chmod +x minusx
    
EXPOSE 8081
    
ENTRYPOINT ["./minusx"]