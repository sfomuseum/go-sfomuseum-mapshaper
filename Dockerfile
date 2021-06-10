FROM golang:1.16-alpine AS builder

RUN mkdir /build

COPY . /build/go-sfomuseum-mapshaper

RUN apk update && apk upgrade \
    && cd /build/go-sfomuseum-mapshaper \
    && go build -mod vendor -o /usr/local/bin/mapshaper-server cmd/server/main.go \
    && cd && rm -rf /build

FROM node:15-alpine

COPY --from=builder /usr/local/bin/mapshaper-server /usr/local/bin/mapshaper-server

RUN mkdir /usr/local/data
VOLUME /usr/local/data

RUN apk update && apk upgrade
RUN npm install -g mapshaper

ENTRYPOINT ["/usr/local/bin/mapshaper-server"]