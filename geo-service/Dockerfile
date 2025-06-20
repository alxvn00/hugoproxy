# syntax=docker/dockerfile:1.4

FROM golang:1.21.5-alpine AS builder

ENV GOPATH=""

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

COPY .env .env
COPY internal/docs ./internal/docs

RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o geo ./cmd/server/main.go

RUN ls -al /app


FROM alpine:latest

WORKDIR /root

COPY --from=builder /app/geo .
COPY --from=builder /app/.env .env

CMD ["./geo"]