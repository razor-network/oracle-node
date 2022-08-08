//Package utils provides the utils functions
package utils

import (
	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"razor/core"
	"razor/pkg/bindings"
)

// GetBondManagerWithOpts function returns the bond manager with opts
func (*UtilsStruct) GetBondManagerWithOpts(client *ethclient.Client) (*bindings.BondManager, bind.CallOpts) {
	return UtilsInterface.GetBondManager(client), UtilsInterface.GetOptions()
}

// GetDataBondCollections returns the collectionIds of data-bond
func (*UtilsStruct) GetDataBondCollections(client *ethclient.Client) ([]uint16, error) {
	var (
		dataBondCollections []uint16
		err                 error
	)
	err = retry.Do(
		func() error {
			dataBondCollections, err = BondManagerInterface.GetDataBondCollections(client)
			if err != nil {
				log.Error("Error in fetching data bonds")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return nil, err
	}
	return dataBondCollections, nil
}
