//Package cmd provides all functions related to command line
package cmd

import (
	"crypto/ecdsa"
	"math/big"
	"os"
	"razor/core"
	"razor/core/types"
	"razor/path"
	"razor/pkg/bindings"
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

var utilsInterface = utils.UtilsInterface

//This function initializes the utils
func InitializeUtils() {
	utilsInterface = &utils.UtilsStruct{}
	utils.UtilsInterface = &utils.UtilsStruct{}
	utils.EthClient = &utils.EthClientStruct{}
	utils.ClientInterface = &utils.ClientStruct{}
	utils.Time = &utils.TimeStruct{}
	utils.OS = &utils.OSStruct{}
	utils.Bufio = &utils.BufioStruct{}
	utils.CoinInterface = &utils.CoinStruct{}
	utils.MerkleInterface = &utils.MerkleTreeStruct{}
	utils.IOInterface = &utils.IOStruct{}
	utils.ABIInterface = &utils.ABIStruct{}
	utils.PathInterface = &utils.PathStruct{}
	utils.BindInterface = &utils.BindStruct{}
	utils.AccountsInterface = &utils.AccountsStruct{}
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
}

func ExecuteTransaction(interfaceName interface{}, methodName string, args ...interface{}) (*Types.Transaction, error) {
	returnedValues := utils.InvokeFunctionWithTimeout(interfaceName, methodName, args...)
	returnedError := utils.CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return nil, returnedError
	}
	return returnedValues[0].Interface().(*Types.Transaction), nil
}

//This function returns the config file path
func (u Utils) GetConfigFilePath() (string, error) {
	return path.PathUtilsInterface.GetConfigFilePath()
}

//This function returns the epoch
func (u Utils) GetEpoch(client *ethclient.Client) (uint32, error) {
	return utilsInterface.GetEpoch(client)
}

//This function returns the options
func (u Utils) GetOptions() bind.CallOpts {
	return utilsInterface.GetOptions()
}

//This function returns the block time
func (u Utils) CalculateBlockTime(client *ethclient.Client) int64 {
	return utilsInterface.CalculateBlockTime(client)
}

//This function returns the transaction opts
func (u Utils) GetTxnOpts(transactionData types.TransactionOptions) *bind.TransactOpts {
	return utilsInterface.GetTxnOpts(transactionData)
}

//This function returns the config data
func (u Utils) GetConfigData() (types.Configurations, error) {
	return cmdUtils.GetConfigData()
}

//This function assigns the password
func (u Utils) AssignPassword(flagSet *pflag.FlagSet) string {
	return utils.AssignPassword(flagSet)
}

//This function returns the string address
func (u Utils) GetStringAddress(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("address")
}

//This function returns the Uint32 bountyId
func (u Utils) GetUint32BountyId(flagSet *pflag.FlagSet) (uint32, error) {
	return flagSet.GetUint32("bountyId")
}

//This function connects to the client
func (u Utils) ConnectToClient(provider string) *ethclient.Client {
	returnedValues := utils.InvokeFunctionWithTimeout(utilsInterface, "ConnectToClient", provider)
	returnedError := utils.CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return nil
	}
	return returnedValues[0].Interface().(*ethclient.Client)
}

//This function waits for the block completion
func (u Utils) WaitForBlockCompletion(client *ethclient.Client, hashToRead string) error {
	return utilsInterface.WaitForBlockCompletion(client, hashToRead)
}

//This function returns the number of active collections
func (u Utils) GetNumActiveCollections(client *ethclient.Client) (uint16, error) {
	return utilsInterface.GetNumActiveCollections(client)
}

//This function returns rogue random value
func (u Utils) GetRogueRandomValue(value int) *big.Int {
	return utils.GetRogueRandomValue(value)
}

//This function returns the rogue median value
func (u Utils) GetRogueRandomMedianValue() uint32 {
	return utils.GetRogueRandomMedianValue()
}

//This function returns the aggregated data of collection
func (u Utils) GetAggregatedDataOfCollection(client *ethclient.Client, collectionId uint16, epoch uint32) (*big.Int, error) {
	return utilsInterface.GetAggregatedDataOfCollection(client, collectionId, epoch)
}

