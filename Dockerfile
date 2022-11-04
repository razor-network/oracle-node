FROM golang:1.17-alpine AS go
FROM ethereum/client-go:alltools-v1.10.7 AS ethereum

FROM node:16.2.0-alpine AS builder

COPY --from=ethereum /usr/local/bin/abigen /usr/local/bin/
COPY --from=go /usr/local/go/ /usr/local/go/

## Attaching current dir to workdir
WORKDIR /app
COPY . /app

## Install and Cleanup

RUN PATH="/usr/local/go/bin:${PATH}" \
    && apk add --update --no-cache python3 && ln -sf python3 /usr/bin/python \
    && apk add --update make gcc musl musl-dev g++ libc-dev bash linux-headers \
    && apk add --no-cache jq \
    && npm install \
    && npm run build-noargs \
    && cp build/bin/razor /usr/local/bin/


FROM alpine:latest
RUN apk add --update bash 
COPY --from=builder /usr/local/bin/razor /usr/local/bin/
RUN set -x \
	&& adduser -u 82 -D -S  razor
USER razor
ENTRYPOINT [ "razor" ]