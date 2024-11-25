#!/bin/bash

set -e -o pipefail

BIN_DIR=./build/bin
RAZOR=${BIN_DIR}/razor

read -rp "Provider: (http://127.0.0.1:8545) " PROVIDER
if [ -z "$PROVIDER" ];
then
  PROVIDER="http://127.0.0.1:8545"
fi

read -rp "Gas Multiplier: (1.0) " GAS_MULTIPLIER
if [ -z "$GAS_MULTIPLIER" ];
then
  GAS_MULTIPLIER=1.0
fi

read -rp "Buffer Percent: (5) " BUFFER
if [ -z "$BUFFER" ];
then
  BUFFER=5
fi

read -rp "Wait Time: (1) " WAIT_TIME
if [ -z "$WAIT_TIME" ]; then
   WAIT_TIME=1
fi

read -rp "Gas Price: (1) " GAS_PRICE
if [ -z "$GAS_PRICE" ]; then
   GAS_PRICE=1
fi

read -rp "Gas Limit Increment : (2) " GAS_LIMIT
if [ -z "$GAS_LIMIT" ]; then
   GAS_LIMIT=2
fi

read -rp "Log File Max Size: (200) " MAX_SIZE
if [ -z "$MAX_SIZE" ]; then
   MAX_SIZE=200
fi

read -rp "Log Files Max Backups: (10) " MAX_BACKUPS
if [ -z "$MAX_BACKUPS" ]; then
   MAX_BACKUPS=10
fi

read -rp "Log Files Max Age: (60) " MAX_AGE
if [ -z "$MAX_AGE" ]; then
   MAX_AGE=60
fi

$RAZOR setConfig -p $PROVIDER -b $BUFFER -g $GAS_MULTIPLIER -w $WAIT_TIME --gasprice $GAS_PRICE --gasLimit $GAS_LIMIT --rpcTimeout 5 --httpTimeout 5 --logFileMaxSize $MAX_SIZE --logFileMaxBackups $MAX_BACKUPS --logFileMaxAge $MAX_AGE