//This function returns the delayed state
func (u Utils) GetBufferedState(client *ethclient.Client, buffer int32) (int64, error) {
	return utilsInterface.GetBufferedState(client, buffer)
}

//This function returns the default path
func (u Utils) GetDefaultPath() (string, error) {
	return path.PathUtilsInterface.GetDefaultPath()
}

//This function returns the job file path
func (u Utils) GetJobFilePath() (string, error) {
	return path.PathUtilsInterface.GetJobFilePath()
}

//This function fetches the balance
func (u Utils) FetchBalance(client *ethclient.Client, accountAddress string) (*big.Int, error) {
	return utilsInterface.FetchBalance(client, accountAddress)
}

//This function checks if the flag is passed
func (u Utils) IsFlagPassed(name string) bool {
	return utilsInterface.IsFlagPassed(name)
}

//This function returns the amount in wei
func (u Utils) GetAmountInWei(amount *big.Int) *big.Int {
	return utils.GetAmountInWei(amount)
}

//This function checks the amount and balance
func (u Utils) CheckAmountAndBalance(amountInWei *big.Int, balance *big.Int) *big.Int {
	return utils.CheckAmountAndBalance(amountInWei, balance)
}

//This function returns the amount in decimal
func (u Utils) GetAmountInDecimal(amountInWei *big.Int) *big.Float {
	return utils.GetAmountInDecimal(amountInWei)
}

//This function returns the epoch which is last committed
func (u Utils) GetEpochLastCommitted(client *ethclient.Client, stakerId uint32) (uint32, error) {
	return utilsInterface.GetEpochLastCommitted(client, stakerId)
}

//This function returns the commitments
func (u Utils) GetCommitments(client *ethclient.Client, address string) ([32]byte, error) {
	return utilsInterface.GetCommitments(client, address)
}

//This function returns if all the values in bytesValue is zero
func (u Utils) AllZero(bytesValue [32]byte) bool {
	return utils.AllZero(bytesValue)
}

//This function converts the Uint array to Uint16 array
func (u Utils) ConvertUintArrayToUint16Array(uintArr []uint) []uint16 {
	return utils.ConvertUintArrayToUint16Array(uintArr)
}

//This function returns the jobs array
func (u Utils) GetJobs(client *ethclient.Client) ([]bindings.StructsJob, error) {
	return utilsInterface.GetJobs(client)
}

//This function if the eth balance is zero or not
func (u Utils) CheckEthBalanceIsZero(client *ethclient.Client, address string) {
	utilsInterface.CheckEthBalanceIsZero(client, address)
}

//This function assigns the stakerId
func (u Utils) AssignStakerId(flagSet *pflag.FlagSet, client *ethclient.Client, address string) (uint32, error) {
	return utilsInterface.AssignStakerId(flagSet, client, address)
}

//This function returns the lock
func (u Utils) GetLock(client *ethclient.Client, address string, stakerId uint32, lockType uint8) (types.Locks, error) {
	return utilsInterface.GetLock(client, address, stakerId, lockType)
}

//This function returns the staker
func (u Utils) GetStaker(client *ethclient.Client, stakerId uint32) (bindings.StructsStaker, error) {
	return utilsInterface.GetStaker(client, stakerId)
}

//This function returns the updated staker
func (u Utils) GetUpdatedStaker(client *ethclient.Client, stakerId uint32) (bindings.StructsStaker, error) {
	return utilsInterface.GetStaker(client, stakerId)
}

//This function returns the staked token
func (u Utils) GetStakedToken(client *ethclient.Client, address common.Address) *bindings.StakedToken {
	return utilsInterface.GetStakedToken(client, address)
}

//This function converts the SRazor to Razor
func (u Utils) ConvertSRZRToRZR(sAmount *big.Int, currentStake *big.Int, totalSupply *big.Int) *big.Int {
	return utils.ConvertSRZRToRZR(sAmount, currentStake, totalSupply)
}

//This function converts the Razor to SRazors
func (u Utils) ConvertRZRToSRZR(sAmount *big.Int, currentStake *big.Int, totalSupply *big.Int) (*big.Int, error) {
	return utils.ConvertRZRToSRZR(sAmount, currentStake, totalSupply)
}

