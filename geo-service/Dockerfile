# syntax=docker/dockerfile:1.4

FROM golang:1.19-alpine AS builder

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o geo ./cmd/server/main.go


FROM alpine:latest

WORKDIR /root

COPY --from=builder /app/geo .

CMD ["./geo"]