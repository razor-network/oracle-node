#!/bin/bash

set -e -o pipefail

mkdir -p ./pkg/bindings

generate_binding() {
  contract=$(echo $1 | awk '{print $1}')
  contractType=$(echo $1 | awk '{print $2}')
  go_source=$(echo $1 | awk '{print $3}')
  echo "Generating binding for (${contract})"
  abigen --abi ./contract-abi/${contract}.json --pkg 'bindings' --type=${contractType} --out ./pkg/bindings/${go_source}
}

contracts=(
  "BlockManager blockManager.go"
  "CollectionManager collectionManager.go"
  "StakeManager stakeManager.go"
  "VoteManger voteManager.go"
  "BlockStorage blockStorage.go"
  "CollectionStorage collectionStorage.go"
  "StakeStorage stakeStorage.go"
  "VoteStorage voteStorage.go"
  "ACL acl.go"
  "Delegator delegator.go"
  "StakedToken stakedToken.go"
  "RAZOR RAZOR.go"
  "AccessControl accessControl.go"
  "ERC20 erc20.go"
  "ERC165 erc165.go"
)

commitId=$1
git clone https://github.com/razor-network/contracts.git
cd contracts
git checkout $1
npm i
npm run compile
npx hardhat export-abi
cp -r abi ../contract-abi
cd ..
rm -rf contracts

for c in "${contracts[@]}"
do
    generate_binding "$c"
done

bash generate-contracts.sh