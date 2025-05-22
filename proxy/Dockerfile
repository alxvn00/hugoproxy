FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go mod tidy
RUN go build -o proxy

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/proxy /proxy
EXPOSE 8080
ENTRYPOINT ["/proxy"]
