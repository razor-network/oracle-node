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
  "Core/BlockManager.sol/BlockManager BlockManager blockManager.go"
  "Core/CollectionManager.sol/CollectionManager CollectionManager collectionManager.go"
  "Core/StakeManager.sol/StakeManager StakeManager stakeManager.go"
  "Core/VoteManager.sol/VoteManager VoteManger voteManager.go"
  "Core/storage/BlockStorage.sol/BlockStorage BlockStorage blockStorage.go"
  "Core/storage/CollectionStorage.sol/CollectionStorage CollectionStorage collectionStorage.go"
  "Core/storage/StakeStorage.sol/StakeStorage StakeStorage stakeStorage.go"
  "Core/storage/VoteStorage.sol/VoteStorage VoteStorage voteStorage.go"
  "Core/parameters/ACL.sol/ACL ACL acl.go"
  "Delegator.sol/Delegator Delegator delegator.go"
  "tokenization/StakedToken.sol/StakedToken StakedToken stakedToken.go"
  "tokenization/RAZOR.sol/RAZOR RAZOR RAZOR.go"
  "access/AccessControl.sol/AccessControl AccessControl accessControl.go"
  "token/ERC20/ERC20.sol/ERC20 ERC20 erc20.go"
  "utils/introspection/ERC165.sol/ERC165 ERC165 erc165.go"
)

commitId=$1
git clone https://github.com/razor-network/contracts.git
cd contracts
git checkout $1
npm i
npm run cp-ci-env
npm run compile
cp -r artifacts/contracts ../contract-abi
cp -r artifacts/@openzeppelin/contracts/* ../contract-abi
cd ..
rm -rf contracts

for c in "${contracts[@]}"
do
    generate_binding "$c"
done

bash generate-contracts.sh