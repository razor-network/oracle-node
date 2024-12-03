// Package cmd provides all functions related to command line
package cmd

import (
	"crypto/ecdsa"
	"errors"
	"math/big"
	"os"
	"razor/core"
	"razor/core/types"
	"razor/path"
	"razor/pkg/bindings"
	"razor/rpc"
	"razor/utils"
	"strconv"
	"time"

	"github.com/avast/retry-go"
	ethAccounts "github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	razorUtils  = utils.UtilsInterface
	pathUtils   = path.PathUtilsInterface
	clientUtils = utils.ClientInterface
	fileUtils   = utils.FileInterface
	gasUtils    = utils.GasInterface
	merkleUtils = utils.MerkleInterface
)

//This function initializes the utils
func InitializeUtils() {
	razorUtils = &utils.UtilsStruct{}
	utils.UtilsInterface = &utils.UtilsStruct{}
	utils.EthClient = &utils.EthClientStruct{}
	utils.ClientInterface = &utils.ClientStruct{}
	utils.Time = &utils.TimeStruct{}
	utils.OS = &utils.OSStruct{}
	utils.CoinInterface = &utils.CoinStruct{}
	utils.MerkleInterface = &utils.MerkleTreeStruct{}
	utils.IOInterface = &utils.IOStruct{}
	utils.ABIInterface = &utils.ABIStruct{}
	utils.PathInterface = &utils.PathStruct{}
	utils.BindInterface = &utils.BindStruct{}
	utils.BlockManagerInterface = &utils.BlockManagerStruct{}
	utils.StakeManagerInterface = &utils.StakeManagerStruct{}
	utils.AssetManagerInterface = &utils.AssetManagerStruct{}
	utils.VoteManagerInterface = &utils.VoteManagerStruct{}
	utils.BindingsInterface = &utils.BindingsStruct{}
	utils.JsonInterface = &utils.JsonStruct{}
	utils.StakedTokenInterface = &utils.StakedTokenStruct{}
	utils.RetryInterface = &utils.RetryStruct{}
	utils.MerkleInterface = &utils.MerkleTreeStruct{}
	utils.FlagSetInterface = &utils.FlagSetStruct{}
	clientUtils = &utils.ClientStruct{}
	utils.ClientInterface = &utils.ClientStruct{}
	fileUtils = &utils.FileStruct{}
	utils.FileInterface = &utils.FileStruct{}
	gasUtils = &utils.GasStruct{}
	utils.GasInterface = &utils.GasStruct{}
	merkleUtils = &utils.MerkleTreeStruct{}
	utils.MerkleInterface = &utils.MerkleTreeStruct{}
}

func ExecuteTransaction(interfaceName interface{}, methodName string, args ...interface{}) (*Types.Transaction, error) {
	returnedValues := utils.InvokeFunctionWithTimeout(interfaceName, methodName, args...)
	returnedError := utils.CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return nil, returnedError
	}
	return returnedValues[0].Interface().(*Types.Transaction), nil
}

// FetchFlagInput fetches input value of the flag with given data type and specified flag keyword
func (flagSetUtils FLagSetUtils) FetchFlagInput(flagSet *pflag.FlagSet, flagName string, dataType string) (interface{}, error) {
	switch dataType {
	case "string":
		return flagSet.GetString(flagName)
	case "float32":
		return flagSet.GetFloat32(flagName)
	case "int32":
		return flagSet.GetInt32(flagName)
	case "int64":
		return flagSet.GetInt64(flagName)
	case "uint64":
		return flagSet.GetUint64(flagName)
	case "int":
		return flagSet.GetInt(flagName)
	case "bool":
		return flagSet.GetBool(flagName)
	default:
		return nil, errors.New("unsupported data type for flag input")
	}
}

// FetchRootFlagInput fetches input value of the root flag with given data type and specified flag keyword
func (flagSetUtils FLagSetUtils) FetchRootFlagInput(flagName string, dataType string) (interface{}, error) {
	switch dataType {
	case "string":
		return rootCmd.PersistentFlags().GetString(flagName)
	case "float32":
		return rootCmd.PersistentFlags().GetFloat32(flagName)
	case "int32":
		return rootCmd.PersistentFlags().GetInt32(flagName)
	case "int64":
		return rootCmd.PersistentFlags().GetInt64(flagName)
	case "uint64":
		return rootCmd.PersistentFlags().GetUint64(flagName)
	case "int":
		return rootCmd.PersistentFlags().GetInt(flagName)
	case "bool":
		return rootCmd.PersistentFlags().GetBool(flagName)
	default:
		return nil, errors.New("unsupported data type for root flag input")
	}
}

