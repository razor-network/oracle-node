#!/bin/bash

echo "Copying addresses.json to core/contracts.go"
touch core/contracts.go

echo "package core" > core/contracts.go
printf "\n" >> core/contracts.go

generate_contract_address() {
  jsonFileKey=$(echo $1 | awk '{print $1}')
  jsonFileKey="."$jsonFileKey
  goContractKey=$(echo $1 | awk '{print $2}')
  varDeclaration="var $goContractKey ="
  contractAddress=$(cat addresses.json | jq $jsonFileKey)
  echo "$varDeclaration $contractAddress" >> core/contracts.go
}

contract_addresses_list=(
  "StakeManager StakeManagerAddress"
  "RAZOR RAZORAddress"
  "Parameters ParametersAddress"
  "AssetManager AssetManagerAddress"
  "VoteManager VoteManagerAddress"
  "BlockManager BlockManagerAddress"
)

for c in "${contract_addresses_list[@]}"
do
    generate_contract_address "$c"
done


