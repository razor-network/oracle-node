package utils

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"io"
	"io/fs"
	"math/big"
	"os"
	"razor/accounts"
	coretypes "razor/core/types"
	"razor/path"
	"razor/pkg/bindings"
	"reflect"
	"time"

	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
)

var RPCTimeout int64

func StartRazor(optionsPackageStruct OptionsPackageStruct) Utils {
	UtilsInterface = optionsPackageStruct.UtilsInterface
	EthClient = optionsPackageStruct.EthClient
	ClientInterface = optionsPackageStruct.ClientInterface
	Time = optionsPackageStruct.Time
	OS = optionsPackageStruct.OS
	CoinInterface = optionsPackageStruct.CoinInterface
	IOInterface = optionsPackageStruct.IOInterface
	ABIInterface = optionsPackageStruct.ABIInterface
	PathInterface = optionsPackageStruct.PathInterface
	BindInterface = optionsPackageStruct.BindInterface
	AccountsInterface = optionsPackageStruct.AccountsInterface
	BlockManagerInterface = optionsPackageStruct.BlockManagerInterface
	StakeManagerInterface = optionsPackageStruct.StakeManagerInterface
	AssetManagerInterface = optionsPackageStruct.AssetManagerInterface
	VoteManagerInterface = optionsPackageStruct.VoteManagerInterface
	BindingsInterface = optionsPackageStruct.BindingsInterface
	JsonInterface = optionsPackageStruct.JsonInterface
	StakedTokenInterface = optionsPackageStruct.StakedTokenInterface
	RetryInterface = optionsPackageStruct.RetryInterface
	MerkleInterface = optionsPackageStruct.MerkleInterface
	FlagSetInterface = optionsPackageStruct.FlagSetInterface
	FileInterface = optionsPackageStruct.FileInterface
	GasInterface = optionsPackageStruct.GasInterface
	return &UtilsStruct{}
}

func InvokeFunctionWithTimeout(interfaceName interface{}, methodName string, args ...interface{}) []reflect.Value {
	var functionCall []reflect.Value
	var gotFunction = make(chan bool)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(RPCTimeout)*time.Second)
	defer cancel()

	go func() {
		inputs := make([]reflect.Value, len(args))
		for i := range args {
			inputs[i] = reflect.ValueOf(args[i])
		}
		log.Debug("Function: ", methodName)
		functionCall = reflect.ValueOf(interfaceName).MethodByName(methodName).Call(inputs)
		gotFunction <- true
	}()
	for {
		select {
		case <-ctx.Done():
			log.Errorf("%s function timeout!", methodName)
			log.Debug("Kindly check your connection")
			return nil

		case <-gotFunction:
			return functionCall
		}
	}
}

func CheckIfAnyError(result []reflect.Value) error {
	if result == nil {
		return errors.New("RPC timeout error")
	}

	errorDataType := reflect.TypeOf((*error)(nil)).Elem()
	errorIndexInReturnedValues := -1

	for i := range result {
		returnedValue := result[i]
		returnedValueDataType := reflect.TypeOf(returnedValue.Interface())
		if returnedValueDataType != nil {
			if returnedValueDataType.Implements(errorDataType) {
				errorIndexInReturnedValues = i
			}
		}
	}
	if errorIndexInReturnedValues == -1 {
		return nil
	}
	returnedError := result[errorIndexInReturnedValues].Interface()
	if returnedError != nil {
		return returnedError.(error)
	}
	return nil
}

func (b BlockManagerStruct) GetBlockIndexToBeConfirmed(client *ethclient.Client) (int8, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(blockManager, "BlockIndexToBeConfirmed", &opts)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return -1, returnedError
	}
	return returnedValues[0].Interface().(int8), nil
}

func (s StakeManagerStruct) WithdrawInitiationPeriod(client *ethclient.Client) (uint16, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(stakeManager, "WithdrawInitiationPeriod", &opts)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return 0, returnedError
	}
	return returnedValues[0].Interface().(uint16), nil
}

