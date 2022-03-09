[![Coverage Status](https://coveralls.io/repos/github/razor-network/razor-go/badge.svg?branch=main)](https://coveralls.io/github/razor-network/razor-go?branch=main)

# Razor-Go

Official node for running stakers in Golang.

## Installation

### Docker quick start

One of the quickest ways to get `razor-go` up and running on your machine is by using Docker:

```
  docker run -d \
  -it \
  --name razor-go \
  -v "$(echo $HOME)"/.razor:/root/.razor \
  razornetwork/razor-go
```

Note that we are leveraging docker bind-mounts to mount `.razor` directory so that we have a shared mount of `.razor` directory between the host and the container. The `.razor` directory holds keys to the addresses that we use in `razor-go`, along with logs and config. We do this to persist data in the host machine, otherwise you would lose your keys once you delete the container.

You need to set a provider before you can operate razor-go cli on docker:

```
docker exec -it razor-go razor setConfig -p <provider_url>
```

You can now execute razor-go cli commands by running:

```
docker exec -it razor-go razor <command>
```

### Setting up dev environment with docker-compose

You can build razor-go docker image by running:

```
docker-compose build
```

> **_NOTE:_** Add platform: linux/x86_64 for Silicon based MAC in docker-compose.yml.

Run razor-go locally with:

```
docker-compose up -d
```

You can interact with razor:

```
docker exec -it razor-go razor ...
```

### Prerequisites

- Golang 1.15 or later must be installed.
- Latest stable version of node is required.
- Silicon chip based Mac users must go for node 15.3.0+
- `geth` and `abigen` should be installed. (Skip this step if you don't want to fetch the bindings and build from scratch)
- `solc` and `jq` must be installed.

### Building the source

1. Run `npm install` to install the node dependencies.
2. If you want to build from scratch i.e., by fetching the smart contract bindings as well, run `npm run build-all`.

   _Note: To build from scratch, `geth` and `abigen` must be installed in your system._

3. If you already have the `pkg/bindings` you can run `npm run build` instead of `npm run build-all` to directly build the binary.
4. If you want to build the binary without wanting to set the configurations use `npm run build-noargs`
5. While building the binary, supply the provider RPC url and the gas multiplier.
6. To bypass the intractive mode of providing password, create file in `.razor` directory with providing password in it.
7. The binary will be generated at `build/bin`.

## Commands

Go to the `build/bin` directory where the razor binary is generated.

`cd build/bin`

### Create Account

Create an account using the `create` command. You'll be asked to enter a password that'll be used to encrypt the keystore file.

razor cli

```
$ ./razor create

```

Docker

```
docker run -it  \
    -v "$(echo $HOME)"/.razor:/root/.razor \
    razornetwork/razor-go:latest razor create
```

Docker providing password file

```
docker run -it  \
    -v "$(echo $HOME)"/.razor:/root/.razor \
    -v /path/of/password-file:/root/.razor/password-file \
    razornetwork/razor-go:latest \
   razor create --password /root/.razor/password-file
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

_Before staking on Razor Network, please ensure your account has eth and RAZOR. For testnet RAZOR, please contact us on Discord._

### Stake

If you have a minimum of 1000 razors in your account, you can stake those using the addStake command.

razor cli

```
$ ./razor addStake --address <address> --value <value>
```

docker

```
docker run -it  \
    -v "$(echo $HOME)"/.razor:/root/.razor \
    razornetwork/razor-go:latest \
    addStake --address <address> --value <value>
```

Example:

```
$ ./razor addStake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 1000
```

_Note: --pow flag is used to stake floating number stake_

_Note: Formula for calculating pow: (value \* (10**18)) / (10**x) where x is no of decimal places and value is integer_

_The value of pow is : 18 - x here_

If you have a 1000.25 razors in your account, you can stake those using the stake command with pow flag.

Example:

```
$ ./razor addStake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 100025 --pow 16
```

If you have a 5678.1001 razors in your account, you can stake those using the stake command with pow flag.

If you have a 5678.1001 razors in your account, you can stake those using the stake command with pow flag.

Example:

```
$ ./razor addStake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 56781001 --pow 14
```

### Staker Info

If you want to know the details of a staker, you can use stakerInfo command.

razor cli

```
$ ./razor stakerInfo --stakerId <staker_id_of_the_staker>
```

docker

```
docker run -it --rm \
    -v "$(echo $HOME)"/.razor:/root/.razor \
    razornetwork/razor-go:latest \
    stakerInfo --stakerId <staker_id_of_the_staker>
```

Example:

```
$ ./razor stakerInfo --stakerId 2
```

### Set Delegation

If you are a staker you can accept delegation from delegators and charge a commission from them.

razor cli

```
$ ./razor setDelegation --address <address> --status <true_or_false> --commission <commission_percent>
```

docker

```
docker run -it --rm \
    -v "$(echo $HOME)"/.razor:/root/.razor \
    razornetwork/razor-go:latest \
    setDelegation --address <address> --status <true_or_false>
```

Example:

```
$ ./razor setDelegation --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --status true -c 20
```

### Update Commission

If you are a staker and have accepted delegation, you can define your commission rate using this command.

razor cli

```
$ ./razor updateCommission --address <address> --commission <commission_percent>

```

docker

```
docker run -it --rm \
    -v "$(echo $HOME)"/.razor:/root/.razor \
    razornetwork/razor-go:latest \
    updateCommission --address <address> --commission <commission_percent>
```

Example:

```
$ ./razor updateCommission --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --commission 10
```

### Delegate

If you want to become a delegator use the `delegate` command. The staker whose `staker_id` is provided, their stake is increased.

razor cli

```
$ ./razor delegate --address <address> --value <value> --pow <power> --stakerId <staker_id>
```

docker

```
docker run -it --rm \
    -v "$(echo $HOME)"/.razor:/root/.razor \
    razornetwork/razor-go:latest \
    delegate --address <address> --value <value> --pow <power> --stakerId <staker_id>
```

Example:

```
$ ./razor delegate --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 1000 --pow 10 --stakerId 1
```

### Vote

You can start voting once you've staked some razors

razor cli

```
$ ./razor vote --address <address>
```

docker

```
docker run -it  -d --name razor-go \
    -v "$(echo $HOME)"/.razor:/root/.razor \
    razornetwork/razor-go:latest \
    vote --address <address>
```

Example:

```
$ ./razor vote --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c
```

If you want to report incorrect values, there is a `rogue` mode available. Just pass an extra flag `--rogue` to start voting in rogue mode and the client will report wrong medians.
The rogueMode key can be used to specify in which particular voting state (commit, reveal, propose) you want to report incorrect values.

Example:

```
$ ./razor vote --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --rogue --rogueMode commit,reveal,propose
```

### Unstake

If you wish to unstake your funds, you can run the `unstake` command.

razor cli

```
$ ./razor unstake --address <address> --stakerId <staker_id> --value <value> --pow <power>
```

docker

```
docker run -it  \
    -v "$(echo $HOME)"/.razor:/root/.razor \
    razornetwork/razor-go:latest \
    unstake --address <address> --stakerId <staker_id> --value <value> --pow <power>
```

Example:

```
$ ./razor unstake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --stakerId 1 --amount --pow 10 1000
```

### Withdraw

Once `unstake` has been called, you can withdraw your funds using the `withdraw` command

razor cli

```
$ ./razor withdraw --address <address> --stakerId <staker_id>
```

docker

```
docker run -it  \
    -v "$(echo $HOME)"/.razor:/root/.razor \
    razornetwork/razor-go:latest \
    withdraw --address <address> --stakerId <staker_id>
```

Example:

```
$ ./razor withdraw --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --stakerId 1
```

### Extend Lock

If the withdrawal period is over, then extendLock can be called to extend the lock period.

razor cli

```
$ ./razor extendLock --address <address> --stakerId <staker_id>
```

docker

```
docker run -it  \
    -v "$(echo $HOME)"/.razor:/root/.razor \
    razornetwork/razor-go:latest \
    extendLock --address <address> --stakerId <staker_id>
```

Example:

```
$ ./razor extendLock --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --stakerId 1
```

### Claim Bounty

If you want to claim your bounty after disputing a rogue staker, you can run `claimBounty` command

razor cli

```
$ ./razor claimBounty --address <address> --bountyId <bounty_id>
```

docker

```
docker run -it  \
    -v "$(echo $HOME)"/.razor:/root/.razor \
    razornetwork/razor-go:latest \
    claimBounty --address <address> --bountyId <bounty_id>
```

Example:

```
$ ./razor claimBounty --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --bountyId 5
```

### Transfer

Transfers razor to other accounts.

razor cli

```
$ ./razor transfer --value <value> --to <transfer_to_address> --from <transfer_from_address>
```

docker

```
docker run -it  \
    -v "$(echo $HOME)"/.razor:/root/.razor \
    razornetwork/razor-go:latest \
    transfer --value <value> --to <transfer_to_address> --from <transfer_from_address>
```

Example:

```
$ ./razor transfer --value 100 --to 0x91b1E6488307450f4c0442a1c35Bc314A505293e --from 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c
```

### Override Job

Jobs URLs are a placeholder for where to fetch values from. There is a chance that these URLs might either fail, or get razor nodes blacklisted, etc.
`overrideJob` command enables users to override the job URLs and selectors so that razor-nodes can fetch data directly from the override jobs.

razor cli

```
$ ./razor overrideJob --jobId <job_id_to_override> --url <new_url_of_job> --selector <selector_in_json_or_XHTML_selector_format> --power <power> --selectorType <0_for_XHTML_or_1_for_JSON>
```

docker

```
docker run -it  \
    -v "$(echo $HOME)"/.razor:/root/.razor \
    razornetwork/razor-go:latest \
    overrideJob --jobId <job_id_to_override> --url <new_url_of_job> --selector <selector_in_json_or_XHTML_selector_format> --power <power> --selectorType <0_for_XHTML_or_1_for_JSON>
```

Example:

```
$ ./razor overrideJob --jobId 2 --url https://api.gemini.com/v1/pubticker/ethusd --selector last --power 2 --selectorType 0
```

### Delete override

The overridden jobs can be deleted using `deleteOverride` command.

razor cli

```
$ ./razor deleteOverride --jobId <jobId>
```

docker

```
docker run -it  \
    -v "$(echo $HOME)"/.razor:/root/.razor \
    razornetwork/razor-go:latest \
    deleteOverride --jobId <jobId>
```

Example:

```
$ ./razor deleteOverride --jobId 2
```

### Set Config

There are a set of parameters that are configurable. These include:

- Provider: The RPC URL of the provider you are using to connect to the blockchain.
- Gas Multiplier: The value with which the gas price will be multiplied while sending every transaction.
- Buffer Size: Buffer size determines, out of all blocks in a state, in how many blocks the voting or any other operation can be performed.
- Wait Time: This is the number of blocks the system will wait while voting.
- Gas Price: The value of gas price if you want to set manually. If you don't provide any value or simply keep it to 0, the razor client will automatically calculate the optimum gas price and send it.
- Log Level: Normally debug logs are not logged into the log file. But if you want you can set `logLevel` to `debug` and fetch the debug logs.
- Gas Limit: The value with which the gas limit will be multiplied while sending every transaction.

The config is set while the build is generated, but if you need to change any of the above parameter, you can use the `setConfig` command.

razor cli

```
$ ./razor setConfig --provider <rpc_provider> --gasmultiplier <multiplier_value> --buffer <buffer_percentage> --wait <wait_for_n_blocks> --gasprice <gas_price> --logLevel <debug_or_info> --gasLimit <gas_limit_multiplier>
```

docker

```
docker run -it  \
    -v "$(echo $HOME)"/.razor:/root/.razor \
    razornetwork/razor-go:latest \
    setConfig --provider <rpc_provider> --gasmultiplier <multiplier_value> --buffer <buffer_percentage> --wait <wait_for_n_blocks> --gasprice <gas_price> --logLevel <debug_or_info> --gasLimit <gas_limit_multiplier>
```

Example:

```
$ ./razor setConfig --provider https://infura/v3/matic --gasmultiplier 1.5 --buffer 20 --wait 70 --gasprice 1 --logLevel debug --gasLimit 0.8
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

razor cli

```
$ ./razor createJob --url <URL> --selector <selector_in_json_or_XHTML_selector_format> --selectorType <0_for_XHTML_or_1_for_JSON> --name <name> --address <address> --power <power> --weight <weight>
```

docker

```
docker run -it  \
    -v "$(echo $HOME)"/.razor:/root/.razor \
    razornetwork/razor-go:latest \
    createJob --url <URL> --selector <selector_in_json_or_XHTML_selector_format> --selectorType <0_for_XHTML_or_1_for_JSON> --name <name> --address <address> --power <power> --weight <weight>
```

Example:

```
$ ./razor createJob --url https://www.alphavantage.co/query\?function\=GLOBAL_QUOTE\&symbol\=MSFT\&apikey\=demo --selector '[`Global Quote`][`05. price`]" --selectorType 1 --name msft --power 2 --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --weight 32
```

OR

```
$  ./razor createJob --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c -n btc_gecko --power 2 -s 'table tbody tr td span[data-coin-id="1"][data-target="price.price"] span' -u https://www.coingecko.com/en --selectorType 0 --weight 100
```

### Create Collection

Create new collections using `creteCollection` command.

_Note: This command is restricted to "Admin Role"_

razor cli

```
$ ./razor createCollection --name <collection_name> --address <address> --jobIds <list_of_job_ids> --aggregation <aggregation_method> --power <power> --tolerance <tolerance>
```

docker

```
docker run -it  \
    -v "$(echo $HOME)"/.razor:/root/.razor \
    razornetwork/razor-go:latest \
    createCollection --name <collection_name> --address <address> --jobIds <list_of_job_ids> --aggregation <aggregation_method> --power <power> --tolerance <tolerance>
```

Example:

```
$ ./razor createCollection --name btcCollectionMean --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --jobIds 1,2 --aggregation 2 --power 2 --tolerance 200
```

### Modify Asset Status

Modify the active status of an asset using the `modifyAssetStatus` command.

_Note: This command is restricted to "Admin Role"_

razor cli

```
$ ./razor modifyAssetStatus --assetId <assetId> --address <address> --status <true_or_false>
```

docker

```
docker run -it  \
    -v "$(echo $HOME)"/.razor:/root/.razor \
    razornetwork/razor-go:latest \
    modifyAssetStatus --assetId <assetId> --address <address> --status <true_or_false>
```

Example:

```
$ ./razor modifyAssetStatus --assetId 1 --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --status false
```

### Update Collection

Update the collection using `updateCollection` command.

_Note: This command is restricted to "Admin Role"_

razor cli

```
$ ./razor updateCollection --collectionId <collection_id> --jobIds <list_of_jobs> --address <address> --aggregation <aggregation_method> --power <power> --tolerance <tolerance>
```

docker

```
docker run -it  \
    -v "$(echo $HOME)"/.razor:/root/.razor \
    razornetwork/razor-go:latest \
    updateCollection --collectionId <collection_id> --jobIds <list_of_jobs> --address <address> --aggregation <aggregation_method> --power <power> --tolerance <tolerance>
```

Example:

```
$ ./razor updateCollection -a 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --collectionId 3 --jobIds 1,3 --aggregation 2 --power 4 --tolerance 5
```

### Update Job

Update the existing parameters of the Job using `updateJob` command.

_Note: This command is restricted to "Admin Role"_

razor cli

```
./razor updateJob --address <address> --jobID <job_Id> -s <selector> --selectorType <selectorType> -u <job_url> --power <power> --weight <weight>
```

docker

```
docker run -it  \
    -v "$(echo $HOME)"/.razor:/root/.razor \
    razornetwork/razor-go:latest \
    updateJob --address <address> --jobID <job_Id> -s <selector> --selectorType <selectorType> -u <job_url> --power <power> --weight <weight>
```

Example:

```
$ ./razor updateJob -a 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --jobId 1 -s last -u https://api.gemini.com/v1/pubticker/btcusd --power 2 --weight 10
```

### Job details

Get the list of all jobs with the details like weight, power, Id etc.

Example:

razor cli

```
$ ./razor jobList
```

docker

```
docker run -it  \
    -v "$(echo $HOME)"/.razor:/root/.razor \
    razornetwork/razor-go:latest \
    jobList
```

### Collection details

Get the list of all collections with the details like power, Id, name etc.

Example:

razor cli

```
$ ./razor collectionList
```

docker

```
docker run -it  \
    -v "$(echo $HOME)"/.razor:/root/.razor \
    razornetwork/razor-go:latest \
    collectionList
```

Note : _All the commands have an additional --password flag that you can provide with the file path from which password must be picked._



### Expose Metrics
Expose Prometheus-based metrics for monitoring

Example:

razor cli

```
$ ./razor setConfig --exposeMetrics 2112
```

docker

```
# Create docker network

docker network create razor_network

# Expose Metrics
docker run -it  --network razor_network\
    -v "$(echo $HOME)"/.razor:/root/.razor \
    razornetwork/razor-go:latest \
    setConfig --exposeMetrics 2112
```
### Contribute to razor-go

We would really appreciate your contribution. To see our [contribution guideline](https://github.com/razor-network/razor-go/blob/main/.github/CONTRIBUTING.md)

## Setting up development environment

1. Must have `docker` and `docker-compose` installed
2. Building the source `docker-compose build`
3. Create razor.yaml at $HOME/.razor/

   ```bash
   vi $HOME/.razor/razor.yaml
   ```

4. Add in razor.yaml and use :wq to exit form editor

   ```bash
   buffer: 20
   gaslimit: 2
   gasmultiplier: 1
   gasprice: 0
   provider: <rpc-url>
   wait: 30
   ```

5. Create account , and note address.

   ```bash
   docker-compose run razor-go /usr/local/bin/razor create
   ```

6. Import account

   ```
   docker-compose run razor-go /usr/local/bin/razor import
   ```

7. Get some **RAZOR** and **MATIC** token (or Token of respective RPC) to this address
8. Start **Staking**

   ```bash
   #Provide password through CLI
   docker-compose run razor-go /usr/local/bin/razor addStake --address <address> --value 50000

   #Provide password throudh File

     #Create file and put password string
       vi ~/.razor/pass
     #Start Staking
       docker-compose run razor-go /usr/local/bin/razor addStake --address <address> --value 50000 --password /root/.razor/pass

   ```

9. To Start **Voting**,

   1. Provide password through **CLI**

   ```bash
   # Run process in foreground and provide password through cli
   docker-compose run razor-go /usr/local/bin/razor vote --address <address>

   # Run process in background and provide password through file
   docker-compose run -d razor-go /usr/local/bin/razor vote --address <address> --password /root/.razor/pass
   ```

   1. Provide password through **File** and run in background with compose up
      1. replace <address> in docker-compose.yml with your address and create file pass and add your password in file

   ```bash
   docker-compose up -d
   ```

10. Enable Delegation

    ```bash
    #Provide password with cli
    docker-compose run razor-go /usr/local/bin/razor setDelegation --address <address> --status true --commission 10

    #provide password through file
    docker-compose run razor-go /usr/local/bin/razor setDelegation --address <address> --status true --commission 10 --password /root/.razor/pass
    ```
