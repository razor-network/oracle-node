[![Coverage Status](https://coveralls.io/repos/github/razor-network/razor-go/badge.svg?branch=main)](https://coveralls.io/github/razor-network/razor-go?branch=main)

# Razor-Go

Official node for running stakers in Golang.

## Installation

### Linux quick start

Install `razor-go` pre build binary directly from github and configure into host.

  For linux-amd64
  ```
  curl -sSL https://raw.githubusercontent.com/razor-network/razor-go/main/install.sh | bash 
  ```

  For linux-arm64
  ```
  export PLATFORM=arm64

  curl -sSL https://raw.githubusercontent.com/razor-network/razor-go/main/install.sh | bash 
  ```

Check installation

```
razor -v
```
>**_NOTE:_** To install the version you want, you can set VERSION:<git-tag> environment variable before running above command.
## Docker quick start

One of the quickest ways to get `razor-go` up and running on your machine is by using Docker:
```
docker run -d -it--entrypoint /bin/sh  --name razor-go -v "$(echo $HOME)"/.razor:/root/.razor razornetwork/razor-go:v1.0.1-incentivised-testnet-phase2
```

>**_NOTE:_** that we are leveraging docker bind-mounts to mount `.razor` directory so that we have a shared mount of `.razor` directory between the host and the container. The `.razor` directory holds keys to the addresses that we use in `razor-go`, along with logs and config. We do this to persist data in the host machine, otherwise you would lose your keys once you delete the container.

You need to set a provider before you can operate razor-go cli on docker:

```
docker exec -it razor-go razor setConfig -p <provider_url>
```

You can now execute razor-go cli commands by running:

```
docker exec -it razor-go razor <command>
```

### Prerequisites to building the source

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
6. To bypass the interactive mode of providing password, create file in `.razor` directory with providing password in it.
7. The binary will be generated at `build/bin`.

## Commands

Go to the `build/bin` directory where the razor binary is generated.

`cd build/bin`


### Set Config

There are a set of parameters that are configurable. These include:

- Provider: The RPC URL of the provider you are using to connect to the blockchain.
- Gas Multiplier: The value with which the gas price will be multiplied while sending every transaction.
- Buffer Size: Buffer size determines, out of all blocks in a state, in how many blocks the voting or any other operation can be performed.
- Wait Time: This is the number of blocks the system will wait while voting.
- Gas Price: The value of gas price if you want to set manually. If you don't provide any value or simply keep it to 1, the razor client will automatically calculate the optimum gas price and send it.
- Log Level: Normally debug logs are not logged into the log file. But if you want you can set `logLevel` to `debug` and fetch the debug logs.
- Gas Limit: The value with which the gas limit will be multiplied while sending every transaction.

The config is set while the build is generated, but if you need to change any of the above parameter, you can use the `setConfig` command.

razor cli

```
$ ./razor setConfig --provider <rpc_provider> --gasmultiplier <multiplier_value> --buffer <buffer_percentage> --wait <wait_for_n_blocks> --gasprice <gas_price> --logLevel <debug_or_info> --gasLimit <gas_limit_multiplier>
```

docker

```
docker exec -it razor-go razor setConfig --provider <rpc_provider> --gasmultiplier <multiplier_value> --buffer <buffer_percentage> --wait <wait_for_n_blocks> --gasprice <gas_price> --logLevel <debug_or_info> --gasLimit <gas_limit_multiplier>
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



## Razor commands

### Create Account

Create an account using the `create` command. You'll be asked to enter a password that'll be used to encrypt the keystore file.

razor cli

```
$ ./razor create

```

Docker

```
docker exec -it razor-go razor create
```
Example:

```
$ ./razor create
Password:
```

### Import Account

If you already have an account created, and have it's private key, that account can be imported into the `razor-go` client.
To do that, you can use the `import` command. You'll be asked the private key first and then the password which you want to encrypt your keystore file with.

razor cli

```
$ ./razor import
```

docker

```
docker exec -it razor-go razor import
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
docker exec -it razor-go razor addStake --address <address> --value <value>
```

Example:

```
$ ./razor addStake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 1000
```

_Note: --weiRazor flag can be passed to provide values in wei_

If you have a 1000.25 razors in your account, you can stake those using the stake command with weiRazor flag.

Example:

```
$ razor addStake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 1000250000000000000000 --weiRazor true
```

If you have a 5678.1001 razors in your account, you can stake those using the stake command with weiRazor flag.

Example:

```
$ ./razor addStake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 5678100100000000000000 --weiRazor true
```

### Staker Info

If you want to know the details of a staker, you can use stakerInfo command.

razor cli

```
$ ./razor stakerInfo --stakerId <staker_id_of_the_staker>
```

docker

```
docker exec -it razor-go razor stakerInfo --stakerId <staker_id_of_the_staker>
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
docker exec -it razor-go razor setDelegation --address <address> --status <true_or_false> --commission <commission_percent>
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
docker exec -it razor-go razor updateCommission --address <address> --commission <commission_percent>
```

Example:

```
$ ./razor updateCommission --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --commission 10
```

### Delegate

If you want to become a delegator use the `delegate` command. The staker whose `staker_id` is provided, their stake is increased.

razor cli

```
$ ./razor delegate --address <address> --value <value> --weiRazor <bool> --stakerId <staker_id>
```

docker

```
docker exec -it razor-go razor delegate --address <address> --value <value> --weiRazor <bool> --stakerId <staker_id>
```

Example:

```
$ ./razor delegate --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 1000 --weiRazor false --stakerId 1
```

### Claim Commission 

Staker can claim the rewards earned from delegator's pool share as commission using `claimCommission`

razor cli

```
$ ./razor claimCommission --address <address> 
```

docker

```
docker exec -it razor-go razor claimCommission --address <address> 
```

Example:

```
$ ./razor claimCommission --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c 
```

### Vote

You can start voting once you've staked some razors

razor cli

```
$ ./razor vote --address <address>
```

docker

```
docker exec -it razor-go razor vote --address <address>
```

run vote command in background
```
docker exec -it -d razor-go razor vote --address <address> 
```


Example:

```
$ ./razor vote --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c
```
If you want to claim your bounty automatically after disputing staker, you can just pass `--autoClaimBounty` flag in your vote command.

If you want to report incorrect values, there is a `rogue` mode available. Just pass an extra flag `--rogue` to start voting in rogue mode and the client will report wrong medians.
The rogueMode key can be used to specify in which particular voting state (commit, reveal) or for which values i.e. medians/revealedIds (medians, missingIds, extraIds, unsortedIds)you want to report incorrect values.

Example:

```
$ ./razor vote --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --rogue --rogueMode commit,reveal,medians,missingIds,extraIds,unsortedIds
```

### Unstake

If you wish to unstake your funds, you can run the `unstake` command.

razor cli

```
$ ./razor unstake --address <address> --stakerId <staker_id> --value <value> --weiRazor <bool>
```

docker

```
docker exec -it razor-go razor unstake --address <address> --stakerId <staker_id> --value <value> --weiRazor <bool>
```

Example:

```
$ ./razor unstake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --stakerId 1 --value 1000 --weiRazor false
```

### Withdraw

Once `unstake` has been called, you can withdraw your funds using the `initiateWithdraw` and `unlockWithdraw` commands

You need to start the withdrawal process using `initiateWithdraw` command and once the withdraw lock period is over you can use `unlockWithdraw` command to get the RZR's back to your account.

razor cli

```
$ ./razor initiateWithdraw --address <address> --stakerId <staker_id>
```

```
$ ./razor unlockWithdraw --address <address> --stakerId <staker_id>
```
docker

```
docker exec -it razor-go razor initiateWithdraw --address <address> --stakerId <staker_id>
```

```
docker exec -it razor-go razor unlockWithdraw --address <address> --stakerId <staker_id>
```

Example:

```
$ ./razor initiateWithdraw --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --stakerId 1
```
```
$ ./razor unlockWithdraw --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --stakerId 1
```

### Extend Lock

If the withdrawal period is over, then extendLock can be called to extend the lock period.

razor cli

```
$ ./razor extendLock --address <address> --stakerId <staker_id>
```

docker

```
docker exec -it razor-go razor extendLock --address <address> --stakerId <staker_id>
```

Example:

```
$ ./razor extendLock --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --stakerId 1
```

### Claim Bounty

If you want to claim your bounty after disputing a rogue staker, you can run `claimBounty` command

>**_NOTE:_**  bountyIds are stored in .razor directory with file name in format `YOUR_ADDRESS_disputeData.json file.`
>
> e.g: `0x2EDc3c6F93e4e20590F480272AB490D2620557xY_disputeData.json`
If you know the bountyId, you can pass the value to `bountyId` flag.

razor cli

```
$ ./razor claimBounty --address <address> --bountyId <bounty_id>
```

docker

```
docker exec -it razor-go razor claimBounty --address <address> --bountyId <bounty_id>
```

Example:

```
$ ./razor claimBounty --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --bountyId 5
```

You can also run claimBounty command without passing `bountyId` flag as it will pick up bountyIds associated to your address from the file one at a time.

razor cli

```
$ ./razor claimBounty --address <address> 
```

docker

```
docker exec -it razor-go razor claimBounty --address <address> 
```

### Transfer

Transfers razor to other accounts.

razor cli

```
$ ./razor transfer --value <value> --to <transfer_to_address> --from <transfer_from_address>
```

docker

```
docker exec -it razor-go razor transfer --value <value> --to <transfer_to_address> --from <transfer_from_address>
```

Example:

```
$ ./razor transfer --value 100 --to 0x91b1E6488307450f4c0442a1c35Bc314A505293e --from 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c
```

### Create Job

Create new jobs using `creteJob` command.

_Note: This command is restricted to "Admin Role"_

razor cli

```
$ ./razor createJob --url <URL> --selector <selector_in_json_or_XHTML_selector_format> --selectorType <0_for_XHTML_or_1_for_JSON> --name <name> --address <address> --power <power> --weight <weight>
```

docker

```
docker exec -it razor-go razor createJob --url <URL> --selector <selector_in_json_or_XHTML_selector_format> --selectorType <0_for_XHTML_or_1_for_JSON> --name <name> --address <address> --power <power> --weight <weight>
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
docker exec -it razor-go razor createCollection --name <collection_name> --address <address> --jobIds <list_of_job_ids> --aggregation <aggregation_method> --power <power> --tolerance <tolerance>
```

Example:

```
$ ./razor createCollection --name btcCollectionMean --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --jobIds 1,2 --aggregation 2 --power 2 --tolerance 200
```

### Modify Collection Status

Modify the active status of an collection using the `modifyCollectionStatus` command.

_Note: This command is restricted to "Admin Role"_

razor cli

```
$ ./razor modifyCollectionStatus --collectionId <collectionId> --address <address> --status <true_or_false>
```

docker

```
docker exec -it razor-go razor modifyCollectionStatus --collectionId <collectionId> --address <address> --status <true_or_false>
```

Example:

```
$ ./razor modifyCollectionStatus --collectionId 1 --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --status false
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
docker exec -it razor-go razor updateCollection --collectionId <collection_id> --jobIds <list_of_jobs> --address <address> --aggregation <aggregation_method> --power <power> --tolerance <tolerance>
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
docker exec -it razor-go razor updateJob --address <address> --jobID <job_Id> -s <selector> --selectorType <selectorType> -u <job_url> --power <power> --weight <weight>
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
docker exec -it razor-go razor jobList
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
docker exec -it razor-go razorcollectionList
```

### Expose Metrics
Expose Prometheus-based metrics for monitoring

Example:

razor cli

Without TLS
```
$ ./razor setConfig --exposeMetrics 2112
```
With TLS
```
$ ./razor setConfig --exposeMetrics 2112 --certFile /cert/file/path/certfile.crt --certKey key/file/path/keyfile.key
```

docker

```
# Create docker network

docker network create razor_network

# Expose Metrics without TLS
docker exec -it razor-go razor setConfig --exposeMetrics 2112

# Expose Metrics with TLS
docker exec -it razor-go razor setConfig --exposeMetrics 2112 --certFile /cert/file/path/certfile.crt --certKey key/file/path/keyfile.key
```

### Override Job and Adding Your Custom Jobs

Jobs URLs are a placeholder from where to fetch values from. There is a chance that these URLs might either fail, or get razor nodes blacklisted, etc.
You can override the existing job and also add your custom jobs by adding `assets.json` file in `.razor` directory so that razor-nodes can fetch data directly from the provided jobs.

Shown below is an example of how your `assets.json` file should be -
```
{
  "assets": {
    "collection": {
      "ethCollectionMean": {
        "power": 2,
        "official jobs": {
          "1": {
            "URL": "https://data.messari.io/api/v1/assets/eth/metrics",
            "selector": "[`data`][`market_data`][`price_usd`]",
            "power": 2,
            "weight": 2
          },
        },
        "custom jobs": [
          {
            "URL": "https://api.lunarcrush.com/v2?data=assets&symbol=ETH",
            "selector": "[`data`][`0`][`price`]",
            "power": 3,
            "weight": 2
          },
        ]
      }
    }
  }
}
```

Breaking down into components
- The existing jobs that you want to override should be included in `official jobs` and fields like URL, selector should be replaced with your provided inputs respectively.

In the above example for the collection `ethCollectionMean`, job having `jobId:1` is override by provided URL, selector, power and weight.
```
"official jobs": {
          "1": {
            "URL": "https://data.messari.io/api/v1/assets/eth/metrics",
            "selector": "[`data`][`market_data`][`price_usd`]",
            "power": 2,
            "weight": 2
          },
```

- Additional jobs that you want to add to a collection should be added in `custom jobs` field with their respective URLs and selectors.

In the above example for the collection `ethCollectionMean`, new custom job having URL `https://api.lunarcrush.com/v2?data=assets&symbol=ETH` is added.
```
 "custom jobs": [
          {
            "URL": "https://api.lunarcrush.com/v2?data=assets&symbol=ETH",
            "selector": "[`data`][`0`][`price`]",
            "power": 3,
            "weight": 2
          },
        ]
```

### Logs

User can pass a separate flag --logFile followed with any name for log file along with command. The logs will be stored in ```.razor/logs``` directory.

razor cli
```
$ ./razor addStake --address <address> --value <value> --logFile stakingLogs
```
docker
```
docker exec -it razor-go razo addStake --address <address> --value <value> --logFile stakingLogs
```
_The logs for above command will be stored at "home/.razor/stakingLogs.log" path_

razor cli
```
$ ./razor delegate --address <address> --value <value> --weiRazor <bool> --stakerId <staker_id> --logFile delegationLogs
```
docker
```
docker exec -it razor-go razo delegate --address <address> --value <value> --weiRazor <bool> --stakerId <staker_id> --logFile delegationLogs
```
_The logs for above command will be stored at "home/.razor/delegationLogs.log" path_

_Note: If the user runs multiple commands with the same log file name all the logs will be appended in the same log file._


### Contract Addresses

This command provides the list of contract addresses.

razor cli

```
$ ./razor contractAddresses
```

docker

```
docker exec -it razor-go razor contractAddresses
```

Example:

```
$ ./razor contractAddresses
```


### Setting up razor-go and commands using docker-compose

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
   docker-compose run razor-go /usr
   /local/bin/razor create
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
   ```

9. To Start **Voting**,

   1. Provide password through **CLI**

   ```bash
   # Run process in foreground and provide password through cli
   docker-compose run razor-go /usr/local/bin/razor vote --address <address>
   ```

   ```bash
   docker-compose up -d
   ```

10. Enable Delegation

    ```bash
    #Provide password with cli
    docker-compose run razor-go /usr/local/bin/razor setDelegation --address <address> --status true --commission 10
    ```

### Contribute to razor-go

We would really appreciate your contribution. To see our [contribution guideline](https://github.com/razor-network/razor-go/blob/main/.github/CONTRIBUTING.md)