func (a AssetManagerStruct) GetNumJobs(client *ethclient.Client) (uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(collectionManager, "GetNumJobs", &opts)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return 0, returnedError
	}
	return returnedValues[0].Interface().(uint16), nil
}

func (a AssetManagerStruct) GetCollection(client *ethclient.Client, id uint16) (bindings.StructsCollection, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(collectionManager, "GetCollection", &opts, id)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return bindings.StructsCollection{}, returnedError
	}
	return returnedValues[0].Interface().(bindings.StructsCollection), nil
}

func (a AssetManagerStruct) GetJob(client *ethclient.Client, id uint16) (bindings.StructsJob, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(collectionManager, "GetJob", &opts, id)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return bindings.StructsJob{}, returnedError
	}
	return returnedValues[0].Interface().(bindings.StructsJob), nil
}

func (a AssetManagerStruct) GetCollectionIdFromIndex(client *ethclient.Client, index uint16) (uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(collectionManager, "LeafIdToCollectionIdRegistry", &opts, index)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return 0, returnedError
	}
	return returnedValues[0].Interface().(uint16), nil
}

func (a AssetManagerStruct) GetCollectionIdFromLeafId(client *ethclient.Client, leafId uint16) (uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(collectionManager, "GetCollectionIdFromLeafId", &opts, leafId)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return 0, returnedError
	}
	return returnedValues[0].Interface().(uint16), nil
}

func (a AssetManagerStruct) GetLeafIdOfACollection(client *ethclient.Client, collectionId uint16) (uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(collectionManager, "GetLeafIdOfCollection", &opts, collectionId)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return 0, returnedError
	}
	return returnedValues[0].Interface().(uint16), nil
}

func (v VoteManagerStruct) ToAssign(client *ethclient.Client) (uint16, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(voteManager, "ToAssign", &opts)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return 0, returnedError
	}
	return returnedValues[0].Interface().(uint16), nil
}

func (v VoteManagerStruct) GetSaltFromBlockchain(client *ethclient.Client) ([32]byte, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(voteManager, "GetSalt", &opts)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return [32]byte{}, returnedError
	}
	return returnedValues[0].Interface().([32]byte), nil
}

func (a AccountsStruct) GetPrivateKey(address string, password string, keystorePath string) (*ecdsa.PrivateKey, error) {
	return accounts.AccountUtilsInterface.GetPrivateKey(address, password, keystorePath)
}

func (b BlockManagerStruct) GetNumProposedBlocks(client *ethclient.Client, epoch uint32) (uint8, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(blockManager, "GetNumProposedBlocks", &opts, epoch)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return 0, returnedError
	}
	return returnedValues[0].Interface().(uint8), nil
}

func (b BlockManagerStruct) GetProposedBlock(client *ethclient.Client, epoch uint32, proposedBlock uint32) (bindings.StructsBlock, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(blockManager, "GetProposedBlock", &opts, epoch, proposedBlock)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return bindings.StructsBlock{}, returnedError
	}
	return returnedValues[0].Interface().(bindings.StructsBlock), nil
}

func (b BlockManagerStruct) GetBlock(client *ethclient.Client, epoch uint32) (bindings.StructsBlock, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(blockManager, "GetBlock", &opts, epoch)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return bindings.StructsBlock{}, returnedError
	}
	return returnedValues[0].Interface().(bindings.StructsBlock), nil
}

func (b BlockManagerStruct) MinStake(client *ethclient.Client) (*big.Int, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(blockManager, "MinStake", &opts)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return nil, returnedError
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

func (b BlockManagerStruct) StateBuffer(client *ethclient.Client) (uint8, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(blockManager, "Buffer", &opts)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return 0, returnedError
	}
	return returnedValues[0].Interface().(uint8), nil
}

