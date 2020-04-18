FROM golang:1.14.2-alpine as build

WORKDIR /go/app

COPY . .

RUN set -ex \
    && go build -o crawler


FROM alpine

WORKDIR /app

COPY --from=build /go/app/crawler .

RUN addgroup go \
    && adduser -D -G go go \
    && chown -R go:go /app/crawler

CMD ["./crawler"]