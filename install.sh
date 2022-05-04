#!/bin/bash
set -e
version="${VERSION:-v0.0.0}"
bin_dir="${BIN_DIR:-/usr/local/bin}"
platform="${PLATFORM:-amd64}"
wget "https://github.com/razor-network/razor-go/releases/download/$version/razor_go.linux-$platform.tar.gz" \
    -O /tmp/razor_go.tar.gz

mkdir -p /tmp/razor_go

cd /tmp || { echo "ERROR! No /tmp found.."; exit 1; }

tar xfz /tmp/razor_go.tar.gz -C /tmp/razor_go || { echo "ERROR! Extracting the razor_go tar"; exit 1; }

cp "/tmp/razor_go/razor" "$bin_dir"
chown root:staff "$bin_dir/razor"

echo "SUCCESS! Installation succeeded!"
