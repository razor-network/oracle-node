#!/bin/bash

set -e -o pipefail

mkdir -p ./pkg/bindings

network=$2

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
  "Delegator delegator.go"
  "ERC20 erc20.go"
  "ERC165 erc165.go"
  "CollectionManager collectionManager.go"
  "CollectionStorage collectionStorage.go"
  "RAZOR RAZOR.go"
  "StakeManager stakeManager.go"
  "StakeStorage stakeStorage.go"
  "StakedToken stakedToken.go"
  "VoteManager voteManager.go"
  "VoteStorage voteStorage.go"
)

for c in "${contracts[@]}"
do
    generate_binding "$c"
done

bash generate-contracts.sh $network