//This function returns the withdraw initiation period
func (u Utils) GetWithdrawInitiationPeriod(client *ethclient.Client) (uint16, error) {
	return utilsInterface.GetWithdrawInitiationPeriod(client)
}

//This function returns the influence snapshot
func (u Utils) GetInfluenceSnapshot(client *ethclient.Client, stakerId uint32, epoch uint32) (*big.Int, error) {
	return utilsInterface.GetInfluenceSnapshot(client, stakerId, epoch)
}

//This function returns the collections
func (u Utils) GetCollections(client *ethclient.Client) ([]bindings.StructsCollection, error) {
	return utilsInterface.GetAllCollections(client)
}

//This function returns the number of stakers
func (u Utils) GetNumberOfStakers(client *ethclient.Client) (uint32, error) {
	return utilsInterface.GetNumberOfStakers(client)
}

//This function returns the number of proposed blocks
func (u Utils) GetNumberOfProposedBlocks(client *ethclient.Client, epoch uint32) (uint8, error) {
	return utilsInterface.GetNumberOfProposedBlocks(client, epoch)
}

//This function returns the maximum alt blocks
func (u Utils) GetMaxAltBlocks(client *ethclient.Client) (uint8, error) {
	return utilsInterface.GetMaxAltBlocks(client)
}

//This function returns the proposed block
func (u Utils) GetProposedBlock(client *ethclient.Client, epoch uint32, proposedBlockId uint32) (bindings.StructsBlock, error) {
	return utilsInterface.GetProposedBlock(client, epoch, proposedBlockId)
}

//This function returns the epoch which is last revealed
func (u Utils) GetEpochLastRevealed(client *ethclient.Client, stakerId uint32) (uint32, error) {
	return utilsInterface.GetEpochLastRevealed(client, stakerId)
}

//This function returns the epoch which is last proposed
func (u Utils) GetEpochLastProposed(client *ethclient.Client, stakerId uint32) (uint32, error) {
	return utilsInterface.GetEpochLastProposed(client, stakerId)
}

//This function returns the vote value
func (u Utils) GetVoteValue(client *ethclient.Client, epoch uint32, stakerId uint32, medianIndex uint16) (*big.Int, error) {
	return utilsInterface.GetVoteValue(client, epoch, stakerId, medianIndex)
}

//This function returns the total influence revealed
func (u Utils) GetTotalInfluenceRevealed(client *ethclient.Client, epoch uint32, medianIndex uint16) (*big.Int, error) {
	return utilsInterface.GetTotalInfluenceRevealed(client, epoch, medianIndex)
}

//This function returns the Uint32 Array to BigInt array
func (u Utils) ConvertUint32ArrayToBigIntArray(uint32Array []uint32) []*big.Int {
	return utils.ConvertUint32ArrayToBigIntArray(uint32Array)
}

//This function returns the active collections
func (u Utils) GetActiveCollections(client *ethclient.Client) ([]uint16, error) {
	return utilsInterface.GetActiveCollectionIds(client)
}

//This function retrns the block manager
func (u Utils) GetBlockManager(client *ethclient.Client) *bindings.BlockManager {
	return utilsInterface.GetBlockManager(client)
}

//This function returns the sorted proposed block Ids
func (u Utils) GetSortedProposedBlockIds(client *ethclient.Client, epoch uint32) ([]uint32, error) {
	return utilsInterface.GetSortedProposedBlockIds(client, epoch)
}

//This function returns the stakerId
func (u Utils) GetStakerId(client *ethclient.Client, address string) (uint32, error) {
	return utilsInterface.GetStakerId(client, address)
}

//This function returns the stake
func (u Utils) GetStake(client *ethclient.Client, stakerId uint32) (*big.Int, error) {
	return utilsInterface.GetStake(client, stakerId)
}

//This function prompts the private key
func (u Utils) PrivateKeyPrompt() string {
	return utils.PrivateKeyPrompt()
}

//This function prompts the password
func (u Utils) PasswordPrompt() string {
	return utils.PasswordPrompt()
}

