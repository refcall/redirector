# syntax=docker/dockerfile:1

## Build
FROM golang:1.21-bullseye AS build

# Authorize SSH Host
RUN mkdir -p /root/.ssh && \
    chmod 0700 /root/.ssh && \
    ssh-keyscan github.com > /root/.ssh/known_hosts

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /backend

## Deploy
FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /backend /backend

USER nonroot:nonroot

ENTRYPOINT ["/backend"]