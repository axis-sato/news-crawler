FROM golang:1.14.2-alpine

WORKDIR /go/app

RUN set -ex \
    && apk update \
    && apk add --no-cache alpine-sdk \
    && go get -v github.com/rubenv/sql-migrate/...
