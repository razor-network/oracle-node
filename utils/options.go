package utils

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum"
	"razor/accounts"
	"razor/core/types"
	"strings"

	"github.com/ethereum/go-ethereum/ethclient"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func GetOptions() bind.CallOpts {
	block, _ := new(big.Int).SetString("", 10)
	return bind.CallOpts{
		Pending:     false,
		BlockNumber: block,
		Context:     context.Background(),
	}
}

func GetTxnOpts(transactionData types.TransactionOptions, razorUtils RazorUtilsInterface) *bind.TransactOpts {
	defaultPath, err := razorUtils.GetDefaultPath()
	CheckError("Error in fetching default path: ", err)
	privateKey := razorUtils.GetPrivateKey(transactionData.AccountAddress, transactionData.Password, defaultPath, accounts.AccountUtilsInterface)
	if privateKey == nil {
		CheckError("Error in fetching private key: ", errors.New(transactionData.AccountAddress+" not present in razor-go"))
	}
	nonce, err := razorUtils.GetPendingNonceAtWithRetry(transactionData.Client, common.HexToAddress(transactionData.AccountAddress))
	CheckError("Error in fetching pending nonce: ", err)

	gasPrice := razorUtils.getGasPrice(transactionData.Client, transactionData.Config, razorUtils)

	txnOpts, err := razorUtils.NewKeyedTransactorWithChainID(privateKey, transactionData.ChainId)
	CheckError("Error in getting transactor: ", err)
	txnOpts.Nonce = big.NewInt(int64(nonce))
	txnOpts.GasPrice = gasPrice
	txnOpts.Value = transactionData.EtherValue

	gasLimit, err := razorUtils.getGasLimit(transactionData, txnOpts, razorUtils)
	if err != nil {
		log.Error("Error in getting gas limit: ", err)
	}
	log.Debug("Gas after increment: ", gasLimit)
	txnOpts.GasLimit = gasLimit
	return txnOpts
}

func getGasPrice(client *ethclient.Client, config types.Configurations, razorUtils RazorUtilsInterface) *big.Int {
	var gas *big.Int
	if config.GasPrice != 0 {
		gas = big.NewInt(1).Mul(big.NewInt(int64(config.GasPrice)), big.NewInt(1e9))
	} else {
		var err error
		gas, err = razorUtils.SuggestGasPriceWithRetry(client)
		if err != nil {
			log.Fatal(err)
		}
	}
	gasPrice := razorUtils.MultiplyFloatAndBigInt(gas, float64(config.GasMultiplier))
	return gasPrice
}

func getGasLimit(transactionData types.TransactionOptions, txnOpts *bind.TransactOpts, razorUtils RazorUtilsInterface) (uint64, error) {
	if transactionData.MethodName == "" {
		return 0, nil
	}
	parsed, err := razorUtils.parse(strings.NewReader(transactionData.ABI))
	if err != nil {
		log.Error("Error in parsing ABI: ", err)
		return 0, err
	}
	inputData, err := razorUtils.Pack(parsed, transactionData.MethodName, transactionData.Parameters...)
	if err != nil {
		log.Error("Error in calculating inputData: ", err)
		return 0, err
	}
	contractAddress := common.HexToAddress(transactionData.ContractAddress)
	msg := ethereum.CallMsg{
		From:     common.HexToAddress(transactionData.AccountAddress),
		To:       &contractAddress,
		GasPrice: txnOpts.GasPrice,
		Value:    txnOpts.Value,
		Data:     inputData,
	}
	gasLimit, err := razorUtils.EstimateGasWithRetry(transactionData.Client, msg)
	if err != nil {
		return 0, err
	}
	log.Debug("Estimated Gas: ", gasLimit)
	return razorUtils.increaseGasLimitValue(transactionData.Client, gasLimit, transactionData.Config.GasLimitMultiplier, razorUtils)
}

func increaseGasLimitValue(client *ethclient.Client, gasLimit uint64, gasLimitMultiplier float32, razorUtils RazorUtilsInterface) (uint64, error) {
	if gasLimit == 0 || gasLimitMultiplier <= 0 {
		return gasLimit, nil
	}
	gasLimitIncremented := float64(gasLimitMultiplier) * float64(gasLimit)
	gasLimit = uint64(gasLimitIncremented)

	latestBlock, err := razorUtils.GetLatestBlockWithRetry(client)
	if err != nil {
		log.Error("Error in fetching block: ", err)
		return 0, err
	}

	if gasLimit > latestBlock.GasLimit {
		return latestBlock.GasLimit, nil
	}

	return gasLimit, nil
}
