#!/bin/bash

# Define the ChainId value based on the network argument
NETWORK=$1
CHAINID=""

if [[ "$NETWORK" == "mainnet" ]]; then
    CHAINID="0x109B4597"
elif [[ "$NETWORK" == "testnet" ]]; then
    CHAINID="0x5A79C44E"
else
    echo "Invalid network specified. Please choose 'mainnet' or 'testnet'."
    exit 1
fi

sed -i '' "s/var ChainId = big.NewInt(0x[a-fA-F0-9]*)/var ChainId = big.NewInt($CHAINID)/" core/constants.go

