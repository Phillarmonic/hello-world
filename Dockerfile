FROM golang:1-alpine AS builder
WORKDIR /app

COPY main.go .
COPY go.mod .

RUN go build -o server

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/server .
COPY static/ static/

ENV SERVER_IP=0.0.0.0

EXPOSE 8080
CMD ["./server"]