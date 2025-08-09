FROM golang:1.24-alpine

RUN apk update \
    && apk add --no-cache \
        git \
        make

RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.64.8 \
    && go install github.com/golang/mock/mockgen@v1.6 \
    && go install github.com/cweill/gotests/gotests@v1.6.0