// Changed returns true if flag was passed in the command else returns false
func (flagSetUtils FLagSetUtils) Changed(flagSet *pflag.FlagSet, flagName string) bool {
	return flagSet.Changed(flagName)
}

//This function returns the hash
func (transactionUtils TransactionUtils) Hash(txn *Types.Transaction) common.Hash {
	return txn.Hash()
}

//This function is of staking the razors
func (stakeManagerUtils StakeManagerUtils) Stake(client *ethclient.Client, txnOpts *bind.TransactOpts, epoch uint32, amount *big.Int) (*Types.Transaction, error) {
	stakeManager := razorUtils.GetStakeManager(client)
	return ExecuteTransaction(stakeManager, "Stake", txnOpts, epoch, amount)
}

//This function resets the unstake lock
func (stakeManagerUtils StakeManagerUtils) ResetUnstakeLock(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32) (*Types.Transaction, error) {
	stakeManager := razorUtils.GetStakeManager(client)
	return ExecuteTransaction(stakeManager, "ResetUnstakeLock", opts, stakerId)
}

//This function is for delegation
func (stakeManagerUtils StakeManagerUtils) Delegate(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32, amount *big.Int) (*Types.Transaction, error) {
	stakeManager := razorUtils.GetStakeManager(client)
	return ExecuteTransaction(stakeManager, "Delegate", opts, stakerId, amount)
}

//This function initiates the withdraw
func (stakeManagerUtils StakeManagerUtils) InitiateWithdraw(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32) (*Types.Transaction, error) {
	stakeManager := razorUtils.GetStakeManager(client)
	return ExecuteTransaction(stakeManager, "InitiateWithdraw", opts, stakerId)
}

//This function unlocks the withdraw amount
func (stakeManagerUtils StakeManagerUtils) UnlockWithdraw(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32) (*Types.Transaction, error) {
	stakeManager := razorUtils.GetStakeManager(client)
	return ExecuteTransaction(stakeManager, "UnlockWithdraw", opts, stakerId)
}

//This function sets the delegation acceptance or rejection
func (stakeManagerUtils StakeManagerUtils) SetDelegationAcceptance(client *ethclient.Client, opts *bind.TransactOpts, status bool) (*Types.Transaction, error) {
	stakeManager := razorUtils.GetStakeManager(client)
	return ExecuteTransaction(stakeManager, "SetDelegationAcceptance", opts, status)
}

//This function updates the commission
func (stakeManagerUtils StakeManagerUtils) UpdateCommission(client *ethclient.Client, opts *bind.TransactOpts, commission uint8) (*Types.Transaction, error) {
	stakeManager := razorUtils.GetStakeManager(client)
	return ExecuteTransaction(stakeManager, "UpdateCommission", opts, commission)
}

//This function allows to unstake the razors
func (stakeManagerUtils StakeManagerUtils) Unstake(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32, sAmount *big.Int) (*Types.Transaction, error) {
	stakeManager := razorUtils.GetStakeManager(client)
	return ExecuteTransaction(stakeManager, "Unstake", opts, stakerId, sAmount)
}

//This function approves the unstake your razor
func (stakeManagerUtils StakeManagerUtils) ApproveUnstake(client *ethclient.Client, opts *bind.TransactOpts, stakerTokenAddress common.Address, amount *big.Int) (*Types.Transaction, error) {
	stakedToken := razorUtils.GetStakedToken(client, stakerTokenAddress)
	log.Debugf("ApproveUnstake: Executing Approve transaction for stakedToken address: %s with arguments amount : %s", stakerTokenAddress, amount)
	return ExecuteTransaction(stakedToken, "Approve", opts, common.HexToAddress(core.StakeManagerAddress), amount)
}

//This function is used to redeem the bounty
func (stakeManagerUtils StakeManagerUtils) RedeemBounty(client *ethclient.Client, opts *bind.TransactOpts, bountyId uint32) (*Types.Transaction, error) {
	stakeManager := razorUtils.GetStakeManager(client)
	return ExecuteTransaction(stakeManager, "RedeemBounty", opts, bountyId)
}

//This function is used to claim the staker reward
func (stakeManagerUtils StakeManagerUtils) ClaimStakerReward(client *ethclient.Client, opts *bind.TransactOpts) (*Types.Transaction, error) {
	stakeManager := razorUtils.GetStakeManager(client)
	return ExecuteTransaction(stakeManager, "ClaimStakerReward", opts)
}

