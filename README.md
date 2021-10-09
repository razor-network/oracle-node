[![Coverage Status](https://coveralls.io/repos/github/razor-network/razor-go/badge.svg?branch=main)](https://coveralls.io/github/razor-network/razor-go?branch=main)

# Razor-Go

Official node for running stakers in Golang.

## Installation

### Docker quick start
One of the quickest ways to get `razor-go` up and running on your machine is by using Docker:

`docker run -it razornetwork/razor-go /bin/bash`

### Prerequisites
* Golang 1.15 or later must be installed.
* Latest stable version of node is required.
* Silicon chip based Mac users must go for node 15.3.0+
* `geth` and `abigen` should be installed. (Skip this step if you don't want to fetch the bindings and build from scratch)

### Building the source
1. Run `npm install` to install the node dependencies.
2. If you want to build from scratch i.e., by fetching the smart contract bindings as well, run `npm run build-all`.

   _Note: To build from scratch, `geth` and `abigen` must be installed in your system._
3. If you already have the `pkg/bindings` you can run `npm run build` instead of `npm run build-all` to directly build the binary. 
4. If you want to build the binary without wanting to set the configurations use `npm run dockerize-build`
5. While building the binary, supply the provider RPC url and the gas multiplier.
6. The binary will be generated at `build/bin`.

## Commands

Go to the `build/bin` directory where the razor binary is generated.

`cd build/bin`

### Create Account
Create an account using the `create` command. You'll be asked to enter a password that'll be used to encrypt the keystore file.

```
$ ./razor create
```

Example:

```
$ ./razor create
Password: 
```

### Import Account
If you already have an account created, and have it's private key, that account can be imported into the `razor-go` client.
To do that, you can use the `import` command. You'll be asked the private key first and then the password which you want to encrypt your keystore file with.

```
$ ./razor import
```

Example:

```
$ ./razor import
🔑 Private Key:
Password: 
```

__Before staking on Razor Network, please ensure your account has eth and RAZOR. For testnet RAZOR, please use the faucet here - https://razorscan.io/dashboard/faucet__

### Stake

If you have a minimum of 1000 razors in your account, you can stake those using the stake command.
```
$ ./razor stake --address <address> --value <value>
```

Example:
```
$ ./razor stake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 1000
```

### Set Delegation

If you are a staker you can accept delegation from delegators and charge a commission from them.
```
$ ./razor setDelegation --address <address> --status <true_or_false> --commission <commission>
```

Example:
```
$ ./razor setDelegation --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --status true --commission 100
```

### Delegate

If you want to become a delegator use the `delegate` command. The staker whose `staker_id` is provided, their stake is increased.
```
$ ./razor delegate --address <address> --value <value> --stakerId <staker_id>
```

Example:
```
$ ./razor delegate --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 1000 --stakerId 1
```

### Vote
You can start voting once you've staked some razors
```
$ ./razor vote --address <address>
```

Example:
```
$ ./razor vote --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c
```

If you want to report incorrect values, there is a `rogue` mode available. Just pass an extra flag `--rogue` to start voting in rogue mode and the client will report wrong medians.

Example:

```
$ ./razor vote --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --rogue 
```

### Unstake
If you wish to withdraw your funds, you can run the `unstake` command followed by the `withdraw` command.
```
$ ./razor unstake --address <address> --stakerId <staker_id> --value <value> --autoWithdraw
```

Example:
```
$ ./razor unstake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --stakerId 1 --amount 1000 --autoWithdraw
```

### Withdraw
Once `unstake` has been called, you can withdraw your funds using the `withdraw` command

```
$ ./razor withdraw --address <address> --stakerId <staker_id>
```

Example:

```
$ ./razor withdraw --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --stakerId 1
```

### Reset Lock
If the withdrawal period is over, then the lock must be reset otherwise the user cannot unstake.
```
$ ./razor resetLock --address <address> --stakerId <staker_id>
```

Example:

```
$ ./razor resetLock --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --stakerId 1
```

### Transfer
Transfers razor to other accounts.

```
$ ./razor transfer --value <value> --to <transfer_to_address> --from <transfer_from_address>
```

Example:
```
$ ./razor transfer --value 100 --to 0x91b1E6488307450f4c0442a1c35Bc314A505293e --from 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c
```

### Set Config
There are a set of parameters that are configurable. These include:

* Provider: The RPC URL of the provider you are using to connect to the blockchain.
* Gas Multiplier: The value with which the gas price will be multiplied while sending every transaction.
* Buffer Size: Buffer size determines, out of all blocks in a state, in how many blocks the voting or any other operation can be performed.
* Wait Time: This is the number of blocks the system will wait while voting.
* Gas Price: The value of gas price if you want to set manually. If you don't provide any value or simply keep it to 0, the razor client will automatically calculate the optimum gas price and send it.
* Log Level: Normally debug logs are not logged into the log file. But if you want you can set `logLevel` to `debug` and fetch the debug logs.

The config is set while the build is generated, but if you need to change any of the above parameter, you can use the `setConfig` command.

```
$ ./razor setConfig --provider <rpc_provider> --gasmultiplier <multiplier_value> --buffer <buffer_percentage> --wait <wait_for_n_blocks> --gasprice <gas_price> --logLevel <debug_or_info>
```

Example:
```
$ ./razor setConfig --provider https://infura/v3/matic --gasmultiplier 1.5 --buffer 20 --wait 70 --gasprice 1 --logLevel debug
```

Other than setting these parameters in the config, you can use different values of these parameters in different command. Just add the same flag to any command you want to use and the new config changes will appear for that command.

Example:
```
$ ./razor vote --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --gasprice 10 
```
This will cause this particular vote command to run with a gas price of 10.


### Create Job
Create new jobs using `creteJob` command. 

_Note: This command is restricted to "Admin Role"_
```
$ ./razor createJob --url <URL> --selector <selector_in_json_selector_format> --name <name> --address <address> --power <power>
```

Example:
```
$ ./razor createJob --url https://www.alphavantage.co/query\?function\=GLOBAL_QUOTE\&symbol\=MSFT\&apikey\=demo --selector '[`Global Quote`][`05. price`]" --name msft --power 2 --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c
```
OR
```
$  ./razor createJob --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c -n ethusd -power 2 -s last -u https://api.gemini.com/v1/pubticker/ethusd
```

### Create Collection
Create new collections using `creteCollection` command.

_Note: This command is restricted to "Admin Role"_

```
$ ./razor createCollection --name <collection_name> --address <address> --jobIds <list_of_job_ids> --aggregation <aggregation_method> -power <power>
```

Example:
```
$ ./razor createCollection --name btcCollectionMean --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --jobIds 1,2 --aggregation 2 --power 2
```

### Add Job to Collection
Add existing jobs to existing collections using `addJobToCollection` command.

_Note: This command is restricted to "Admin Role"_

```
$ ./razor addJobToCollection --address <address> --jobId <job_id> --collectionId <collection_id>
```

Example:
```
$ ./razor addJobToCollection --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --collectionId 6 --jobId 7
```

### Remove Job from Collection
Remove jobs from existing collections using `removeJobFromCollection` command.

_Note: This command is restricted to "Admin Role"_

```
$ ./razor removeJobFromCollection --address <address> --jobId <job_id> --collectionId <collection_id>
```

Example:
```
$ ./razor removeJobFromCollection --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --collectionId 6 --jobId 7
```

### Modify Asset Status
Modify the active status of an asset using the `modifyAssetStatus` command.

_Note: This command is restricted to "Admin Role"_

```
$ ./razor modifyAssetStatus --assetId <assetId> --address <address> --status <true_or_false>
```

Example:
```
$ ./razor modifyAssetStatus --assetId 1 --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --status false
```

### Update Collection
Update the collection using `updateCollection` command.

_Note: This command is restricted to "Admin Role"_

```
$ ./razor updateCollection --collectionId <collection_id> --address <address> --aggregation <aggregation_method> --power <power> 
```

Example:
```
$ ./razor updateCollection -a 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --collectionId 3 --aggregation 2 --power 4```
```

### Update Job
Update the existing parameters of the Job using `updateJob` command.
  
_Note: This command is restricted to "Admin Role"_

```
./razor updateJob -address <address> --jobID <job_Id> -s <selector> -u <job_url>
```

Example:
```
$ ./razor updateJob -a 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --jobId 1 -s last -u https://api.gemini.com/v1/pubticker/btcusd 
```

### Contribute to razor-go 
We would really appreciate your contribution. To see our [contribution guideline](https://github.com/razor-network/razor-go/blob/main/.github/CONTRIBUTING.md)
