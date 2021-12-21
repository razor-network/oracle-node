package cmd

import (
	"crypto/ecdsa"
	"math/big"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
	"time"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	Types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
)

type utilsInterface interface {
	GetOptions() bind.CallOpts
	GetTxnOpts(types.TransactionOptions, utils.Utils) *bind.TransactOpts
	WaitForBlockCompletion(*ethclient.Client, string) int
	AssignPassword(*pflag.FlagSet) string
	ConnectToClient(string) *ethclient.Client
	GetStakerId(*ethclient.Client, string) (uint32, error)
	GetStaker(*ethclient.Client, string, uint32) (bindings.StructsStaker, error)
	GetUpdatedStaker(*ethclient.Client, string, uint32) (bindings.StructsStaker, error)
	GetConfigData(UtilsStruct) (types.Configurations, error)
	ParseBool(str string) (bool, error)
	GetDelayedState(*ethclient.Client, int32) (int64, error)
	GetEpoch(*ethclient.Client) (uint32, error)
	GetActiveAssetsData(*ethclient.Client, uint32) ([]*big.Int, error)
	ConvertUintArrayToUint16Array(uintArr []uint) []uint16
	PrivateKeyPrompt() string
	PasswordPrompt() string
	FetchBalance(*ethclient.Client, string) (*big.Int, error)
	CheckAmountAndBalance(*big.Int, *big.Int) *big.Int
	GetAmountInDecimal(*big.Int) *big.Float
	GetDefaultPath() (string, error)
	GetNumberOfStakers(*ethclient.Client, string) (uint32, error)
	GetRandaoHash(*ethclient.Client) ([32]byte, error)
	GetNumberOfProposedBlocks(*ethclient.Client, string, uint32) (uint8, error)
	GetMaxAltBlocks(*ethclient.Client, string) (uint8, error)
	GetProposedBlock(*ethclient.Client, string, uint32, uint32) (bindings.StructsBlock, error)
	GetEpochLastRevealed(*ethclient.Client, string, uint32) (uint32, error)
	GetVoteValue(*ethclient.Client, uint16, uint32) (*big.Int, error)
	GetInfluenceSnapshot(*ethclient.Client, uint32, uint32) (*big.Int, error)
	GetNumActiveAssets(*ethclient.Client) (*big.Int, error)
	GetTotalInfluenceRevealed(*ethclient.Client, uint32) (*big.Int, error)
	ConvertBigIntArrayToUint32Array([]*big.Int) []uint32
	GetLock(*ethclient.Client, string, uint32) (types.Locks, error)
	GetWithdrawReleasePeriod(*ethclient.Client, string) (uint8, error)
	GetMaxCommission(client *ethclient.Client) (uint8, error)
	GetEpochLimitForUpdateCommission(client *ethclient.Client) (uint16, error)
	GetCommitments(*ethclient.Client, string) ([32]byte, error)
	AllZero([32]byte) bool
	GetEpochLastCommitted(*ethclient.Client, uint32) (uint32, error)
	GetConfigFilePath() (string, error)
	ViperWriteConfigAs(string) error
	IsEqual([]uint32, []uint32) (bool, int)
	GetActiveAssetIds(*ethclient.Client) ([]uint16, error)
	GetBlockManager(*ethclient.Client) *bindings.BlockManager
	GetVotes(*ethclient.Client, uint32) (bindings.StructsVote, error)
	Contains(interface{}, interface{}) bool
	CheckEthBalanceIsZero(*ethclient.Client, string)
	AssignStakerId(*pflag.FlagSet, *ethclient.Client, string) (uint32, error)
	GetSortedProposedBlockIds(*ethclient.Client, string, uint32) ([]uint32, error)
	CheckError(msg string, err error)
	GetLatestBlock(*ethclient.Client) (*Types.Header, error)
	GetUpdatedEpoch(*ethclient.Client) (uint32, error)
	GetStateName(int64) string
	getBufferPercent(UtilsStruct) (int32, error)
	IsFlagPassed(string) bool
	GetFractionalAmountInWei(*big.Int, string) (*big.Int, error)
	GetAmountInWei(*big.Int) *big.Int
	Sleep(time.Duration)
	CalculateBlockTime(*ethclient.Client) int64
	getProvider(UtilsStruct) (string, error)
	getMultiplier(UtilsStruct) (float32, error)
	getWaitTime(UtilsStruct) (int32, error)
	getGasPrice(UtilsStruct) (int32, error)
	getLogLevel(UtilsStruct) (string, error)
	getGasLimit(UtilsStruct) (float32, error)
	GetStakedToken(*ethclient.Client, common.Address) *bindings.StakedToken
	ConvertSRZRToRZR(*big.Int, *big.Int, *big.Int) *big.Int
	ConvertRZRToSRZR(*big.Int, *big.Int, *big.Int) (*big.Int, error)
	GetRogueRandomValue(int) *big.Int
}