//This function returns the max commission
func (u Utils) GetMaxCommission(client *ethclient.Client) (uint8, error) {
	return utilsInterface.GetMaxCommission(client)
}

//This function returns the epoch limit for updated commission
func (u Utils) GetEpochLimitForUpdateCommission(client *ethclient.Client) (uint16, error) {
	return utilsInterface.GetEpochLimitForUpdateCommission(client)
}

//This function returns the stake snapshot
func (u Utils) GetStakeSnapshot(client *ethclient.Client, stakerId uint32, epoch uint32) (*big.Int, error) {
	return utilsInterface.GetStakeSnapshot(client, stakerId, epoch)
}

//This function converts the wei to eth
func (u Utils) ConvertWeiToEth(data *big.Int) (*big.Float, error) {
	return utils.ConvertWeiToEth(data)
}

//This function wait till next N seconds
func (u Utils) WaitTillNextNSecs(seconds int32) {
	utilsInterface.WaitTillNextNSecs(seconds)
}

//This function deletes the job from JSON
func (u Utils) DeleteJobFromJSON(s string, jobId string) error {
	return utilsInterface.DeleteJobFromJSON(s, jobId)
}

//This function adds the job to JSON
func (u Utils) AddJobToJSON(s string, job *types.StructsJob) error {
	return utilsInterface.AddJobToJSON(s, job)
}

//This function converts seconds into readable time
func (u Utils) SecondsToReadableTime(time int) string {
	return utilsInterface.SecondsToReadableTime(time)
}

//This function returns the staker SRZR balance
func (u Utils) GetStakerSRZRBalance(client *ethclient.Client, staker bindings.StructsStaker) (*big.Int, error) {
	return utilsInterface.GetStakerSRZRBalance(client, staker)
}

//This function saves the data to commit JSON File
func (u Utils) SaveDataToCommitJsonFile(flePath string, epoch uint32, commitFileData types.CommitData) error {
	return utilsInterface.SaveDataToCommitJsonFile(flePath, epoch, commitFileData)
}

//This function reads from the commit JSON file
func (u Utils) ReadFromCommitJsonFile(filePath string) (types.CommitFileData, error) {
	return utilsInterface.ReadFromCommitJsonFile(filePath)
}

//This function assigns the log file
func (u Utils) AssignLogFile(flagSet *pflag.FlagSet) {
	utilsInterface.AssignLogFile(flagSet)
}

//This function reads from propose JSON file
func (u Utils) ReadFromProposeJsonFile(filePath string) (types.ProposeFileData, error) {
	return utilsInterface.ReadFromProposeJsonFile(filePath)
}

//This function saves the data to propose JSON file
func (u Utils) SaveDataToProposeJsonFile(flePath string, proposeFileData types.ProposeFileData) error {
	return utilsInterface.SaveDataToProposeJsonFile(flePath, proposeFileData)
}

//This function saves data to Dispute JSON file
func (u Utils) SaveDataToDisputeJsonFile(filePath string, bountyIdQueue []uint32) error {
	return utilsInterface.SaveDataToDisputeJsonFile(filePath, bountyIdQueue)
}

//This function reads from Dispute JSON file
func (u Utils) ReadFromDisputeJsonFile(filePath string) (types.DisputeFileData, error) {
	return utilsInterface.ReadFromDisputeJsonFile(filePath)
}

//This function returns the proposed data JSON file
func (u Utils) GetProposeDataFileName(address string) (string, error) {
	return path.PathUtilsInterface.GetProposeDataFileName(address)
}

//This function returns the commit data file name
func (u Utils) GetCommitDataFileName(address string) (string, error) {
	return path.PathUtilsInterface.GetCommitDataFileName(address)
}

//This function returns the dispute data file name
func (u Utils) GetDisputeDataFileName(address string) (string, error) {
	return path.PathUtilsInterface.GetDisputeDataFileName(address)
}

//This function returns the hash
func (transactionUtils TransactionUtils) Hash(txn *Types.Transaction) common.Hash {
	return txn.Hash()
}