//This function is used to claim the block reward
func (blockManagerUtils BlockManagerUtils) ClaimBlockReward(client *ethclient.Client, opts *bind.TransactOpts) (*Types.Transaction, error) {
	blockManager := razorUtils.GetBlockManager(client)
	return ExecuteTransaction(blockManager, "ClaimBlockReward", opts)
}

// Thid function is used to finalize the dispute
func (blockManagerUtils BlockManagerUtils) FinalizeDispute(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, blockIndex uint8, positionOfCollectionInBlock *big.Int) (*Types.Transaction, error) {
	blockManager := razorUtils.GetBlockManager(client)
	var (
		txn *Types.Transaction
		err error
	)
	err = retry.Do(func() error {
		txn, err = ExecuteTransaction(blockManager, "FinalizeDispute", opts, epoch, blockIndex, positionOfCollectionInBlock)
		if err != nil {
			log.Error("Error in finalizing dispute.. Retrying")
			return err
		}
		return nil
	}, retry.Attempts(3))
	if err != nil {
		return nil, err
	}
	return txn, nil
}

//This function is used to dispute the biggest staker which is proposed
func (blockManagerUtils BlockManagerUtils) DisputeBiggestStakeProposed(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, blockIndex uint8, correctBiggestStakerId uint32) (*Types.Transaction, error) {
	blockManager := razorUtils.GetBlockManager(client)
	var (
		txn *Types.Transaction
		err error
	)
	err = retry.Do(func() error {
		txn, err = ExecuteTransaction(blockManager, "DisputeBiggestStakeProposed", opts, epoch, blockIndex, correctBiggestStakerId)
		if err != nil {
			log.Error("Error in disputing biggest influence proposed.. Retrying")
			return err
		}
		return nil
	}, retry.Attempts(3))
	if err != nil {
		return nil, err
	}
	return txn, nil
}

//This function is used to check if dispute collection Id is absent or not
func (blockManagerUtils BlockManagerUtils) DisputeCollectionIdShouldBeAbsent(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, blockIndex uint8, id uint16, positionOfCollectionInBlock *big.Int) (*Types.Transaction, error) {
	blockManager := razorUtils.GetBlockManager(client)
	var (
		txn *Types.Transaction
		err error
	)
	err = retry.Do(func() error {
		txn, err = ExecuteTransaction(blockManager, "DisputeCollectionIdShouldBeAbsent", opts, epoch, blockIndex, id, positionOfCollectionInBlock)

		if err != nil {
			log.Error("Error in disputing collection id should be absent... Retrying")
			return err
		}
		return nil
	}, retry.Attempts(3))
	if err != nil {
		return nil, err
	}
	return txn, nil
}

//This function is used to check if dispute collection Id is present or not
func (blockManagerUtils BlockManagerUtils) DisputeCollectionIdShouldBePresent(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, blockIndex uint8, id uint16) (*Types.Transaction, error) {
	blockManager := razorUtils.GetBlockManager(client)
	var (
		txn *Types.Transaction
		err error
	)
	err = retry.Do(func() error {
		txn, err = ExecuteTransaction(blockManager, "DisputeCollectionIdShouldBePresent", opts, epoch, blockIndex, id)
		if err != nil {
			log.Error("Error in disputing collection id should be present... Retrying")
			return err
		}
		return nil
	}, retry.Attempts(3))
	if err != nil {
		return nil, err
	}
	return txn, nil
}

//This function is used to do dispute on order of Ids
func (blockManagerUtils BlockManagerUtils) DisputeOnOrderOfIds(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, blockIndex uint8, index0 *big.Int, index1 *big.Int) (*Types.Transaction, error) {
	blockManager := razorUtils.GetBlockManager(client)
	var (
		txn *Types.Transaction
		err error
	)
	err = retry.Do(func() error {
		txn, err = ExecuteTransaction(blockManager, "DisputeOnOrderOfIds", opts, epoch, blockIndex, index0, index1)
		if err != nil {
			log.Error("Error in disputing order of ids proposed... Retrying")
			return err
		}
		return nil
	}, retry.Attempts(3))
	if err != nil {
		return nil, err
	}
	return txn, nil
}

