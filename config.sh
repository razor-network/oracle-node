#!/bin/bash

set -e -o pipefail

BIN_DIR=./build/bin
RAZOR=${BIN_DIR}/razor

read -rp "Provider: (http://127.0.0.1:8545) " PROVIDER
if [ -z "$PROVIDER" ];
then
  PROVIDER="http://127.0.0.1:8545"
fi

read -rp "Alternate Provider: " ALTERNATE_PROVIDER

read -rp "Gas Multiplier: (1.0) " GAS_MULTIPLIER
if [ -z "$GAS_MULTIPLIER" ];
then
  GAS_MULTIPLIER=1.0
fi

read -rp "Buffer Percent: (20) " BUFFER
if [ -z "$BUFFER" ];
then
  BUFFER=20
fi

read -rp "Wait Time: (5) " WAIT_TIME
if [ -z "$WAIT_TIME" ]; then
   WAIT_TIME=5
fi

read -rp "Gas Price: (0) " GAS_PRICE
if [ -z "$GAS_PRICE" ]; then
   GAS_PRICE=0
fi

read -rp "Gas Limit Increment : (2) " GAS_LIMIT
if [ -z "$GAS_LIMIT" ]; then
   GAS_LIMIT=2
fi

read -rp "Log File Max Size: (200) " MAX_SIZE
if [ -z "$MAX_SIZE" ]; then
   MAX_SIZE=200
fi

read -rp "Log Files Max Backups: (52) " MAX_BACKUPS
if [ -z "$MAX_BACKUPS" ]; then
   MAX_BACKUPS=52
fi

read -rp "Log Files Max Age: (365) " MAX_AGE
if [ -z "$MAX_AGE" ]; then
   MAX_AGE=365
fi

ALT_PROVIDER_OPTION=""
if [ -n "$ALTERNATE_PROVIDER" ]; then
    ALT_PROVIDER_OPTION="--alternateProvider $ALTERNATE_PROVIDER"
fi

$RAZOR setConfig -p $PROVIDER $ALT_PROVIDER_OPTION -b $BUFFER -g $GAS_MULTIPLIER -w $WAIT_TIME --gasprice $GAS_PRICE --gasLimit $GAS_LIMIT --rpcTimeout 10 --httpTimeout 10 --logFileMaxSize $MAX_SIZE --logFileMaxBackups $MAX_BACKUPS --logFileMaxAge $MAX_AGE