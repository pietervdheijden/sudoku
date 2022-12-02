# Source: https://docs.docker.com/language/golang/build-images/

# Build
FROM golang:1.19-alpine AS build
# CGO has to be disabled to use the distroless image in final layer (docs: https://stackoverflow.com/a/72727838/3737152).
ENV CGO_ENABLED=0
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . .
RUN go build -o /sudoku-api
RUN go test ./...

# Deploy
FROM gcr.io/distroless/static AS final
ENV GIN_MODE=release
WORKDIR /
EXPOSE 8080
USER nonroot:nonroot
COPY --from=build /sudoku-api /sudoku-api
ENTRYPOINT [ "/sudoku-api" ]