//This function is used for proposing the block
func (blockManagerUtils BlockManagerUtils) Propose(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, ids []uint16, medians []*big.Int, iteration *big.Int, biggestInfluencerId uint32) (*Types.Transaction, error) {
	blockManager := razorUtils.GetBlockManager(client)
	var (
		txn *Types.Transaction
		err error
	)
	err = retry.Do(func() error {
		txn, err = ExecuteTransaction(blockManager, "Propose", opts, epoch, ids, medians, iteration, biggestInfluencerId)
		if err != nil {
			log.Error("Error in proposing... Retrying")
			return err
		}
		return nil
	}, retry.Attempts(3))
	if err != nil {
		return nil, err
	}
	return txn, nil
}

//This function returns the sorted Ids
func (blockManagerUtils BlockManagerUtils) GiveSorted(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, leafId uint16, sortedValues []*big.Int) (*Types.Transaction, error) {
	blockManager := razorUtils.GetBlockManager(client)
	return ExecuteTransaction(blockManager, "GiveSorted", opts, epoch, leafId, sortedValues)
}

//This function resets the dispute
func (blockManagerUtils BlockManagerUtils) ResetDispute(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32) (*Types.Transaction, error) {
	blockManager := razorUtils.GetBlockManager(client)
	return ExecuteTransaction(blockManager, "ResetDispute", opts, epoch)
}

//This function is used to reveal the values
func (voteManagerUtils VoteManagerUtils) Reveal(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, tree bindings.StructsMerkleTree, signature []byte) (*Types.Transaction, error) {
	voteManager := razorUtils.GetVoteManager(client)
	var (
		txn *Types.Transaction
		err error
	)
	err = retry.Do(func() error {
		txn, err = ExecuteTransaction(voteManager, "Reveal", opts, epoch, tree, signature)
		if err != nil {
			log.Error("Error in revealing... Retrying")
			return err
		}
		return nil
	}, retry.Attempts(3))
	if err != nil {
		return nil, err
	}
	return txn, nil
}

//This function is used to commit the values
func (voteManagerUtils VoteManagerUtils) Commit(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, commitment [32]byte) (*Types.Transaction, error) {
	voteManager := razorUtils.GetVoteManager(client)
	var (
		txn *Types.Transaction
		err error
	)
	err = retry.Do(func() error {
		txn, err = ExecuteTransaction(voteManager, "Commit", opts, epoch, commitment)
		if err != nil {
			log.Error("Error in committing... Retrying")
			return err
		}
		return nil
	}, retry.Attempts(3))
	if err != nil {
		return nil, err
	}
	return txn, nil
}

//This function is used to approve the transaction
func (tokenManagerUtils TokenManagerUtils) Approve(client *ethclient.Client, opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*Types.Transaction, error) {
	tokenManager := razorUtils.GetTokenManager(client)
	return ExecuteTransaction(tokenManager, "Approve", opts, spender, amount)
}

//This function is used to transfer the tokens
func (tokenManagerUtils TokenManagerUtils) Transfer(client *ethclient.Client, opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*Types.Transaction, error) {
	tokenManager := razorUtils.GetTokenManager(client)
	return ExecuteTransaction(tokenManager, "Transfer", opts, recipient, amount)
}

//This function is used to create the job
func (assetManagerUtils AssetManagerUtils) CreateJob(client *ethclient.Client, opts *bind.TransactOpts, weight uint8, power int8, selectorType uint8, name string, selector string, url string) (*Types.Transaction, error) {
	assetManager := razorUtils.GetCollectionManager(client)
	return ExecuteTransaction(assetManager, "CreateJob", opts, weight, power, selectorType, name, selector, url)
}

//This function is used to set the collection status
func (assetManagerUtils AssetManagerUtils) SetCollectionStatus(client *ethclient.Client, opts *bind.TransactOpts, assetStatus bool, id uint16) (*Types.Transaction, error) {
	assetManager := razorUtils.GetCollectionManager(client)
	return ExecuteTransaction(assetManager, "SetCollectionStatus", opts, assetStatus, id)
}

//This function is used to get the active status
func (assetManagerUtils AssetManagerUtils) GetActiveStatus(client *ethclient.Client, opts *bind.CallOpts, id uint16) (bool, error) {
	assetMananger := razorUtils.GetCollectionManager(client)
	returnedValues := utils.InvokeFunctionWithTimeout(assetMananger, "GetCollectionStatus", opts, id)
	returnedError := utils.CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return false, returnedError
	}
	return returnedValues[0].Interface().(bool), nil
}

