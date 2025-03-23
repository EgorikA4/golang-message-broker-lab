FROM golang:1.24-alpine3.21 AS builder

WORKDIR /build

COPY ./cmd ./cmd

COPY ./vendor ./vendor
COPY ./go.mod ./go.sum ./

COPY ./config ./config
COPY ./internal ./internal

RUN go build -o producer ./cmd/producer/main.go
RUN go build -o consumer ./cmd/consumer/main.go

FROM alpine:3.21

WORKDIR /service

COPY --from=builder /build/producer ./
COPY --from=builder /build/consumer ./