//This function is of staking the razors
func (stakeManagerUtils StakeManagerUtils) Stake(client *ethclient.Client, txnOpts *bind.TransactOpts, epoch uint32, amount *big.Int) (*Types.Transaction, error) {
	stakeManager := utilsInterface.GetStakeManager(client)
	return ExecuteTransaction(stakeManager, "Stake", txnOpts, epoch, amount)
}

//This function resets the unstake lock
func (stakeManagerUtils StakeManagerUtils) ResetUnstakeLock(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32) (*Types.Transaction, error) {
	stakeManager := utilsInterface.GetStakeManager(client)
	return ExecuteTransaction(stakeManager, "ResetUnstakeLock", opts, stakerId)
}

//This function is for delegation
func (stakeManagerUtils StakeManagerUtils) Delegate(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32, amount *big.Int) (*Types.Transaction, error) {
	stakeManager := utilsInterface.GetStakeManager(client)
	return ExecuteTransaction(stakeManager, "Delegate", opts, stakerId, amount)
}

//This function initiates the withdraw
func (stakeManagerUtils StakeManagerUtils) InitiateWithdraw(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32) (*Types.Transaction, error) {
	stakeManager := utilsInterface.GetStakeManager(client)
	return ExecuteTransaction(stakeManager, "InitiateWithdraw", opts, stakerId)
}

//This function unlocks the withdraw amount
func (stakeManagerUtils StakeManagerUtils) UnlockWithdraw(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32) (*Types.Transaction, error) {
	stakeManager := utilsInterface.GetStakeManager(client)
	return ExecuteTransaction(stakeManager, "UnlockWithdraw", opts, stakerId)
}

//This function sets the delegation acceptance or rejection
func (stakeManagerUtils StakeManagerUtils) SetDelegationAcceptance(client *ethclient.Client, opts *bind.TransactOpts, status bool) (*Types.Transaction, error) {
	stakeManager := utilsInterface.GetStakeManager(client)
	return ExecuteTransaction(stakeManager, "SetDelegationAcceptance", opts, status)
}

//This function updates the commission
func (stakeManagerUtils StakeManagerUtils) UpdateCommission(client *ethclient.Client, opts *bind.TransactOpts, commission uint8) (*Types.Transaction, error) {
	stakeManager := utilsInterface.GetStakeManager(client)
	return ExecuteTransaction(stakeManager, "UpdateCommission", opts, commission)
}

//This function allows to unstake the razors
func (stakeManagerUtils StakeManagerUtils) Unstake(client *ethclient.Client, opts *bind.TransactOpts, stakerId uint32, sAmount *big.Int) (*Types.Transaction, error) {
	stakeManager := utilsInterface.GetStakeManager(client)
	return ExecuteTransaction(stakeManager, "Unstake", opts, stakerId, sAmount)
}

//This function approves the unstake your razor
func (stakeManagerUtils StakeManagerUtils) ApproveUnstake(client *ethclient.Client, opts *bind.TransactOpts, staker bindings.StructsStaker, amount *big.Int) (*Types.Transaction, error) {
	stakedToken := razorUtils.GetStakedToken(client, staker.TokenAddress)
	return ExecuteTransaction(stakedToken, "Approve", opts, common.HexToAddress(core.StakeManagerAddress), amount)
}

//This function is used to redeem the bounty
func (stakeManagerUtils StakeManagerUtils) RedeemBounty(client *ethclient.Client, opts *bind.TransactOpts, bountyId uint32) (*Types.Transaction, error) {
	stakeManager := utilsInterface.GetStakeManager(client)
	return ExecuteTransaction(stakeManager, "RedeemBounty", opts, bountyId)
}

//This function returns the staker Info
func (stakeManagerUtils StakeManagerUtils) StakerInfo(client *ethclient.Client, opts *bind.CallOpts, stakerId uint32) (types.Staker, error) {
	stakeManager := utilsInterface.GetStakeManager(client)
	returnedValues := utils.InvokeFunctionWithTimeout(stakeManager, "Stakers", opts, stakerId)
	returnedError := utils.CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return types.Staker{}, returnedError
	}
	staker := returnedValues[0].Interface().(struct {
		AcceptDelegation                bool
		IsSlashed                       bool
		Commission                      uint8
		Id                              uint32
		Age                             uint32
		Address                         common.Address
		TokenAddress                    common.Address
		EpochFirstStakedOrLastPenalized uint32
		EpochCommissionLastUpdated      uint32
		Stake                           *big.Int
		StakerReward                    *big.Int
	})
	return staker, nil
}

