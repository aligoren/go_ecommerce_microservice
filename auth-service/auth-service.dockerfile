# Golang

FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o authApp ./cmd/api

RUN chmod +x /app/authApp

# Make a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/authApp /app
COPY ./.env.prod ./.env
CMD ["/app/authApp"]