type tokenManagerInterface interface {
	Allowance(*ethclient.Client, *bind.CallOpts, common.Address, common.Address) (*big.Int, error)
	Approve(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) (*Types.Transaction, error)
	Transfer(*ethclient.Client, *bind.TransactOpts, common.Address, *big.Int) (*Types.Transaction, error)
}

type transactionInterface interface {
	Hash(*Types.Transaction) common.Hash
}

type assetManagerInterface interface {
	CreateJob(*ethclient.Client, *bind.TransactOpts, uint8, int8, uint8, string, string, string) (*Types.Transaction, error)
	SetCollectionStatus(*ethclient.Client, *bind.TransactOpts, bool, uint16) (*Types.Transaction, error)
	GetActiveStatus(*ethclient.Client, *bind.CallOpts, uint16) (bool, error)
	CreateCollection(client *ethclient.Client, opts *bind.TransactOpts, jobIDs []uint16, aggregationMethod uint32, power int8, name string) (*Types.Transaction, error)
	UpdateJob(*ethclient.Client, *bind.TransactOpts, uint16, uint8, int8, uint8, string, string) (*Types.Transaction, error)
	UpdateCollection(*ethclient.Client, *bind.TransactOpts, uint16, uint32, int8, []uint16) (*Types.Transaction, error)
}

type stakeManagerInterface interface {
	Stake(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) (*Types.Transaction, error)
	ExtendLock(*ethclient.Client, *bind.TransactOpts, uint32) (*Types.Transaction, error)
	Delegate(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) (*Types.Transaction, error)
	Withdraw(*ethclient.Client, *bind.TransactOpts, uint32) (*Types.Transaction, error)
	SetDelegationAcceptance(*ethclient.Client, *bind.TransactOpts, bool) (*Types.Transaction, error)
	Unstake(*ethclient.Client, *bind.TransactOpts, uint32, *big.Int) (*Types.Transaction, error)
	RedeemBounty(*ethclient.Client, *bind.TransactOpts, uint32) (*Types.Transaction, error)
	UpdateCommission(client *ethclient.Client, opts *bind.TransactOpts, commission uint8) (*Types.Transaction, error)

	//Getter methods
	StakerInfo(*ethclient.Client, *bind.CallOpts, uint32) (types.Staker, error)
	GetMaturity(*ethclient.Client, *bind.CallOpts, uint32) (uint16, error)
	GetBountyLock(*ethclient.Client, *bind.CallOpts, uint32) (types.BountyLock, error)
	BalanceOf(*bindings.StakedToken, *bind.CallOpts, common.Address) (*big.Int, error)
	GetTotalSupply(*bindings.StakedToken, *bind.CallOpts) (*big.Int, error)
}

type keystoreInterface interface {
	Accounts(string) []accounts.Account
	ImportECDSA(string, *ecdsa.PrivateKey, string) (accounts.Account, error)
}

type flagSetInterface interface {
	GetStringFrom(*pflag.FlagSet) (string, error)
	GetStringTo(*pflag.FlagSet) (string, error)
	GetStringAddress(*pflag.FlagSet) (string, error)
	GetUint32StakerId(*pflag.FlagSet) (uint32, error)
	GetStringName(*pflag.FlagSet) (string, error)
	GetStringUrl(*pflag.FlagSet) (string, error)
	GetStringSelector(*pflag.FlagSet) (string, error)
	GetInt8Power(*pflag.FlagSet) (int8, error)
	GetUint8Weight(*pflag.FlagSet) (uint8, error)
	GetUint16AssetId(*pflag.FlagSet) (uint16, error)
	GetUint8SelectorType(set *pflag.FlagSet) (uint8, error)
	GetStringStatus(*pflag.FlagSet) (string, error)
	GetUint8Commission(*pflag.FlagSet) (uint8, error)
	GetUintSliceJobIds(*pflag.FlagSet) ([]uint, error)
	GetUint32Aggregation(*pflag.FlagSet) (uint32, error)
	GetUint16JobId(*pflag.FlagSet) (uint16, error)
	GetUint16CollectionId(*pflag.FlagSet) (uint16, error)
	GetStringProvider(*pflag.FlagSet) (string, error)
	GetFloat32GasMultiplier(*pflag.FlagSet) (float32, error)
	GetInt32Buffer(*pflag.FlagSet) (int32, error)
	GetInt32Wait(*pflag.FlagSet) (int32, error)
	GetInt32GasPrice(*pflag.FlagSet) (int32, error)
	GetFloat32GasLimit(set *pflag.FlagSet) (float32, error)
	GetStringLogLevel(*pflag.FlagSet) (string, error)
	GetStringValue(*pflag.FlagSet) (string, error)
	GetStringPow(*pflag.FlagSet) (string, error)
	GetBoolAutoWithdraw(*pflag.FlagSet) (bool, error)
	GetUint32BountyId(*pflag.FlagSet) (uint32, error)
	GetRootStringProvider() (string, error)
	GetRootFloat32GasMultiplier() (float32, error)
	GetRootInt32Buffer() (int32, error)
	GetRootInt32Wait() (int32, error)
	GetRootInt32GasPrice() (int32, error)
	getRootStringLogLevel() (string, error)
	GetRootFloat32GasLimit() (float32, error)
}

