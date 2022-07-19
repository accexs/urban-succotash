FROM golang:1.18-alpine AS base
WORKDIR /app

ENV GO111MODULE="on"
ENV GOOS="linux"
ENV CGO_ENABLED=1

# System dependencies
RUN apk update \
    && apk add --no-cache \
    ca-certificates \
    git \
    gcc \
    musl-dev \
    && update-ca-certificates

### Development with hot reload
FROM base AS dev
WORKDIR /app

# Hot reloading mod
RUN go install github.com/cosmtrek/air@latest
EXPOSE 8080

ENTRYPOINT ["air"]

### Executable builder
FROM base AS builder
WORKDIR /app

# Application dependencies
COPY . /app
RUN go mod download \
    && go mod verify

RUN go build -o me-wallet -a .