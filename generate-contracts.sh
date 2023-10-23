#!/bin/bash
touch core/contracts.go

echo "package core" > core/contracts.go
printf "\n" >> core/contracts.go

network=$1
addresses_path=""

if [[ "$network" == "mainnet" ]]; then
  addresses_path="addresses/mainnet.json"
elif [[ "$network" == "testnet" ]]; then
  addresses_path="addresses/testnet.json"
else
  echo "Invalid network environment specified. Please use 'mainnet' or 'testnet'."
  exit 1
fi

echo "Copying $addresses_path to core/contracts.go"

generate_contract_address() {
  jsonFileKey=$(echo $1 | awk '{print $1}')
  jsonFileKey="."$jsonFileKey
  goContractKey=$(echo $1 | awk '{print $2}')
  varDeclaration="var $goContractKey ="
  contractAddress=$(cat $addresses_path | jq $jsonFileKey)
  echo "$varDeclaration $contractAddress" >> core/contracts.go
}

contract_addresses_list=(
  "StakeManager StakeManagerAddress"
  "RAZOR RAZORAddress"
  "CollectionManager CollectionManagerAddress"
  "VoteManager VoteManagerAddress"
  "BlockManager BlockManagerAddress"
)

for c in "${contract_addresses_list[@]}"
do
    generate_contract_address "$c"
done