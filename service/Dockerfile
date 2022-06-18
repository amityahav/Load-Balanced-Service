# builder
FROM golang:1.18-alpine AS builder

RUN set -ex &&\
    apk add --no-progress --no-cache \
      gcc \
      musl-dev \
      git
ENV GO111MODULE=on

WORKDIR /app
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o app .

# image
FROM alpine:edge
WORKDIR /
COPY --from=builder /app/app .
ENTRYPOINT ["/app"]
