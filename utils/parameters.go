package utils

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"razor/core"
)

// TODO: Move this away from parameters.go
func GetEpoch(client *ethclient.Client, address string) (uint32, error) {
	latestHeader, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Error("Error in fetching block: ", err)
		return 0, err
	}
	epoch := latestHeader.Number.Int64() / core.EpochLength
	return uint32(epoch), nil
}
