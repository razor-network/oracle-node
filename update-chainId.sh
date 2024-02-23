#!/bin/bash

# Define the path to the network details Go file
NETWORK_DETAILS_GO="core/network_details.go"

# Define the ChainId and ChainName values based on the network argument
NETWORK=$1
CHAIN_ID=""
CHAIN_NAME=""

# Define the path to the JSON file based on the chosen network
JSON_FILE="./network_details/${NETWORK}.json"

# Check if the JSON file exists
if [[ ! -f "$JSON_FILE" ]]; then
    echo "JSON file for ${NETWORK} does not exist."
    exit 1
fi

# Use jq to parse the JSON file and extract the chain_id and chain_name values
CHAIN_ID=$(jq -r '.chain_id' "$JSON_FILE")
CHAIN_NAME=$(jq -r '.chain_name' "$JSON_FILE")

# Check if the CHAIN_ID or CHAIN_NAME is empty
if [[ -z "$CHAIN_ID" ]] || [[ -z "$CHAIN_NAME" ]]; then
    echo "Chain ID or Chain Name could not be found in the network_details JSON file."
    exit 1
fi

# Detect the OS and set the appropriate -i option for sed
if [ "$(uname)" = "Darwin" ]; then  # macOS
    SED_I_OPTION=("-i" "")
else  # GNU/Linux and others
    SED_I_OPTION=("-i")
fi

# Update ChainId in the Go file
sed "${SED_I_OPTION[@]}" "s/var ChainId = big.NewInt([^)]*)/var ChainId = big.NewInt($CHAIN_ID)/" "$NETWORK_DETAILS_GO"

# Update ChainName in the Go file
sed "${SED_I_OPTION[@]}" "s/var ChainName = \".*\"/var ChainName = \"$CHAIN_NAME\"/" "$NETWORK_DETAILS_GO"

echo "Chain ID and Chain Name have been updated in $NETWORK_DETAILS_GO."
