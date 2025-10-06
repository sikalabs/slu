FROM golang:1.24 AS build
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build

FROM debian:13-slim
COPY --from=build /build/slu /usr/local/bin/slu
