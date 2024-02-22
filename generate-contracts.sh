#!/bin/bash

# The file that will be updated
NETWORK_DETAILS_GO="core/network_details.go"

# Detect the OS and set the appropriate -i option for sed
if [ "$(uname)" = "Darwin" ]; then
    SED_I_OPTION=("-i" "")
else
    SED_I_OPTION=("-i")
fi

# Decide which addresses file to use based on the network specified as a parameter
network=$1
addresses_path=""

if [[ "$network" == "mainnet" ]]; then
  addresses_path="./network-details/mainnet.json"
elif [[ "$network" == "testnet" ]]; then
  addresses_path="./network-details/testnet.json"
else
  echo "Invalid network environment specified. Please use 'mainnet' or 'testnet'."
  exit 1
fi

echo "Updating contract addresses in $NETWORK_DETAILS_GO from $addresses_path"

# Function to update contract address in the Go file
update_address() {
  local json_key="$1"
  local go_var_name="$2"
  local address

  # Extract the address using jq
  address=$(jq -r ".addresses.$json_key" "$addresses_path")

  # Update the Go file using sed
  sed "${SED_I_OPTION[@]}" "s/var $go_var_name = \".*\"/var $go_var_name = \"$address\"/" "$NETWORK_DETAILS_GO"
}

# List of contracts to update
contract_addresses_list=(
  "StakeManager StakeManagerAddress"
  "RAZOR RAZORAddress"
  "CollectionManager CollectionManagerAddress"
  "VoteManager VoteManagerAddress"
  "BlockManager BlockManagerAddress"
)

# Loop through the list and update each contract address
for entry in "${contract_addresses_list[@]}"
do
  # Split the entry into an array containing the json key and the go variable name
  IFS=' ' read -ra ADDR <<< "$entry"
  update_address "${ADDR[0]}" "${ADDR[1]}"
done

echo "Contract addresses have been updated."