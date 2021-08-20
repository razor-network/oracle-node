FROM golang:1.16-alpine as ethereum

RUN apk add --no-cache gcc musl-dev linux-headers git \
    && git clone https://github.com/ethereum/go-ethereum/ \
    && cd go-ethereum \
    && go run build/ci.go install ./cmd/abigen \
    && cp build/bin/abigen /usr/local/bin/

FROM golang:1.16-alpine AS go

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
    && npm install \
    && npm run dockerize-build \
    && cp build/bin/razor /usr/local/bin/


FROM alpine:latest
RUN apk add --update bash 
COPY --from=builder /usr/local/bin/razor /usr/local/bin/

