# Build stage
FROM golang:1.24rc1-alpine3.21 AS builder
WORKDIR /app
COPY . .
RUN apk add --no-cache curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.12.2/migrate.linux-amd64.tar.gz | tar xvz
RUN go build -o main main.go

# Run stage
FROM alpine:3.21
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate.linux-amd64 ./migrate
COPY wait-for.sh .
COPY app.env .
COPY start.sh .
COPY db/migration ./migration

EXPOSE 8080
ENTRYPOINT ["/app/start.sh","/app/main"]
