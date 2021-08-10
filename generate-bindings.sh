#!/bin/bash

set -e -o pipefail

mkdir -p ./pkg/bindings

echo "\n"

# TODO: Copy bindings directly from node_modules to ./pkg/bindings

generate_binding() {
  contract=$(echo $1 | awk '{print $1}')
  go_source=$(echo $1 | awk '{print $2}')
  echo "Generating binding for (${contract})"
  abigen --abi ./node_modules/@razor-network/contracts/abi/${contract}.json --pkg 'bindings' --type=${contract} --out ./pkg/bindings/${go_source}
}

contracts=(
  "AccessControl accessControl.go"
  "ACL acl.go"
  "BlockManager blockManager.go"
  "BlockStorage blockStorage.go"
  "Parameters parameters.go"
  "Delegator delegator.go"
  "ERC20 erc20.go"
  "ERC165 erc165.go"
  "AssetManager assetManager.go"
  "AssetStorage assetStorage.go"
  "Random random.go"
  "RAZOR RAZOR.go"
  "StakeManager stakeManager.go"
  "StakeStorage stakeStorage.go"
  "VoteManager voteManager.go"
  "VoteStorage voteStorage.go"
)

for c in "${contracts[@]}"
do
    generate_binding "$c"
done