type utilsCmdInterface interface {
	Withdraw(*ethclient.Client, *bind.TransactOpts, uint32, UtilsStruct) (common.Hash, error)
	CheckCurrentStatus(*ethclient.Client, uint16, UtilsStruct) (bool, error)
	Dispute(*ethclient.Client, types.Configurations, types.Account, uint32, uint8, int, UtilsStruct) error
	GiveSorted(*ethclient.Client, *bindings.BlockManager, *bind.TransactOpts, uint32, uint16, []uint32)
	GetEpochAndState(*ethclient.Client, string, UtilsStruct) (uint32, int64, error)
	WaitForAppropriateState(*ethclient.Client, string, string, UtilsStruct, ...int) (uint32, error)
	WaitIfCommitState(*ethclient.Client, string, string, UtilsStruct) (uint32, error)
	AssignAmountInWei(*pflag.FlagSet, UtilsStruct) (*big.Int, error)
	Unstake(types.Configurations, *ethclient.Client, types.UnstakeInput, UtilsStruct) (types.TransactionOptions, error)
	AutoWithdraw(types.TransactionOptions, uint32, UtilsStruct) error
	withdrawFunds(*ethclient.Client, types.Account, types.Configurations, uint32, UtilsStruct) (common.Hash, error)
	Create(string, UtilsStruct) (accounts.Account, error)
	claimBounty(types.Configurations, *ethclient.Client, types.RedeemBountyInput, UtilsStruct) (common.Hash, error)
	GetAmountInSRZRs(*ethclient.Client, string, bindings.StructsStaker, *big.Int, UtilsStruct) (*big.Int, error)
}

type cryptoInterface interface {
	HexToECDSA(string) (*ecdsa.PrivateKey, error)
}

type voteManagerInterface interface {
	Commit(*ethclient.Client, *bind.TransactOpts, uint32, [32]byte) (*Types.Transaction, error)
	Reveal(*ethclient.Client, *bind.TransactOpts, uint32, []*big.Int, [32]byte) (*Types.Transaction, error)
}

type blockManagerInterface interface {
	ClaimBlockReward(*ethclient.Client, *bind.TransactOpts) (*Types.Transaction, error)
	Propose(*ethclient.Client, *bind.TransactOpts, uint32, []uint32, *big.Int, uint32) (*Types.Transaction, error)
	FinalizeDispute(*ethclient.Client, *bind.TransactOpts, uint32, uint8) (*Types.Transaction, error)
	DisputeBiggestInfluenceProposed(*ethclient.Client, *bind.TransactOpts, uint32, uint8, uint32) (*Types.Transaction, error)
}

type proposeUtilsInterface interface {
	getBiggestInfluenceAndId(*ethclient.Client, string, uint32, UtilsStruct) (*big.Int, uint32, error)
	getIteration(*ethclient.Client, string, types.ElectedProposer, UtilsStruct) int
	isElectedProposer(*ethclient.Client, types.ElectedProposer, UtilsStruct) bool
	pseudoRandomNumberGenerator([]byte, uint32, []byte) *big.Int
	MakeBlock(*ethclient.Client, string, types.Rogue, UtilsStruct) ([]uint32, error)
	getSortedVotes(*ethclient.Client, string, uint16, uint32, UtilsStruct) ([]*big.Int, error)
	influencedMedian([]*big.Int, *big.Int) *big.Int
}