//This function is used to update the job
func (assetManagerUtils AssetManagerUtils) UpdateJob(client *ethclient.Client, opts *bind.TransactOpts, jobId uint16, weight uint8, power int8, selectorType uint8, selector string, url string) (*Types.Transaction, error) {
	assetManager := razorUtils.GetCollectionManager(client)
	return ExecuteTransaction(assetManager, "UpdateJob", opts, jobId, weight, power, selectorType, selector, url)
}

//This function is used to create the collection
func (assetManagerUtils AssetManagerUtils) CreateCollection(client *ethclient.Client, opts *bind.TransactOpts, tolerance uint32, power int8, aggregationMethod uint32, jobIDs []uint16, name string) (*Types.Transaction, error) {
	assetManager := razorUtils.GetCollectionManager(client)
	return ExecuteTransaction(assetManager, "CreateCollection", opts, tolerance, power, aggregationMethod, jobIDs, name)
}

//This function is used to update the collection
func (assetManagerUtils AssetManagerUtils) UpdateCollection(client *ethclient.Client, opts *bind.TransactOpts, collectionId uint16, tolerance uint32, aggregationMethod uint32, power int8, jobIds []uint16) (*Types.Transaction, error) {
	assetManager := razorUtils.GetCollectionManager(client)
	return ExecuteTransaction(assetManager, "UpdateCollection", opts, collectionId, tolerance, aggregationMethod, power, jobIds)
}

//This function returns BountyId in Uint32
func (flagSetUtils FLagSetUtils) GetUint32BountyId(flagSet *pflag.FlagSet) (uint32, error) {
	return flagSet.GetUint32("bountyId")
}

//This function returns the from in string
func (flagSetUtils FLagSetUtils) GetStringFrom(flagSet *pflag.FlagSet) (string, error) {
	from, err := flagSet.GetString("from")
	if err != nil {
		return "", err
	}
	return utils.ValidateAddress(from)
}

//This function returns the to in string
func (flagSetUtils FLagSetUtils) GetStringTo(flagSet *pflag.FlagSet) (string, error) {
	to, err := flagSet.GetString("to")
	if err != nil {
		return "", err
	}
	return utils.ValidateAddress(to)
}

//This function returns the address in string
func (flagSetUtils FLagSetUtils) GetStringAddress(flagSet *pflag.FlagSet) (string, error) {
	address, err := flagSet.GetString("address")
	if err != nil {
		return "", err
	}
	return utils.ValidateAddress(address)
}

//This function returns the stakerId in Uint32
func (flagSetUtils FLagSetUtils) GetUint32StakerId(flagSet *pflag.FlagSet) (uint32, error) {
	return flagSet.GetUint32("stakerId")
}

//This function returns the name in string
func (flagSetUtils FLagSetUtils) GetStringName(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("name")
}

//This function returns the Url in string
func (flagSetUtils FLagSetUtils) GetStringUrl(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("url")
}

//This function returns the selector in string
func (flagSetUtils FLagSetUtils) GetStringSelector(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("selector")
}

//This function returns the power in string
func (flagSetUtils FLagSetUtils) GetInt8Power(flagSet *pflag.FlagSet) (int8, error) {
	return flagSet.GetInt8("power")
}

//This function returns the weight in Uint8
func (flagSetUtils FLagSetUtils) GetUint8Weight(flagSet *pflag.FlagSet) (uint8, error) {
	return flagSet.GetUint8("weight")
}

//This function returns the selectorType in Uint8
func (flagSetUtils FLagSetUtils) GetUint8SelectorType(flagSet *pflag.FlagSet) (uint8, error) {
	return flagSet.GetUint8("selectorType")
}

//This function returns the status in string
func (flagSetUtils FLagSetUtils) GetStringStatus(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("status")
}

//This function returns the commission in Uint8
func (flagSetUtils FLagSetUtils) GetUint8Commission(flagSet *pflag.FlagSet) (uint8, error) {
	return flagSet.GetUint8("commission")
}

//This function returns the jobIds in Uint
func (flagSetUtils FLagSetUtils) GetUintSliceJobIds(flagSet *pflag.FlagSet) ([]uint, error) {
	return flagSet.GetUintSlice("jobIds")
}

//This function returns the aggregation in Uint32
func (flagSetUtils FLagSetUtils) GetUint32Aggregation(flagSet *pflag.FlagSet) (uint32, error) {
	return flagSet.GetUint32("aggregation")
}

//This function returns the JobId in Uint16
func (flagSetUtils FLagSetUtils) GetUint16JobId(flagSet *pflag.FlagSet) (uint16, error) {
	return flagSet.GetUint16("jobId")
}

