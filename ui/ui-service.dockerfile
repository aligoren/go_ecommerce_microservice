# Golang

FROM golang:1.18-alpine as builder

RUN mkdir /app

COPY . /app

WORKDIR /app

RUN CGO_ENABLED=0 go build -o uiApp ./cmd/web

RUN chmod +x /app/uiApp

# Make a tiny docker image
FROM alpine:latest

RUN mkdir /app
RUN mkdir -p /app/templates

COPY --from=builder /app/uiApp /app
COPY ./services.production.json .
COPY ./.env.prod ./.env


ADD ./templates ./templates

COPY ./templates/ ./app/templates/

CMD ["/app/uiApp"]