func (b BlockManagerStruct) MaxAltBlocks(client *ethclient.Client) (uint8, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(blockManager, "MaxAltBlocks", &opts)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return 0, returnedError
	}
	return returnedValues[0].Interface().(uint8), nil
}

func (b BlockManagerStruct) SortedProposedBlockIds(client *ethclient.Client, arg0 uint32, arg1 *big.Int) (uint32, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(blockManager, "SortedProposedBlockIds", &opts, arg0, arg1)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return 0, returnedError
	}
	return returnedValues[0].Interface().(uint32), nil
}

func (b BlockManagerStruct) GetEpochLastProposed(client *ethclient.Client, stakerId uint32) (uint32, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(blockManager, "EpochLastProposed", &opts, stakerId)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return 0, returnedError
	}
	return returnedValues[0].Interface().(uint32), nil
}

func (s StakeManagerStruct) GetStakerId(client *ethclient.Client, address common.Address) (uint32, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(stakeManager, "GetStakerId", &opts, address)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return 0, returnedError
	}
	return returnedValues[0].Interface().(uint32), nil
}

func (s StakeManagerStruct) GetNumStakers(client *ethclient.Client) (uint32, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(stakeManager, "GetNumStakers", &opts)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return 0, returnedError
	}
	return returnedValues[0].Interface().(uint32), nil
}

func (s StakeManagerStruct) MinSafeRazor(client *ethclient.Client) (*big.Int, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(stakeManager, "MinSafeRazor", &opts)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return nil, returnedError
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

func (s StakeManagerStruct) Locks(client *ethclient.Client, address common.Address, address1 common.Address, lockType uint8) (coretypes.Locks, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(stakeManager, "Locks", &opts, address, address1, lockType)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return coretypes.Locks{}, returnedError
	}
	locks := returnedValues[0].Interface().(struct {
		Amount      *big.Int
		UnlockAfter *big.Int
	})
	return locks, nil
}

func (s StakeManagerStruct) MaxCommission(client *ethclient.Client) (uint8, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(stakeManager, "MaxCommission", &opts)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return 0, returnedError
	}
	return returnedValues[0].Interface().(uint8), nil
}

func (s StakeManagerStruct) EpochLimitForUpdateCommission(client *ethclient.Client) (uint16, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(stakeManager, "EpochLimitForUpdateCommission", &opts)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return 0, returnedError
	}
	return returnedValues[0].Interface().(uint16), nil
}

func (s StakeManagerStruct) GetStaker(client *ethclient.Client, stakerId uint32) (bindings.StructsStaker, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(stakeManager, "GetStaker", &opts, stakerId)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return bindings.StructsStaker{}, returnedError
	}
	return returnedValues[0].Interface().(bindings.StructsStaker), nil
}

func (a AssetManagerStruct) GetNumCollections(client *ethclient.Client) (uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(collectionManager, "GetNumCollections", &opts)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return 0, returnedError
	}
	return returnedValues[0].Interface().(uint16), nil
}

func (a AssetManagerStruct) GetNumActiveCollections(client *ethclient.Client) (uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(collectionManager, "GetNumActiveCollections", &opts)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return 0, returnedError
	}
	return returnedValues[0].Interface().(uint16), nil
}

func (a AssetManagerStruct) GetActiveCollections(client *ethclient.Client) ([]uint16, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(collectionManager, "GetActiveCollections", &opts)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return nil, returnedError
	}
	return returnedValues[0].Interface().([]uint16), nil
}

func (a AssetManagerStruct) Jobs(client *ethclient.Client, id uint16) (bindings.StructsJob, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(collectionManager, "Jobs", &opts, id)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return bindings.StructsJob{}, returnedError
	}
	job := returnedValues[0].Interface().(struct {
		Id           uint16
		SelectorType uint8
		Weight       uint8
		Power        int8
		Name         string
		Selector     string
		Url          string
	})
	return job, nil
}

