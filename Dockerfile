FROM golang:1-alpine AS builder
WORKDIR /app

COPY main.go .
COPY go.mod .

RUN go build -o server

FROM alpine:latest
RUN apk --no-cache add ca-certificates nano bash curl wget iputils busybox-extras
WORKDIR /app
COPY --from=builder /app/server .
COPY static/ static/

ENV SERVER_IP=0.0.0.0

EXPOSE 80
CMD ["./server"]