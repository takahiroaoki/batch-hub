FROM golang:1.24-alpine AS builder
RUN apk update && apk add --no-cache make
WORKDIR /build
COPY . .
RUN make build

FROM alpine:3.22.2
COPY --from=builder /build/batch-hub /usr/local/bin/batch-hub
ENTRYPOINT ["/usr/local/bin/batch-hub"]
CMD ["migrate", "up"]