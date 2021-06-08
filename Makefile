GO ?= go
SHELL = sh
BIN_DIR = ./build/bin
RAZOR = ${BIN_DIR}/razor

all: fetch_bindings install_razor set_config

fetch_bindings:
	@echo "Installing contract dependencies..."
	@${SHELL} generate-bindings.sh
	@echo "Contract bindings generated....\n"

install_razor:
	@echo "Installing razor node...."
	${GO} build -o ./build/bin/razor main.go
	@echo "Razor node installed. \n"

set_config:
	@echo "Enter provider: "; \
	read PROVIDER; \
	echo "Enter gas multiplier value: "; \
	read GAS_MULTIPLIER; \
    echo "\n"; \
    echo "Setting initial config..."; \
    ${RAZOR} setconfig -p $${PROVIDER} -g $${GAS_MULTIPLIER}
	@echo "Setup done"
