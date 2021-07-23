#!/bin/bash

set -e -o pipefail

BIN_DIR=./build/bin
RAZOR=${BIN_DIR}/razor

read -rp "Provider: (http://127.0.0.1:8545) " PROVIDER
if [ -z "$PROVIDER" ];
then
  PROVIDER="http://127.0.0.1:8545"
fi

read -rp "Gas Multiplier: (1.0) " GASMULTIPLIER
if [ -z "$GASMULTIPLIER" ];
then
  GASMULTIPLIER=1.0
fi

read -rp "Buffer Percent: (20) " BUFFER
if [ -z "$BUFFER" ];
then
  BUFFER=20
fi

$RAZOR setconfig -p $PROVIDER -b $BUFFER -g $GASMULTIPLIER