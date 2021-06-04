GO ?= go
SHELL = sh
BIN_DIR = ./build/bin
RAZOR = ${BIN_DIR}/razor

all: generate_bindings install_razor set_config

generate_bindings:
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
	echo "Enter gas chain id: "; \
    read CHAIN_ID; \
    echo "\n"; \
    echo "Setting initial config..."; \
    ${RAZOR} setconfig -p $${PROVIDER} -g $${GAS_MULTIPLIER} -c $${CHAIN_ID}
	@echo "Setup done"

# TODO: Check if node_modules need to be deleted
clean:
