yarn install

# Check path of contracts
abigen --abi build/contracts/SchellingCoin.json --pkg main --type SchellingCoin --out SchellingCoin.go

# Install go dependencies
go mod install

# Install razor binary
go install razor

# Ask for provider

# Create .razor dir if not there
# Set the default dir path to home/.razor/
