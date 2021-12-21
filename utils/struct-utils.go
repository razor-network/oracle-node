package utils

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"io"
	"math/big"
	"razor/accounts"
	"razor/path"
)

func StartRazor(optionsPackageStruct OptionsPackageStruct) Utils {
	Options = optionsPackageStruct.Options
	AbiUtils = optionsPackageStruct.AbiUtils
	PathUtils = optionsPackageStruct.PathUtils
	AccountUtils = optionsPackageStruct.AccountUtils
	BindUtils = optionsPackageStruct.BindUtils
	UtilsInterface = optionsPackageStruct.UtilsInterface
	return &UtilsStruct{}
}

func (abiUtils ABIUtilsStruct) Parse(reader io.Reader) (abi.ABI, error) {
	return abi.JSON(reader)
}

func (abiUtils ABIUtilsStruct) Pack(parsedData abi.ABI, name string, args ...interface{}) ([]byte, error) {
	return parsedData.Pack(name, args...)
}

func (pathUtils PathStruct) GetDefaultPath() (string, error) {
	return path.GetDefaultPath()
}

func (account AccountStruct) GetPrivateKey(address string, password string, keystorePath string, accountUtils accounts.AccountInterface) *ecdsa.PrivateKey {
	return accounts.GetPrivateKey(address, password, keystorePath, accountUtils)
}

func (bindUtils BindStruct) NewKeyedTransactorWithChainID(key *ecdsa.PrivateKey, chainID *big.Int) (*bind.TransactOpts, error) {
	return bind.NewKeyedTransactorWithChainID(key, chainID)
}

func (optionUtils *OptionUtilsStruct) GetPendingNonceAtWithRetry(client *ethclient.Client, accountAddress common.Address) (uint64, error) {
	return GetPendingNonceAtWithRetry(client, accountAddress)
}
