package utils

import (
	"context"
	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/core"
)

func (*UtilsStruct) GetPendingNonceAtWithRetry(client *ethclient.Client, accountAddress common.Address) (uint64, error) {
	var (
		nonce uint64
		err   error
	)
	err = retry.Do(
		func() error {
			nonce, err = Options.PendingNonceAt(client, context.Background(), accountAddress)
			if err != nil {
				log.Error("Error in fetching nonce.... Retrying")
				return err
			}
			return nil
		}, Options.RetryAttempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return nonce, nil
}

func (*UtilsStruct) GetLatestBlockWithRetry(client *ethclient.Client) (*types.Header, error) {
	var (
		latestHeader *types.Header
		err          error
	)
	err = retry.Do(
		func() error {
			latestHeader, err = ClientInterface.HeaderByNumber(client, context.Background(), nil)
			if err != nil {
				log.Error("Error in fetching latest block.... Retrying")
				return err
			}
			return nil
		}, Options.RetryAttempts(core.MaxRetries))
	if err != nil {
		return nil, err
	}
	return latestHeader, nil
}

func (o *UtilsStruct) SuggestGasPriceWithRetry(client *ethclient.Client) (*big.Int, error) {
	var (
		gasPrice *big.Int
		err      error
	)
	err = retry.Do(
		func() error {
			gasPrice, err = Options.SuggestGasPrice(client, context.Background())
			if err != nil {
				log.Error("Error in fetching gas price.... Retrying")
				return err
			}
			return nil
		}, Options.RetryAttempts(3))
	if err != nil {
		return nil, err
	}
	return gasPrice, nil
}

func (*UtilsStruct) EstimateGasWithRetry(client *ethclient.Client, message ethereum.CallMsg) (uint64, error) {
	var (
		gasLimit uint64
		err      error
	)
	err = retry.Do(
		func() error {
			gasLimit, err = Options.EstimateGas(client, context.Background(), message)
			if err != nil {
				log.Error("Error in estimating gas limit.... Retrying")
				return err
			}
			return nil
		}, Options.RetryAttempts(3))
	if err != nil {
		return 0, err
	}
	return gasLimit, nil
}

func (*UtilsStruct) FilterLogsWithRetry(client *ethclient.Client, query ethereum.FilterQuery) ([]types.Log, error) {
	var (
		logs []types.Log
		err  error
	)
	err = retry.Do(
		func() error {
			logs, err = Options.FilterLogs(client, context.Background(), query)
			if err != nil {
				log.Error("Error in fetching logs.... Retrying")
				return err
			}
			return nil
		}, Options.RetryAttempts(core.MaxRetries))
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (*UtilsStruct) BalanceAtWithRetry(client *ethclient.Client, account common.Address) (*big.Int, error) {
	var (
		balance *big.Int
		err     error
	)
	err = retry.Do(
		func() error {
			balance, err = ClientInterface.BalanceAt(client, context.Background(), account, nil)
			if err != nil {
				log.Error("Error in fetching logs.... Retrying")
				return err
			}
			return nil
		}, Options.RetryAttempts(core.MaxRetries))
	if err != nil {
		return nil, err
	}
	return balance, nil
}