func (v VoteManagerStruct) Commitments(client *ethclient.Client, stakerId uint32) (coretypes.Commitment, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(voteManager, "Commitments", &opts, stakerId)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedValues != nil {
		return coretypes.Commitment{}, returnedError
	}
	commitment := returnedValues[0].Interface().(struct {
		Epoch          uint32
		CommitmentHash [32]byte
	})
	return commitment, nil
}

func (v VoteManagerStruct) GetVoteValue(client *ethclient.Client, epoch uint32, stakerId uint32, medianIndex uint16) (*big.Int, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(voteManager, "GetVoteValue", &opts, epoch, stakerId, medianIndex)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return nil, returnedError
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

func (v VoteManagerStruct) GetInfluenceSnapshot(client *ethclient.Client, epoch uint32, stakerId uint32) (*big.Int, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(voteManager, "GetInfluenceSnapshot", &opts, epoch, stakerId)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return nil, returnedError
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

func (v VoteManagerStruct) GetStakeSnapshot(client *ethclient.Client, epoch uint32, stakerId uint32) (*big.Int, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(voteManager, "GetStakeSnapshot", &opts, epoch, stakerId)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return nil, returnedError
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

func (v VoteManagerStruct) GetTotalInfluenceRevealed(client *ethclient.Client, epoch uint32, medianIndex uint16) (*big.Int, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(voteManager, "GetTotalInfluenceRevealed", &opts, epoch, medianIndex)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return nil, returnedError
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

func (v VoteManagerStruct) GetEpochLastCommitted(client *ethclient.Client, stakerId uint32) (uint32, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(voteManager, "GetEpochLastCommitted", &opts, stakerId)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return 0, returnedError
	}
	return returnedValues[0].Interface().(uint32), nil
}

func (v VoteManagerStruct) GetEpochLastRevealed(client *ethclient.Client, stakerId uint32) (uint32, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(voteManager, "GetEpochLastRevealed", &opts, stakerId)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return 0, returnedError
	}
	return returnedValues[0].Interface().(uint32), nil
}

func (b BindingsStruct) NewCollectionManager(address common.Address, client *ethclient.Client) (*bindings.CollectionManager, error) {
	return bindings.NewCollectionManager(address, client)
}

func (b BindingsStruct) NewRAZOR(address common.Address, client *ethclient.Client) (*bindings.RAZOR, error) {
	return bindings.NewRAZOR(address, client)
}

func (b BindingsStruct) NewStakeManager(address common.Address, client *ethclient.Client) (*bindings.StakeManager, error) {
	return bindings.NewStakeManager(address, client)
}

func (b BindingsStruct) NewVoteManager(address common.Address, client *ethclient.Client) (*bindings.VoteManager, error) {
	return bindings.NewVoteManager(address, client)
}

func (b BindingsStruct) NewBlockManager(address common.Address, client *ethclient.Client) (*bindings.BlockManager, error) {
	return bindings.NewBlockManager(address, client)
}

func (b BindingsStruct) NewStakedToken(address common.Address, client *ethclient.Client) (*bindings.StakedToken, error) {
	return bindings.NewStakedToken(address, client)
}

func (j JsonStruct) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func (j JsonStruct) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (u UtilsStruct) GetUint32(flagSet *pflag.FlagSet, name string) (uint32, error) {
	return flagSet.GetUint32(name)
}

func (e EthClientStruct) Dial(rawurl string) (*ethclient.Client, error) {
	return ethclient.Dial(rawurl)
}

func (t TimeStruct) Sleep(duration time.Duration) {
	time.Sleep(duration)
}

func (o OSStruct) OpenFile(name string, flag int, perm fs.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}

func (o OSStruct) Open(name string) (*os.File, error) {
	return os.Open(name)
}

func (o OSStruct) WriteFile(name string, data []byte, perm fs.FileMode) error {
	return os.WriteFile(name, data, perm)
}

func (o OSStruct) ReadFile(filename string) ([]byte, error) {
	return os.ReadFile(filename)
}

func (i IOStruct) ReadAll(body io.ReadCloser) ([]byte, error) {
	return io.ReadAll(body)
}

func (c ClientStruct) TransactionReceipt(client *ethclient.Client, ctx context.Context, txHash common.Hash) (*types.Receipt, error) {
	returnedValues := InvokeFunctionWithTimeout(client, "TransactionReceipt", ctx, txHash)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return &types.Receipt{}, returnedError
	}
	return returnedValues[0].Interface().(*types.Receipt), nil
}

func (c ClientStruct) BalanceAt(client *ethclient.Client, ctx context.Context, account common.Address, blockNumber *big.Int) (*big.Int, error) {
	returnedValues := InvokeFunctionWithTimeout(client, "BalanceAt", ctx, account, blockNumber)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return nil, returnedError
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

func (c ClientStruct) HeaderByNumber(client *ethclient.Client, ctx context.Context, number *big.Int) (*types.Header, error) {
	returnedValues := InvokeFunctionWithTimeout(client, "HeaderByNumber", ctx, number)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return &types.Header{}, returnedError
	}
	return returnedValues[0].Interface().(*types.Header), nil
}

func (c ClientStruct) NonceAt(client *ethclient.Client, ctx context.Context, account common.Address) (uint64, error) {
	var blockNumber *big.Int
	returnedValues := InvokeFunctionWithTimeout(client, "NonceAt", ctx, account, blockNumber)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return 0, returnedError
	}
	return returnedValues[0].Interface().(uint64), nil
}

