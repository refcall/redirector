# syntax=docker/dockerfile:1

## Build
FROM golang:1.21-bullseye AS build

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