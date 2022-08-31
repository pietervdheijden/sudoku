# syntax=docker/dockerfile:1

# Source: https://docs.docker.com/language/golang/build-images/

## Build
FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /sudoku-api

## Deploy
FROM gcr.io/distroless/base-debian10

# ENV GIN_MODE=release

WORKDIR /

COPY --from=build /sudoku-api /sudoku-api

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/sudoku-api"]