func (c ClientStruct) SuggestGasPrice(client *ethclient.Client, ctx context.Context) (*big.Int, error) {
	returnedValues := InvokeFunctionWithTimeout(client, "SuggestGasPrice", ctx)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return nil, returnedError
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

func (c ClientStruct) EstimateGas(client *ethclient.Client, ctx context.Context, msg ethereum.CallMsg) (uint64, error) {
	returnedValues := InvokeFunctionWithTimeout(client, "EstimateGas", ctx, msg)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return 0, returnedError
	}
	return returnedValues[0].Interface().(uint64), nil
}

func (c ClientStruct) FilterLogs(client *ethclient.Client, ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	returnedValues := InvokeFunctionWithTimeout(client, "FilterLogs", ctx, q)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return []types.Log{}, returnedError
	}
	return returnedValues[0].Interface().([]types.Log), nil
}

func (c CoinStruct) BalanceOf(erc20Contract *bindings.RAZOR, opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	returnedValues := InvokeFunctionWithTimeout(erc20Contract, "BalanceOf", opts, account)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return nil, returnedError
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

func (a ABIStruct) Parse(reader io.Reader) (abi.ABI, error) {
	return abi.JSON(reader)
}

func (a ABIStruct) Pack(parsedData abi.ABI, name string, args ...interface{}) ([]byte, error) {
	return parsedData.Pack(name, args...)
}

func (p PathStruct) GetDefaultPath() (string, error) {
	return path.PathUtilsInterface.GetDefaultPath()
}

func (p PathStruct) GetJobFilePath() (string, error) {
	return path.PathUtilsInterface.GetJobFilePath()
}

func (b BindStruct) NewKeyedTransactorWithChainID(key *ecdsa.PrivateKey, chainID *big.Int) (*bind.TransactOpts, error) {
	return bind.NewKeyedTransactorWithChainID(key, chainID)
}

func (s StakedTokenStruct) BalanceOf(stakedToken *bindings.StakedToken, callOpts *bind.CallOpts, address common.Address) (*big.Int, error) {
	returnedValues := InvokeFunctionWithTimeout(stakedToken, "BalanceOf", callOpts, address)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return nil, returnedError
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

func (r RetryStruct) RetryAttempts(numberOfAttempts uint) retry.Option {
	return retry.Attempts(numberOfAttempts)
}

func (f FlagSetStruct) GetLogFileName(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("logFile")
}
