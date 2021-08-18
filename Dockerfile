## Use node 16.2 on alpine.
FROM node:16.2.0-alpine

## Install python
RUN apk add --update --no-cache python3 && ln -sf python3 /usr/bin/python

## Install other dependencies
RUN apk add --update make gcc musl musl-dev g++ libc-dev bash linux-headers

## Copy geth, abigen from geth image
COPY --from=ethereum/client-go:alltools-v1.10.7 /usr/local/bin/* /usr/local/bin/

## Copy golang from golang-alpine
COPY --from=golang:1.16-alpine /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"

RUN abigen --version

## Install solc using npm
RUN npm install -g solc

## Caching node_modules
ADD package.json /tmp/package.json
RUN cd /tmp && npm install
RUN mkdir -p /app && cp -a /tmp/node_modules /app

## Attaching current dir to workdir
WORKDIR /app
COPY . /app

## Create build
RUN npm run dockerize-build