//This function returns the maturity
func (stakeManagerUtils StakeManagerUtils) GetMaturity(client *ethclient.Client, opts *bind.CallOpts, age uint32) (uint16, error) {
	stakeManager := utilsInterface.GetStakeManager(client)
	index := age / 10000
	returnedValues := utils.InvokeFunctionWithTimeout(stakeManager, "Maturities", opts, big.NewInt(int64(index)))
	returnedError := utils.CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return 0, returnedError
	}
	return returnedValues[0].Interface().(uint16), nil
}

//This function returns the bounty lock
func (stakeManagerUtils StakeManagerUtils) GetBountyLock(client *ethclient.Client, opts *bind.CallOpts, bountyId uint32) (types.BountyLock, error) {
	stakeManager := utilsInterface.GetStakeManager(client)
	returnedValues := utils.InvokeFunctionWithTimeout(stakeManager, "BountyLocks", opts, bountyId)
	returnedError := utils.CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return types.BountyLock{}, returnedError
	}
	bountyLock := returnedValues[0].Interface().(struct {
		RedeemAfter  uint32
		BountyHunter common.Address
		Amount       *big.Int
	})
	return bountyLock, nil
}

//This function is used to claim the staker reward
func (stakeManagerUtils StakeManagerUtils) ClaimStakerReward(client *ethclient.Client, opts *bind.TransactOpts) (*Types.Transaction, error) {
	stakeManager := utilsInterface.GetStakeManager(client)
	return ExecuteTransaction(stakeManager, "ClaimStakerReward", opts)
}

//This function is used to claim the block reward
func (blockManagerUtils BlockManagerUtils) ClaimBlockReward(client *ethclient.Client, opts *bind.TransactOpts) (*Types.Transaction, error) {
	blockManager := utilsInterface.GetBlockManager(client)
	return ExecuteTransaction(blockManager, "ClaimBlockReward", opts)
}

//Thid function is used to finalize the dispute
func (blockManagerUtils BlockManagerUtils) FinalizeDispute(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, blockIndex uint8, positionOfCollectionInBlock *big.Int) (*Types.Transaction, error) {
	blockManager := utilsInterface.GetBlockManager(client)
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
	blockManager := utilsInterface.GetBlockManager(client)
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
	blockManager := utilsInterface.GetBlockManager(client)
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
	blockManager := utilsInterface.GetBlockManager(client)
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
	blockManager := utilsInterface.GetBlockManager(client)
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
	blockManager := utilsInterface.GetBlockManager(client)
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
func (blockManagerUtils BlockManagerUtils) GiveSorted(blockManager *bindings.BlockManager, opts *bind.TransactOpts, epoch uint32, leafId uint16, sortedValues []*big.Int) (*Types.Transaction, error) {
	return ExecuteTransaction(blockManager, "GiveSorted", opts, epoch, leafId, sortedValues)
}

//This function resets the dispute
func (blockManagerUtils BlockManagerUtils) ResetDispute(blockManager *bindings.BlockManager, opts *bind.TransactOpts, epoch uint32) (*Types.Transaction, error) {
	return ExecuteTransaction(blockManager, "ResetDispute", opts, epoch)
}

//This functiom gets Disputes mapping
func (blockManagerUtils BlockManagerUtils) Disputes(client *ethclient.Client, opts *bind.CallOpts, epoch uint32, address common.Address) (types.DisputesStruct, error) {
	blockManager := utilsInterface.GetBlockManager(client)
	returnedValues := utils.InvokeFunctionWithTimeout(blockManager, "Disputes", opts, epoch, address)
	returnedError := utils.CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return types.DisputesStruct{}, returnedError
	}
	disputesMapping := returnedValues[0].Interface().(struct {
		LeafId           uint16
		LastVisitedValue *big.Int
		AccWeight        *big.Int
		Median           *big.Int
	})
	return disputesMapping, nil
}

