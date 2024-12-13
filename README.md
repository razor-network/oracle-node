[![Coverage Status](https://coveralls.io/repos/github/razor-network/razor-go/badge.svg?branch=main)](https://coveralls.io/github/razor-network/razor-go?branch=main)

# Razor-Go

Official node for running stakers in Golang.

## Installation

### Linux quick start

Install the pre-built razor-go binary directly from GitHub and configure it on the host.

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

> **_NOTE:_** To install a specific version, set the VERSION:<git-tag> environment variable before running the command above.

## Docker quick start

One of the quickest ways to get `razor-go` up and running on your machine is by using Docker:

1. Create docker network

```
docker network create razor_network
```

2. Start razor-go container

```
docker run -d -it --entrypoint /bin/sh --network=razor_network --name razor-go -v "$(echo $HOME)"/.razor:/root/.razor razornetwork/razor-go:v2.0.0
```

> **_NOTE:_** We leverage Docker bind-mounts to mount the .razor directory, ensuring a shared mount between the host and the container. The `.razor` directory holds keys to the addresses that we use in `razor-go`, along with logs and config. We do this to persist data in the host machine, otherwise you would lose your keys once you delete the container.

You need to set a provider before you can operate razor-go cli on docker:

```
docker exec -it razor-go razor setConfig -p <provider_url>
```

You can now execute razor-go cli commands by running:

```
docker exec -it razor-go razor <command>
```

### Prerequisites to building the source

- Golang 1.23 or later must be installed.
- Latest stable version of node is required.
- Mac users with Silicon chips should use Node 18.18.0 LTS or later versions
- `geth` and `abigen` should be installed. (Skip this step if you don't want to fetch the bindings and build from scratch)
- `solc` and `jq` must be installed.

### Building the source

1. Run `npm install` to install the node dependencies.
2. If you want to build from scratch i.e., by fetching the smart contract bindings as well, run `npm run build-all`.

   _Note: To build from scratch, `geth` and `abigen` must be installed in your system._

3. If you already have the `pkg/bindings` you can run `npm run build` instead of `npm run build-all` to directly build the binary.
4. If you want to build the binary without wanting to set the configurations use `npm run build-noargs`
5. While building the binary, supply the provider RPC url and the gas multiplier.
6. The binary will be generated at `build/bin`.

## Commands

Go to the `build/bin` directory where the razor binary is generated.

`cd build/bin`

### Set Config

There are a set of parameters that are configurable. These include:

- Provider: The RPC URL of the provider you are using to connect to the blockchain.
- Alternate Provider: This is the secondary RPC URL of the provider used to connect to the blockchain if the primary one is not working.
- Gas Multiplier: The value with which the gas price will be multiplied while sending every transaction.
- Buffer Size: Buffer size determines, out of all blocks in a state, in how many blocks the voting or any other operation can be performed.
- Wait Time: This is the number of seconds the system will wait while voting.
- Gas Price: The value of gas price if you want to set manually. If you don't provide any value or simply keep it to 1, the razor client will automatically calculate the optimum gas price and send it.
- Log Level: Normally debug logs are not logged into the log file. But if you want you can set `logLevel` to `debug` and fetch the debug logs.
- Gas Limit: The value with which the gas limit will be multiplied while sending every transaction.
- Gas Limit Override: This value would be used as a gas limit for all the transactions instead of estimating for each transaction.
- RPC Timeout: This is the threshold number of seconds after which any contract and client calls will time out.
- HTTP Timeout: This is the threshold number of seconds after which an HTTP request for a job will time out.
- Maximum size of log file: This is the maximum size of log file in MB
- Maximum number of backups of log file: This is the maximum number of old log files to retain.
- Maximum age of log file: This is the maximum number of days to retain old log files.

The config is set while the build is generated, but if you need to change any of the above parameter, you can use the `setConfig` command.

razor cli

```
$ ./razor setConfig --provider <rpc_provider> --gasmultiplier <multiplier_value> --buffer <buffer_percentage> --wait <wait_for_n_blocks> --gasprice <gas_price> --logLevel <debug_or_info> --gasLimit <gas_limit_multiplier> --rpcTimeout <rpc_timeout> --httpTimeout <http_timeout> --logFileMaxSize <file_max_size> --logFileMaxBackups <file_max_backups> --logFileMaxAge <file_max_age>
```

docker

```
docker exec -it razor-go razor setConfig --provider <rpc_provider> --gasmultiplier <multiplier_value> --buffer <buffer_percentage> --wait <wait_for_n_blocks> --gasprice <gas_price> --logLevel <debug_or_info> --gasLimit <gas_limit_multiplier> --rpcTimeout <rpc_timeout> --httpTimeout <http_timeout> --logFileMaxSize <file_max_size> --logFileMaxBackups <file_max_backups> --logFileMaxAge <file_max_age>
```

Example:

```
$ ./razor setConfig --provider https://mainnet.skalenodes.com/v1/elated-tan-skat --gasmultiplier 1 --buffer 5 --wait 1 --gasprice 0 --logLevel debug --gasLimit 2 --rpcTimeout 5 --httpTimeout 5 --logFileMaxSize 200 --logFileMaxBackups 10 --logFileMaxAge 60
```

Besides, setting these parameters in the config, you can use different values for these parameters in various commands. Just add the same flag to any command you want to use and the new config changes will appear for that command.

Example:

```
$ ./razor vote --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --gasprice 1
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

If you already have an account and its private key, you can import that account into the `razor-go` client.
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

_Before staking on Razor Network, please ensure your account has sFUEL and RAZOR. For testnet RAZOR, please contact us on Discord._

### Import Endpoints

You can import the endpoints to file `$HOME/.razor/endpoints.json` on your local by using the `importEndpoints` command.
This command imports multiple providers along with the user input provider, which are then sorted according to the best performance. The best provider is thus chosen by the RPC manager and will be used to make the RPC calls.

razor cli

```
$ ./razor importEndpoints
```

docker

```
docker exec -it razor-go razor importEndpoints
```

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

If you have 1000.25 razors in your account, you can stake those using the stake command with weiRazor flag.

Example:

```
$ ./razor addStake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 1000250000000000000000 --weiRazor true
```

If you have 5678.1001 razors in your account, you can stake those using the stake command with weiRazor flag.

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

If you are a staker, you can accept delegations from delegators and charge them a commission.
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

Stakers can claim the rewards earned from a delegator's pool share as commission using `claimCommission`

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

You can start voting once you've staked some RAZORs.

razor cli

```
$ ./razor vote --address <address>
```

docker

```
docker exec -it razor-go razor vote --address <address>
```

> **Note**: _To run vote command in background you can use `tmux` for that._
>
> 1.  Run: `tmux new -s razor-go`
> 2.  Run vote command
> 3.  To exit from tmux session: press `ctrl+b`, release those keys and press `d`
> 4.  To list your session: `tmux ls`
> 5.  To attach Session back: `tmux attach-session -t razor-go`


Example:

```
$ ./razor vote --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c
```

If you want to claim your bounty automatically after disputing staker, you can just pass `--autoClaimBounty` flag in your vote command.

If you are running an extra backup node, it is suggested to avoid performing few actions. So to do that you need to pass actions that need to be ignored as a value to flag `--backupNode <actions_To_Ignore>`.
For now, we only support `disputeMedians` as actions to be ignored as a value for backup node.
```
$ ./razor vote --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --backupNode disputeMedians
```

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

If the withdrawal period has ended, you can use the extendLock command to extend the lock period.

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

> **_NOTE:_** Bounty IDs are stored in the .razor directory with filenames in the format `YOUR_ADDRESS_disputeData.json file.`
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

Transfers RAZOR to other accounts.

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

Create new jobs using the `createJob` command.

_Note: This command is restricted to users with the "Admin Role"_

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
$ ./razor createJob --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c -n btc_gecko --power 2 -s 'table tbody tr td span[data-coin-id="1"][data-target="price.price"] span' -u https://www.coingecko.com/en --selectorType 0 --weight 100
```

### Create Collection

Create new collections using the `createCollection` command.

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

Modify the active status of a collection using the `modifyCollectionStatus` command.

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

Update the collection using the `updateCollection` command.

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

Update the existing parameters of the Job using the `updateJob` command.

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
docker exec -it razor-go razor collectionList
```

Note : _All commands include an additional --password flag. You can specify a file path to retrieve the password._

### Expose Metrics

Expose Prometheus-based metrics for monitoring

#### Without TLS

```
$ ./razor setConfig --exposeMetrics 2112
```

#### With TLS

```
$ ./razor setConfig --exposeMetrics 2112 --certFile /cert/file/path/certfile.crt --certKey key/file/path/keyfile.key
```

docker

#### Expose Metrics without TLS

```
docker exec -it razor-go razor setConfig --exposeMetrics 2112
```

#### Expose Metrics with TLS

```
docker exec -it razor-go razor setConfig --exposeMetrics 2112 --certFile /cert/file/path/certfile.crt --certKey key/file/path/keyfile.key
```

#### Configuration

Clone the repo and setup monitoring and alerting using Prometheus/Grafana

```
git clone https://github.com/razor-network/monitoring.git
cd monitoring
```

- If your staker runs via binary, then

    1. In `./configs/prometheus.yml`, replace `"razor-go:2112"` with `"<private/public address of host>:2112"`

- For alerting, you can add a webhook in `./configs/alertmanager.yml`. Replace `http://127.0.0.1:5001/` with your webhook URL. This will send an alert every 5 minutes if the metrics stop.

- If you are running multiple stakers and want to monitor via single grafana dashboard

    1. You need to update `./config/prometheus.yml`, add new target block where `job_name: "razor-go"`
        ```
        - targets: ["<second-host-address>:2112"]
          labels:
            staker: "<staker-name>"
        ```
    2. Restart vmagent service `docker-compose restart vmagent`
#### Start monitoring stack
-  You can spin all agents at once via

    ```
    docker-compose up -d
    ``` 
   Can check the status of each service via
    ```
    docker-compose ps
    ```

- You can open grafana at `<private/public address of host>:3000`, and get
    1. Can checkout `Razor` dashboard to monitor your staker.
    2. Insight of host metrics at `Node Exporter Full` dashboard.
    3. Containers Insight at `Docker and OS metrics ( cadvisor, node_exporter )` dashboard.
    4. Can monitor alerts at `Alertmanager` dashboard.

>**_NOTE:_** Configure firewall for port `3000` on your host to access grafana.

#### Troubleshoot Alerting

1. In `docker-compose.yml` uncomment ports for `alertmanager` and `vmalert`.
2. Configure firewall to allow access to ports `8880` and `9093`.

3. Check you get alerts on vmalert via `http://<host_address>:8880/vmalert/alerts`. vmalert is configured to scrap in every 2min.

4. If you see alert in vmalert then look into alertmanager `http://<host_address>:9093/#/alerts?`, if you see alerts in there but you didn't get one then probably you need to check your weebhook.

#### Configuration

Clone repo and setup monitoring and alerting using Prometheus/Grafana

```
git clone https://github.com/razor-network/monitoring.git
cd monitoring
```

- If your staker is running via binary, then

    1. In `./configs/prometheus.yml`, replace `"razor-go:2112"` with `"<private/public address of host>:2112"`

- For alerting you can add webhook in `./configs/alertmanager.yml`, replace `http://127.0.0.1:5001/` with your webhook URL. This will send you an alert in every 5min if metrics stops.

- If you are running multiple stakers and want to monitor via single grafana dashboard
    1. You need to update `./config/prometheus.yml`, add new target block where `job_name: "razor-go"`
       ```
       - targets: ["<second-host-address>:2112"]
         labels:
           staker: "<staker-name>"
       ```
    2. Restart vmagent service `docker-compose restart vmagent`

#### Start monitoring stack

- You can spin all agents at once via

  ```
  docker-compose up -d
  ```

  Can check the status of each service via

  ```
  docker-compose ps
  ```

- You can open grafana at `<private/public address of host>:3000`, and get
    1. Can checkout `Razor` dashboard to monitor your staker.
    2. Insight of host metrics at `Node Exporter Full` dashboard.
    3. Containers Insight at `Docker and OS metrics ( cadvisor, node_exporter )` dashboard.
    4. Can monitor alerts at `Alertmanager` dashboard.

> **_NOTE:_** Configure firewall for port `3000` on your host to access grafana.

#### Troubleshoot Alerting

1. In `docker-compose.yml` uncomment ports for `alertmanager` and `vmalert`.
2. Configure firewall to allow access to ports `8880` and `9093`.

3. Check you get alerts on vmalert via `http://<host_address>:8880/vmalert/alerts`. vmalert is configured to scrap in every 2min.

4. If you see alert in vmalert then look into alertmanager `http://<host_address>:9093/#/alerts?`, if you see alerts in there but you didn't get one then probably you need to check your weebhook.

### Override Job and Adding Your Custom Jobs

Job URLs act as placeholders indicating where values should be fetched from. There is a chance that these URLs might either fail, or get razor nodes blacklisted, etc.
You can override the existing job and also add your custom jobs by adding `assets.json` file in `.razor` directory so that razor-nodes can fetch data directly from the provided jobs.

Shown below is an example of how your `assets.json` file should be -

``` json
{
    "assets": {
      "collection": {
        "ETHUSD": {
          "official jobs": {
            "1": {
              "URL": "https://data.messari.io/api/v1/assets/eth/metrics",
              "selector": "[`data`][`market_data`][`price_usd`]",
              "power": 2,
              "weight": 2
            }
          },
          "custom jobs": [
            {
              "URL": "https://api.kucoin.com/api/v1/prices?base=USD&currencies=ETH",
              "name": "eth_kucoin_usd",
              "selector": "data.ETH",
              "power": 3,
              "weight": 1
            },
            {
              "URL": {
                "type": "POST",
                "url": "https://rpc.ankr.com/eth",
                "body": {
                  "jsonrpc": "2.0",
                  "method": "eth_call",
                  "params": [
                    {
                      "to": "0xb27308f9f90d607463bb33ea1bebb41c27ce5ab6",
                      "data": "0xf7729d43000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb480000000000000000000000000000000000000000000000000000000000000bb80000000000000000000000000000000000000000000000000de0b6b3a76400000000000000000000000000000000000000000000000000000000000000000000"
                    },
                    "latest"
                  ],
                  "id": 5
                },
                "header": {
                  "content-type": "application/json"
                },
                "returnType": "hex"
              },
              "name": "eth_postJob_usd",
              "power": -4,
              "selectorType": 0,
              "selector": "result",
              "weight": 1
            }
          ]
        }
      }
    }
  }
```

Breaking down into components

- The existing jobs that you want to override should be included in `official jobs` and fields like URL, selector should be replaced with your provided inputs respectively.

In the above example for the collection `ethCollectionMean`, the job with `jobId:1` is overriden by provided URL, selector, power and weight.

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

In the above example for the collection `ethCollectionMean` new custom jobs are as shown below, 
1. Job following GET request having URL `https://api.kucoin.com/api/v1/prices?base=USD&currencies=ETH` and
2. Job following POST request having URL `"https://rpc.ankr.com/eth"` with respective `body` and `header` will be added in jobs array.

If any custom job requires authentication via an API key or headers, the staker can export the key using the method shown below:

If the job is:
```
https://api.gemini.com/v1/pubticker/v1/exchangerate/BTC?apikey=YOUR_AUTH_KEY
```

you can change the above job url to 
```
https://api.gemini.com/v1/pubticker/v1/exchangerate/BTC?apikey=${AUTH_KEY}
```

Now staker needs to use the same keyword defined inside `${...}` as an environment variable using `export` command and assigning it a value as users API key as shown below,

```
export AUTH_KEY="YOUR_AUTH_KEY"
```

### Logs

Users can pass a separate `--logFile` flag followed by any desired log file name when executing a command. The logs will be stored in `.razor/logs` directory.

razor cli

```
$ ./razor addStake --address <address> --value <value> --logFile stakingLogs
```

docker

```
docker exec -it razor-go razo addStake --address <address> --value <value> --logFile stakingLogs
```

_The logs for the above command will be stored at the "$HOME/.razor/logs/stakingLogs.log" path_

razor cli

```
$ ./razor delegate --address <address> --value <value> --weiRazor <bool> --stakerId <staker_id> --logFile delegationLogs
```

docker

```
docker exec -it razor-go razo delegate --address <address> --value <value> --weiRazor <bool> --stakerId <staker_id> --logFile delegationLogs
```

_The logs for above command will be stored at "$HOME/.razor/logs/delegationLogs.log" path_

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

7. Get some **RAZOR** and **sFUEL** token (or Token of respective RPC) to this address
8. Start **Staking**

   ```bash
   #Provide password through CLI
   docker-compose run razor-go /usr/local/bin/razor addStake --address <address> --value 50000
   
      #Provide password through File

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

### Contribute to razor-go

We would really appreciate your contribution. To see our [contribution guideline](https://github.com/razor-network/razor-go/blob/main/.github/CONTRIBUTING.md)
