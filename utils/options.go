package utils

import (
	"context"
	"errors"
	"razor/core/types"
	"razor/rpc"
	"strings"

	"github.com/ethereum/go-ethereum"

	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (*UtilsStruct) GetOptions() bind.CallOpts {
	block, _ := new(big.Int).SetString("", 10)
	return bind.CallOpts{
		Pending:     false,
		BlockNumber: block,
		Context:     context.Background(),
	}
}

func (*UtilsStruct) GetTxnOpts(rpcParameters rpc.RPCParameters, transactionData types.TransactionOptions) (*bind.TransactOpts, error) {
	log.Debug("Getting transaction options...")
	account := transactionData.Account
	if account.AccountManager == nil {
		log.Error("Account Manager in transaction data is not initialised")
		return nil, errors.New("account manager not initialised")
	}
	privateKey, err := account.AccountManager.GetPrivateKey(account.Address, account.Password)
	if err != nil {
		log.Error("Error in fetching private key: ", err)
		return nil, err
	}

	nonce, err := ClientInterface.GetNonceAtWithRetry(rpcParameters, common.HexToAddress(account.Address))
	if err != nil {
		log.Error("Error in fetching nonce: ", err)
		return nil, err
	}

	gasPrice := GasInterface.GetGasPrice(rpcParameters, transactionData.Config)
	txnOpts, err := BindInterface.NewKeyedTransactorWithChainID(privateKey, transactionData.ChainId)
	if err != nil {
		log.Error("Error in getting transactor: ", err)
		return nil, err
	}
	txnOpts.Nonce = big.NewInt(int64(nonce))
	txnOpts.GasPrice = gasPrice
	txnOpts.Value = transactionData.EtherValue

	gasLimit, err := GasInterface.GetGasLimit(rpcParameters, transactionData, txnOpts)
	if err != nil {
		errString := err.Error()
		if ContainsStringFromArray(errString, []string{"500", "501", "502", "503", "504"}) || errString == errors.New("intrinsic gas too low").Error() {
			latestBlock, err := ClientInterface.GetLatestBlockWithRetry(rpcParameters)
			if err != nil {
				log.Error("Error in fetching block: ", err)
				return nil, err
			}

			txnOpts.GasLimit = latestBlock.GasLimit
			log.Debug("Error occurred due to RPC issue, sending block gas limit...")
			log.Debug("Gas Limit: ", txnOpts.GasLimit)
			return txnOpts, nil
		}
		log.Error("Error in getting gas limit: ", err)
	}
	log.Debug("Gas after increment: ", gasLimit)
	txnOpts.GasLimit = gasLimit
	return txnOpts, nil
}

func (*GasStruct) GetGasPrice(rpcParameters rpc.RPCParameters, config types.Configurations) *big.Int {
	var gas *big.Int
	if config.GasPrice != 0 {
		gas = big.NewInt(1).Mul(big.NewInt(int64(config.GasPrice)), big.NewInt(1e9))
	} else {
		gas = big.NewInt(0)
	}
	var err error
	suggestedGasPrice, err := ClientInterface.SuggestGasPriceWithRetry(rpcParameters)
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

func (*GasStruct) GetGasLimit(rpcParameters rpc.RPCParameters, transactionData types.TransactionOptions, txnOpts *bind.TransactOpts) (uint64, error) {
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
		From:     common.HexToAddress(transactionData.Account.Address),
		To:       &contractAddress,
		GasPrice: txnOpts.GasPrice,
		Value:    txnOpts.Value,
		Data:     inputData,
	}
	var gasLimit uint64
	if transactionData.MethodName == "reveal" {
		gasLimit, err = getGasLimitForReveal(rpcParameters)
		if err != nil {
			log.Error("GetGasLimit: Error in getting gasLimit for reveal transaction: ", err)
			return transactionData.Config.GasLimitOverride, err
		}
		log.Debug("Calculated gas limit for reveal: ", gasLimit)
	} else {
		gasLimit, err = ClientInterface.EstimateGasWithRetry(rpcParameters, msg)
		if err != nil {
			log.Error("GetGasLimit: Error in getting gasLimit: ", err)
			//If estimateGas throws an error for a transaction than gasLimit should be picked up from the config
			log.Debugf("As there was an error from estimateGas, taking the gas limit value = %d from config", transactionData.Config.GasLimitOverride)
			return transactionData.Config.GasLimitOverride, nil
		}
		log.Debug("Estimated Gas: ", gasLimit)
	}
	return GasInterface.IncreaseGasLimitValue(rpcParameters, gasLimit, transactionData.Config.GasLimitMultiplier)
}

func (*GasStruct) IncreaseGasLimitValue(rpcParameters rpc.RPCParameters, gasLimit uint64, gasLimitMultiplier float32) (uint64, error) {
	if gasLimit == 0 || gasLimitMultiplier <= 0 {
		return gasLimit, nil
	}
	gasLimitIncremented := float64(gasLimitMultiplier) * float64(gasLimit)
	gasLimit = uint64(gasLimitIncremented)

	latestBlock, err := ClientInterface.GetLatestBlockWithRetry(rpcParameters)
	if err != nil {
		log.Error("Error in fetching block: ", err)
		return 0, err
	}

	if gasLimit > latestBlock.GasLimit {
		return latestBlock.GasLimit, nil
	}

	return gasLimit, nil
}

func getGasLimitForReveal(rpcParameters rpc.RPCParameters) (uint64, error) {
	toAssign, err := UtilsInterface.ToAssign(rpcParameters)
	if err != nil {
		return 0, err
	}

	// Apply the formula: gasLimit = 226864 + n * 85000
	gasLimit := 226864 + (uint64(toAssign) * 85000)
	return gasLimit, nil
}
