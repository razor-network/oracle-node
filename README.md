# Razor-Go

Official node for running stakers in Golang.

## Installation

### Prerequisites
* Golang 1.15 or later must be installed.
* Latest stable version of node is required.
* Silicon chip based Mac users must go for node 15.3.0+
* `geth` and `abigen` should be installed. (Skip this step if you don't want to fetch the bindings and build from scratch)

### Building the source
1. Run `npm install` to install the node dependencies.
2. Run `npm run build` to build the binary. While building the binary, supply the provider RPC url and the gas multiplier.
3. If you want to build from scratch i.e., by fetching the smart contract bindings as well, run `npm run build-all` instead of `npm run build`. 
   
   _Note: To build from scratch, `geth` and `abigen` must be installed in your system._
5. The binary will be generated at `build/bin`.

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

### Before staking on Razor Network, please ensure your account has eth and RAZOR. For RAZOR, please use the faucet here - https://razorscan.io/dashboard/faucet

### Stake

If you have a minimum of 1000 razors in your account, you can stake those using the stake command.
```
$ ./razor stake --address <address> --amount <amount>
```

Example:
```
$ ./razor stake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --amount 1000
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

### Unstake
If you wish to withdraw your funds, you can run the `unstake` command followed by the `withdraw` command.
```
$ ./razor unstake --address <address>
```

Example:
```
$ ./razor unstake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c
```

### Withdraw
Once `unstake` has been called, you can withdraw your funds using the `withdraw` command

```
$ ./razor withdraw --address <address>
```

Example:

```
$ ./razor withdraw --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c
```
### Transfer
Transfers razor to other accounts.

```
$ ./razor transfer --amount <amount> --to <transfer_to_address> --from <transfer_from_address>
```

Example:
```
$ ./razor transfer --amount 100 --to 0x91b1E6488307450f4c0442a1c35Bc314A505293e --from 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c
```

### Create Job
You can create new jobs using `creteJob` command.

```
$ ./razor createJob --url <URL> --selector <selector_comma_seperated> --name <name> --fee <fee_to_lock> --address <address>
```

Example:
```
$ ./razor createJob --url https://www.alphavantage.co/query\?function\=GLOBAL_QUOTE\&symbol\=MSFT\&apikey\=demo --selector "Global Quote,05. price" --fee 100 --name msft --repeat false --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c
```

### Create Collection
You can create new collections using `creteCollection` command.

```
$ ./razor createCollection --name <collection_name> --fee <fee_to_lock> --address <address> --jobIds <list_of_job_ids> --aggregation <aggregation_method>
```

Example:
```
$ ./razor createCollection --name btcCollectionMean -f 100 --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --jobIds 1,2 --aggregation 2
```

### Add Job to Collection
You can add existing jobs to existing collections using `addJobToCollection` command.

```
$ ./razor addJobToCollection --address <address> --jobId <job_id> --collectionId <collection_id>
```

Example:
```
$ ./razor addJobToCollection --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --collectionId 6 --jobId 7
```

### Set Config
The config is set while the build is generated, but if you need to change your provider or the gas multiplier, you can use the `setconfig` command.

```
$ ./razor setconfig --provider <rpc_provider> --gasmultiplier <multiplier_value>
```

Example:
```
$ ./razor setconfig --provider https://infura/v3/matic --gasmultiplier 1.5
```
