FROM golang:1.24.1-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o pdf-service

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/pdf-service .

COPY .env . 

EXPOSE 8081
CMD ["./pdf-service"]
