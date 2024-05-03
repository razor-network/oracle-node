#!/bin/bash

# Define the ChainId value based on the network argument
NETWORK=$1
CHAINID=""

if [[ "$NETWORK" == "mainnet" ]]; then
    CHAINID="0x109B4597"
elif [[ "$NETWORK" == "testnet" ]]; then
    CHAINID="1444673419"
else
    echo "Invalid network specified. Please choose 'mainnet' or 'testnet'."
    exit 1
fi

# Detect the OS and set the appropriate -i option
if [ "$(uname)" = "Darwin" ]; then  # macOS
    SED_I_OPTION=("-i" "")
else  # GNU/Linux and others
    SED_I_OPTION=("-i")
fi

# Use the correct option for sed based on the detected OS
sed "${SED_I_OPTION[@]}" "s/var ChainId = big.NewInt(0x[a-fA-F0-9]*)/var ChainId = big.NewInt($CHAINID)/" core/constants.go