//This function returns the CollectionId in Uint16
func (flagSetUtils FLagSetUtils) GetUint16CollectionId(flagSet *pflag.FlagSet) (uint16, error) {
	return flagSet.GetUint16("collectionId")
}

//This function returns the value in string
func (flagSetUtils FLagSetUtils) GetStringValue(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("value")
}

//This function is used to check if weiRazor is passed or not
func (flagSetUtils FLagSetUtils) GetBoolWeiRazor(flagSet *pflag.FlagSet) (bool, error) {
	return flagSet.GetBool("weiRazor")
}

//This function returns the tolerance in Uint32
func (flagSetUtils FLagSetUtils) GetUint32Tolerance(flagSet *pflag.FlagSet) (uint32, error) {
	return flagSet.GetUint32("tolerance")
}

//This function is used to check if rogue is passed or not
func (flagSetUtils FLagSetUtils) GetBoolRogue(flagSet *pflag.FlagSet) (bool, error) {
	return flagSet.GetBool("rogue")
}

//This function is used to check if rogueMode is passed or not
func (flagSetUtils FLagSetUtils) GetStringSliceRogueMode(flagSet *pflag.FlagSet) ([]string, error) {
	return flagSet.GetStringSlice("rogueMode")
}

//This function is used to check the inputs gor backupNode flag
func (flagSetUtils FLagSetUtils) GetStringSliceBackupNode(flagSet *pflag.FlagSet) ([]string, error) {
	return flagSet.GetStringSlice("backupNode")
}

//This function is used to check if exposeMetrics is passed or not
func (flagSetUtils FLagSetUtils) GetStringExposeMetrics(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("exposeMetrics")
}

//This function is used to check if CertFile  is passed or not
func (flagSetUtils FLagSetUtils) GetStringCertFile(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("certFile")
}

//This function is used to check if CertFile  is passed or not
func (flagSetUtils FLagSetUtils) GetStringCertKey(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("certKey")
}

//This function returns the max size of log file in Int
func (flagSetUtils FLagSetUtils) GetIntLogFileMaxSize(flagSet *pflag.FlagSet) (int, error) {
	return flagSet.GetInt("logFileMaxSize")
}

//This function returns the max number of backups for logFile in Int
func (flagSetUtils FLagSetUtils) GetIntLogFileMaxBackups(flagSet *pflag.FlagSet) (int, error) {
	return flagSet.GetInt("logFileMaxBackups")
}

//This function returns the max nage for logFle in Int
func (flagSetUtils FLagSetUtils) GetIntLogFileMaxAge(flagSet *pflag.FlagSet) (int, error) {
	return flagSet.GetInt("logFileMaxAge")
}

//This function returns the accounts
func (keystoreUtils KeystoreUtils) Accounts(path string) []ethAccounts.Account {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	return ks.Accounts()
}

//This function is used to import the ECDSA
func (keystoreUtils KeystoreUtils) ImportECDSA(path string, priv *ecdsa.PrivateKey, passphrase string) (ethAccounts.Account, error) {
	ks := keystore.NewKeyStore(path, keystore.StandardScryptN, keystore.StandardScryptP)
	return ks.ImportECDSA(priv, passphrase)
}

//This function is used to convert from Hex to ECDSA
func (c CryptoUtils) HexToECDSA(hexKey string) (*ecdsa.PrivateKey, error) {
	return crypto.HexToECDSA(hexKey)
}

//This function is used to give the sorted Ids
func (*UtilsStruct) GiveSorted(rpcParameters rpc.RPCParameters, txnArgs types.TransactionOptions, epoch uint32, assetId uint16, sortedStakers []*big.Int) error {
	return GiveSorted(rpcParameters, txnArgs, epoch, assetId, sortedStakers)
}

//This function is used to write config as
func (v ViperUtils) ViperWriteConfigAs(path string) error {
	return viper.WriteConfigAs(path)
}

//This function is used for sleep
func (t TimeUtils) Sleep(duration time.Duration) {
	utils.Time.Sleep(duration)
}

//This function is used to parse the bool
func (s StringUtils) ParseBool(str string) (bool, error) {
	return strconv.ParseBool(str)
}

//This function is used for unpacking
func (a AbiUtils) Unpack(abi abi.ABI, name string, data []byte) ([]interface{}, error) {
	return abi.Unpack(name, data)
}

//This function is used for exiting the code
func (o OSUtils) Exit(code int) {
	os.Exit(code)
}
