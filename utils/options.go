//Package utils provides the utils functions
package utils

import (
	"context"
	"errors"
	"path"
	"razor/core/types"
	"strings"

	"github.com/ethereum/go-ethereum"

	"github.com/ethereum/go-ethereum/ethclient"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

//This function returns the options
func (*UtilsStruct) GetOptions() bind.CallOpts {
	block, _ := new(big.Int).SetString("", 10)
	return bind.CallOpts{
		Pending:     false,
		BlockNumber: block,
		Context:     context.Background(),
	}
}

//This function returns the transaction opts
func (*UtilsStruct) GetTxnOpts(transactionData types.TransactionOptions) *bind.TransactOpts {
	defaultPath, err := PathInterface.GetDefaultPath()
	CheckError("Error in fetching default path: ", err)
	keystorePath := path.Join(defaultPath, "keystoreFiles")
	privateKey := AccountsInterface.GetPrivateKey(transactionData.AccountAddress, transactionData.Password, keystorePath)
	if privateKey == nil {
		CheckError("Error in fetching private key: ", errors.New(transactionData.AccountAddress+" not present in razor-go"))
	}
	nonce, err := UtilsInterface.GetPendingNonceAtWithRetry(transactionData.Client, common.HexToAddress(transactionData.AccountAddress))
	CheckError("Error in fetching pending nonce: ", err)

	gasPrice := UtilsInterface.GetGasPrice(transactionData.Client, transactionData.Config)
	txnOpts, err := BindInterface.NewKeyedTransactorWithChainID(privateKey, transactionData.ChainId)
	CheckError("Error in getting transactor: ", err)
	txnOpts.Nonce = big.NewInt(int64(nonce))
	txnOpts.GasPrice = gasPrice
	txnOpts.Value = transactionData.EtherValue

	gasLimit, err := UtilsInterface.GetGasLimit(transactionData, txnOpts)
	if err != nil {
		errString := err.Error()
		if ContainsStringFromArray(errString, []string{"500", "501", "502", "503", "504"}) || errString == errors.New("intrinsic gas too low").Error() {
			latestBlock, err := UtilsInterface.GetLatestBlockWithRetry(transactionData.Client)
			CheckError("Error in fetching block: ", err)

			txnOpts.GasLimit = latestBlock.GasLimit
			log.Debug("Error occurred due to RPC issue, sending block gas limit...")
			log.Debug("Gas Limit: ", txnOpts.GasLimit)
			return txnOpts
		}
		log.Error("Error in getting gas limit: ", err)
	}
	log.Debug("Gas after increment: ", gasLimit)
	txnOpts.GasLimit = gasLimit
	return txnOpts
}

//This function returns the gas price
func (*UtilsStruct) GetGasPrice(client *ethclient.Client, config types.Configurations) *big.Int {
	var gas *big.Int
	if config.GasPrice != 0 {
		gas = big.NewInt(1).Mul(big.NewInt(int64(config.GasPrice)), big.NewInt(1e9))
	} else {
		gas = big.NewInt(0)
	}
	var err error
	suggestedGasPrice, err := UtilsInterface.SuggestGasPriceWithRetry(client)
	if err != nil {
		log.Error(err)
		return UtilsInterface.MultiplyFloatAndBigInt(gas, float64(config.GasMultiplier))
	}
	log.Debugf("Suggested gas price: %d", suggestedGasPrice)
	log.Debugf("Gas Price set in config: %d", gas)
	if suggestedGasPrice.Cmp(gas) > 0 {
		log.Debugf("Going with suggested gas price!")
		gas = suggestedGasPrice
	}
	gasPrice := UtilsInterface.MultiplyFloatAndBigInt(gas, float64(config.GasMultiplier))
	return gasPrice
}

//This function returns the gas limit
func (*UtilsStruct) GetGasLimit(transactionData types.TransactionOptions, txnOpts *bind.TransactOpts) (uint64, error) {
	if transactionData.MethodName == "" {
		return 0, nil
	}
	parsed, err := ABIInterface.Parse(strings.NewReader(transactionData.ABI))
	if err != nil {
		log.Error("Error in parsing abi: ", err)
		return 0, err
	}
	inputData, err := ABIInterface.Pack(parsed, transactionData.MethodName, transactionData.Parameters...)
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
	gasLimit, err := UtilsInterface.EstimateGasWithRetry(transactionData.Client, msg)
	if err != nil {
		return 0, err
	}
	log.Debug("Estimated Gas: ", gasLimit)
	return UtilsInterface.IncreaseGasLimitValue(transactionData.Client, gasLimit, transactionData.Config.GasLimitMultiplier)
}

//This function increases the gas limit value
func (*UtilsStruct) IncreaseGasLimitValue(client *ethclient.Client, gasLimit uint64, gasLimitMultiplier float32) (uint64, error) {
	if gasLimit == 0 || gasLimitMultiplier <= 0 {
		return gasLimit, nil
	}
	gasLimitIncremented := float64(gasLimitMultiplier) * float64(gasLimit)
	gasLimit = uint64(gasLimitIncremented)

	latestBlock, err := UtilsInterface.GetLatestBlockWithRetry(client)
	if err != nil {
		log.Error("Error in fetching block: ", err)
		return 0, err
	}

	if gasLimit > latestBlock.GasLimit {
		return latestBlock.GasLimit, nil
	}

	return gasLimit, nil
}
