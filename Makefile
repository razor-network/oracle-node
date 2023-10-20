GO ?= go
SHELL = /bin/bash
BIN_DIR = ./build/bin
RAZOR = ${BIN_DIR}/razor

all: fetch_bindings install_razor set_config
build: install_razor set_config
build-noargs: fetch_bindings install_razor
setup: fetch_bindings

all-testnet: fetch_bindings_testnet install_razor set_config
build-noargs-testnet: fetch_bindings_testnet install_razor
setup-testnet: fetch_bindings_testnet

fetch_bindings:
    @echo "Installing contract dependencies..."
    @echo ""
    @${SHELL} generate-bindings.sh --environment mainnet
    @echo "Contract bindings generated...."
    @echo ""

fetch_bindings_testnet:
    @echo "Installing contract dependencies..."
    @echo ""
    @${SHELL} generate-bindings.sh --environment testnet
    @echo "Contract bindings generated...."
    @echo ""

install_razor:
    @echo "Installing razor node...."
    ${GO} build -ldflags "-s -w" -o ./build/bin/razor main.go
    @echo "Razor node installed."
    @echo ""

set_config:
    @echo "Setup initial config"
    @${SHELL} config.sh
    @echo ""
    @echo "Razor node is set up and ready to use"