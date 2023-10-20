GO ?= go
SHELL = /bin/bash
BIN_DIR = ./build/bin
RAZOR = ${BIN_DIR}/razor

all: update_chainId fetch_bindings install_razor set_config
build: install_razor set_config
build-noargs: update_chainId fetch_bindings install_razor
setup: update_chainId fetch_bindings

all-testnet: update_chainId_testnet fetch_bindings_testnet install_razor set_config
build-noargs-testnet: update_chainId_testnet fetch_bindings_testnet install_razor
setup-testnet: update_chainId_testnet fetch_bindings_testnet

fetch_bindings:
	@echo "Installing contract dependencies..."
	@echo ""
	@${SHELL} generate-bindings.sh mainnet
	@echo "Contract bindings generated...."
	@echo ""

fetch_bindings_testnet:
	@echo "Installing contract dependencies..."
	@echo ""
	@${SHELL} generate-bindings.sh testnet
	@echo "Contract bindings generated...."
	@echo ""

update_chainId:
	@echo "Update chainId..."
	@echo ""
	@${SHELL} update-chainId.sh mainnet
	@echo "ChainId updated to mainnet...."
	@echo ""

update_chainId_testnet:
	@echo "Update chainId..."
	@echo ""
	@${SHELL} update-chainId.sh testnet
	@echo "ChainId updated to testnet...."
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