//This function is used to reveal the values
func (voteManagerUtils VoteManagerUtils) Reveal(client *ethclient.Client, opts *bind.TransactOpts, epoch uint32, tree bindings.StructsMerkleTree, signature []byte) (*Types.Transaction, error) {
	voteManager := utilsInterface.GetVoteManager(client)
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
	voteManager := utilsInterface.GetVoteManager(client)
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

//This function is used to check the allowance of staker
func (tokenManagerUtils TokenManagerUtils) Allowance(client *ethclient.Client, opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	tokenManager := utilsInterface.GetTokenManager(client)
	returnedValues := utils.InvokeFunctionWithTimeout(tokenManager, "Allowance", opts, owner, spender)
	returnedError := utils.CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return nil, returnedError
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

//This function is used to approve the transaction
func (tokenManagerUtils TokenManagerUtils) Approve(client *ethclient.Client, opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*Types.Transaction, error) {
	tokenManager := utilsInterface.GetTokenManager(client)
	return ExecuteTransaction(tokenManager, "Approve", opts, spender, amount)
}

//This function is used to transfer the tokens
func (tokenManagerUtils TokenManagerUtils) Transfer(client *ethclient.Client, opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*Types.Transaction, error) {
	tokenManager := utilsInterface.GetTokenManager(client)
	return ExecuteTransaction(tokenManager, "Transfer", opts, recipient, amount)
}

//This function is used to create the job
func (assetManagerUtils AssetManagerUtils) CreateJob(client *ethclient.Client, opts *bind.TransactOpts, weight uint8, power int8, selectorType uint8, name string, selector string, url string) (*Types.Transaction, error) {
	assetManager := utilsInterface.GetCollectionManager(client)
	return ExecuteTransaction(assetManager, "CreateJob", opts, weight, power, selectorType, name, selector, url)
}

//This function is used to set the collection status
func (assetManagerUtils AssetManagerUtils) SetCollectionStatus(client *ethclient.Client, opts *bind.TransactOpts, assetStatus bool, id uint16) (*Types.Transaction, error) {
	assetManager := utilsInterface.GetCollectionManager(client)
	return ExecuteTransaction(assetManager, "SetCollectionStatus", opts, assetStatus, id)
}

//This function is used to get the active status
func (assetManagerUtils AssetManagerUtils) GetActiveStatus(client *ethclient.Client, opts *bind.CallOpts, id uint16) (bool, error) {
	assetMananger := utilsInterface.GetCollectionManager(client)
	returnedValues := utils.InvokeFunctionWithTimeout(assetMananger, "GetCollectionStatus", opts, id)
	returnedError := utils.CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return false, returnedError
	}
	return returnedValues[0].Interface().(bool), nil
}

//This function is used to update the job
func (assetManagerUtils AssetManagerUtils) UpdateJob(client *ethclient.Client, opts *bind.TransactOpts, jobId uint16, weight uint8, power int8, selectorType uint8, selector string, url string) (*Types.Transaction, error) {
	assetManager := utilsInterface.GetCollectionManager(client)
	return ExecuteTransaction(assetManager, "UpdateJob", opts, jobId, weight, power, selectorType, selector, url)
}

//This function is used to create the collection
func (assetManagerUtils AssetManagerUtils) CreateCollection(client *ethclient.Client, opts *bind.TransactOpts, tolerance uint32, power int8, aggregationMethod uint32, jobIDs []uint16, name string) (*Types.Transaction, error) {
	assetManager := utilsInterface.GetCollectionManager(client)
	return ExecuteTransaction(assetManager, "CreateCollection", opts, tolerance, power, aggregationMethod, jobIDs, name)
}

//This function is used to update the collection
func (assetManagerUtils AssetManagerUtils) UpdateCollection(client *ethclient.Client, opts *bind.TransactOpts, collectionId uint16, tolerance uint32, aggregationMethod uint32, power int8, jobIds []uint16) (*Types.Transaction, error) {
	assetManager := utilsInterface.GetCollectionManager(client)
	return ExecuteTransaction(assetManager, "UpdateCollection", opts, collectionId, tolerance, aggregationMethod, power, jobIds)
}

//This function returns the provider in string
func (flagSetUtils FLagSetUtils) GetStringProvider(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("provider")
}

//This function returns gas multiplier in float 32
func (flagSetUtils FLagSetUtils) GetFloat32GasMultiplier(flagSet *pflag.FlagSet) (float32, error) {
	return flagSet.GetFloat32("gasmultiplier")
}

//This function returns Buffer in Int32
func (flagSetUtils FLagSetUtils) GetInt32Buffer(flagSet *pflag.FlagSet) (int32, error) {
	return flagSet.GetInt32("buffer")
}

//This function returns Wait in Int32
func (flagSetUtils FLagSetUtils) GetInt32Wait(flagSet *pflag.FlagSet) (int32, error) {
	return flagSet.GetInt32("wait")
}

//This function returns GasPrice in Int32
func (flagSetUtils FLagSetUtils) GetInt32GasPrice(flagSet *pflag.FlagSet) (int32, error) {
	return flagSet.GetInt32("gasprice")
}

//This function returns Log Level in string
func (flagSetUtils FLagSetUtils) GetStringLogLevel(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("logLevel")
}

func (flagSetUtils FLagSetUtils) GetInt64RPCTimeout(flagSet *pflag.FlagSet) (int64, error) {
	return flagSet.GetInt64("rpcTimeout")
}

//This function returns Gas Limit in Float32
func (flagSetUtils FLagSetUtils) GetFloat32GasLimit(flagSet *pflag.FlagSet) (float32, error) {
	return flagSet.GetFloat32("gasLimit")
}

//This function returns BountyId in Uint32
func (flagSetUtils FLagSetUtils) GetUint32BountyId(flagSet *pflag.FlagSet) (uint32, error) {
	return flagSet.GetUint32("bountyId")
}

//This function returns the provider of root in string
func (flagSetUtils FLagSetUtils) GetRootStringProvider() (string, error) {
	return rootCmd.PersistentFlags().GetString("provider")
}

//This function returns the gas multiplier of root in float32
func (flagSetUtils FLagSetUtils) GetRootFloat32GasMultiplier() (float32, error) {
	return rootCmd.PersistentFlags().GetFloat32("gasmultiplier")
}

//This function returns the buffer of root in Int32
func (flagSetUtils FLagSetUtils) GetRootInt32Buffer() (int32, error) {
	return rootCmd.PersistentFlags().GetInt32("buffer")
}

//This function returns the wait of root in Int32
func (flagSetUtils FLagSetUtils) GetRootInt32Wait() (int32, error) {
	return rootCmd.PersistentFlags().GetInt32("wait")
}

//This function returns the gas price of root in Int32
func (flagSetUtils FLagSetUtils) GetRootInt32GasPrice() (int32, error) {
	return rootCmd.PersistentFlags().GetInt32("gasprice")
}

//This function returns the log level of root in string
func (flagSetUtils FLagSetUtils) GetRootStringLogLevel() (string, error) {
	return rootCmd.PersistentFlags().GetString("logLevel")
}

//This function returns the gas limit of root in Float32
func (flagSetUtils FLagSetUtils) GetRootFloat32GasLimit() (float32, error) {
	return rootCmd.PersistentFlags().GetFloat32("gasLimit")
}

//This function returns the gas limit of root in Float32
func (flagSetUtils FLagSetUtils) GetRootInt64RPCTimeout() (int64, error) {
	return rootCmd.PersistentFlags().GetInt64("rpcTimeout")
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

//This function returns the AssetId in Uint16
func (flagSetUtils FLagSetUtils) GetUint16AssetId(flagSet *pflag.FlagSet) (uint16, error) {
	return flagSet.GetUint16("assetId")
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
func (*UtilsStruct) GiveSorted(client *ethclient.Client, blockManager *bindings.BlockManager, txnArgs types.TransactionOptions, epoch uint32, assetId uint16, sortedStakers []*big.Int) error {
	return GiveSorted(client, blockManager, txnArgs, epoch, assetId, sortedStakers)
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
