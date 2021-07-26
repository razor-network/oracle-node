GO ?= go
SHELL = /bin/bash
BIN_DIR = ./build/bin
RAZOR = ${BIN_DIR}/razor

all: fetch_bindings install_razor set_config
build: install_razor set_config
docker-build: fetch_bindings install_razor

fetch_bindings:
	@echo "Installing contract dependencies..."
	@${SHELL} generate-bindings.sh
	@echo "Contract bindings generated....\n"

install_razor:
	@echo "Installing razor node...."
	${GO} build -o ./build/bin/razor main.go
	@echo "Razor node installed."
	@echo ""

set_config:
	@echo "Setup initial config"
	@${SHELL} config.sh
	@echo ""
	@echo "Razor node is set up and ready to use"