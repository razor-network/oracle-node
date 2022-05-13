#!/bin/bash
set -e

bin_dir="${BIN_DIR:-/usr/local/bin}"
platform="${PLATFORM:-amd64}"

if [[ -z "$VERSION" ]]
then
    version=$(curl -L -s https://storage.googleapis.com/razor-go-release/stable.txt)
else
    version="${VERSION:-v0.0.0}"
fi
echo "Installing linux-$platform:$version at path $bin_dir."

wget -q "https://github.com/razor-network/razor-go/releases/download/$version/razor_go.linux-$platform.tar.gz" \
    -O /tmp/razor_go.tar.gz
echo "Downloaded linux-$platform:$version at path $bin_dir."
mkdir -p /tmp/razor_go

cd /tmp || { echo "ERROR! No /tmp found.."; exit 1; }

tar xfz /tmp/razor_go.tar.gz -C /tmp/razor_go || { echo "ERROR! Extracting the razor_go tar"; exit 1; }
echo "Moving to bin directory"
mv "/tmp/razor_go/razor" "$bin_dir"
chown root:staff "$bin_dir/razor"
echo "SUCCESS! Installation succeeded!"
