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
	"razor/core"
	coretypes "razor/core/types"
	"razor/path"
	"razor/pkg/bindings"
	"razor/rpc"
	"reflect"
	"strings"
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
		log.Debug("Blockchain function: ", methodName)
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

func InvokeFunctionWithRetryAttempts(rpcParameters rpc.RPCParameters, interfaceName interface{}, methodName string, args ...interface{}) ([]reflect.Value, error) {
	var returnedValues []reflect.Value
	var err error
	var contextError bool

	// Ensure inputs has space for the client and any additional arguments
	inputs := make([]reflect.Value, len(args)+1)

	// Always use the current best client for each retry
	client, err := rpcParameters.RPCManager.GetBestRPCClient()
	if err != nil {
		log.Errorf("Failed to get current best client: %v", err)
		return returnedValues, err
	}

	// Set the client as the first argument
	inputs[0] = reflect.ValueOf(client)
	// Add the rest of the args to inputs starting from index 1
	for i := 0; i < len(args); i++ {
		inputs[i+1] = reflect.ValueOf(args[i])
	}

	err = retry.Do(
		func() error {
			// Check if the context has been cancelled or timed out
			select {
			case <-rpcParameters.Ctx.Done():
				// If context is done, return the context error timeout
				log.Debugf("Context timed out for method: %s", methodName)
				contextError = true
				return retry.Unrecoverable(rpcParameters.Ctx.Err())
			default:
				// Proceed with the RPC call
				returnedValues = reflect.ValueOf(interfaceName).MethodByName(methodName).Call(inputs)
				err = CheckIfAnyError(returnedValues)
				if err != nil {
					log.Debug("Function to retry: ", methodName)
					log.Errorf("Error in %v....Retrying", methodName)
					return err
				}
				return nil
			}
		}, RetryInterface.RetryAttempts(core.MaxRetries), retry.Delay(time.Second*time.Duration(core.RetryDelayDuration)), retry.DelayType(retry.FixedDelay))
	if err != nil {
		if contextError {
			// If context error, we don't switch the client
			log.Warnf("Skipping switching to the next best client due to context error: %v", err)
			return returnedValues, err
		}

		// Only switch to the next best client if the error is identified as an RPC error
		if isRPCError(err) {
			log.Errorf("%v error after retries: %v", methodName, err)
			log.Info("Attempting to switch to a new best RPC endpoint...")

			switched, switchErr := rpcParameters.RPCManager.SwitchToNextBestRPCClient()
			if switchErr != nil {
				log.Errorf("Failed to switch to the next best client: %v", switchErr)
				return returnedValues, switchErr
			}

			if switched {
				log.Infof("Successfully switched to a new RPC endpoint after RPC error.")
			} else {
				log.Warnf("No switch occurred. Retaining the current RPC client.")
			}
		}
	}

	return returnedValues, err
}

func isRPCError(err error) bool {
	// Check for common RPC error patterns
	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
		return true
	}

	// Add other checks based on specific RPC errors (timeouts, connection issues, etc.)
	if strings.Contains(err.Error(), "connection refused") ||
		strings.Contains(err.Error(), "connection reset by peer") ||
		strings.Contains(err.Error(), "EOF") ||
		strings.Contains(err.Error(), "dial") ||
		strings.Contains(err.Error(), "no such host") ||
		strings.Contains(err.Error(), "i/o timeout") {
		return true
	}

	// Check for HTTP 500–504 errors
	if strings.Contains(err.Error(), "500") ||
		strings.Contains(err.Error(), "501") ||
		strings.Contains(err.Error(), "502") ||
		strings.Contains(err.Error(), "503") ||
		strings.Contains(err.Error(), "504") {
		return true
	}

	// Check for the custom RPC timeout error
	if strings.Contains(err.Error(), "RPC timeout error") {
		return true
	}

	// If it's not an RPC error, return false
	return false
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

func (b BlockManagerStruct) GetConfirmedBlocks(client *ethclient.Client, epoch uint32) (coretypes.ConfirmedBlock, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(blockManager, "Blocks", &opts, epoch)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return coretypes.ConfirmedBlock{}, returnedError
	}
	return returnedValues[0].Interface().(struct {
		Valid        bool
		ProposerId   uint32
		Iteration    *big.Int
		BiggestStake *big.Int
	}), nil
}

func (b BlockManagerStruct) Disputes(client *ethclient.Client, epoch uint32, address common.Address) (coretypes.DisputesStruct, error) {
	blockManager, opts := UtilsInterface.GetBlockManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(blockManager, "Disputes", &opts, epoch, address)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return coretypes.DisputesStruct{}, returnedError
	}
	disputesMapping := returnedValues[0].Interface().(struct {
		LeafId           uint16
		LastVisitedValue *big.Int
		AccWeight        *big.Int
		Median           *big.Int
	})
	return disputesMapping, nil
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

func (s StakeManagerStruct) StakerInfo(client *ethclient.Client, stakerId uint32) (coretypes.Staker, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(stakeManager, "Stakers", &opts, stakerId)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return coretypes.Staker{}, returnedError
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

func (s StakeManagerStruct) GetMaturity(client *ethclient.Client, age uint32) (uint16, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	index := age / 10000
	returnedValues := InvokeFunctionWithTimeout(stakeManager, "Maturities", opts, big.NewInt(int64(index)))
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return 0, returnedError
	}
	return returnedValues[0].Interface().(uint16), nil
}

func (s StakeManagerStruct) GetBountyLock(client *ethclient.Client, bountyId uint32) (coretypes.BountyLock, error) {
	stakeManager, opts := UtilsInterface.GetStakeManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(stakeManager, "BountyLocks", &opts, bountyId)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return coretypes.BountyLock{}, returnedError
	}
	bountyLock := returnedValues[0].Interface().(struct {
		RedeemAfter  uint32
		BountyHunter common.Address
		Amount       *big.Int
	})
	return bountyLock, nil
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

func (a AssetManagerStruct) GetActiveStatus(client *ethclient.Client, id uint16) (bool, error) {
	collectionManager, opts := UtilsInterface.GetCollectionManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(collectionManager, "GetCollectionStatus", &opts, id)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return false, returnedError
	}
	return returnedValues[0].Interface().(bool), nil
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

func (v VoteManagerStruct) GetCommitment(client *ethclient.Client, stakerId uint32) (coretypes.Commitment, error) {
	voteManager, opts := UtilsInterface.GetVoteManagerWithOpts(client)
	returnedValues := InvokeFunctionWithTimeout(voteManager, "Commitments", &opts, stakerId)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
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

func (s StakedTokenStruct) BalanceOf(client *ethclient.Client, tokenAddress common.Address, address common.Address) (*big.Int, error) {
	stakedToken, opts := UtilsInterface.GetStakedTokenManagerWithOpts(client, tokenAddress)
	returnedValues := InvokeFunctionWithTimeout(stakedToken, "BalanceOf", &opts, address)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return nil, returnedError
	}
	return returnedValues[0].Interface().(*big.Int), nil
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

func (c CoinStruct) BalanceOf(client *ethclient.Client, account common.Address) (*big.Int, error) {
	tokenManager := UtilsInterface.GetTokenManager(client)
	opts := UtilsInterface.GetOptions()
	returnedValues := InvokeFunctionWithTimeout(tokenManager, "BalanceOf", &opts, account)
	returnedError := CheckIfAnyError(returnedValues)
	if returnedError != nil {
		return nil, returnedError
	}
	return returnedValues[0].Interface().(*big.Int), nil
}

//This function is used to check the allowance of staker
func (c CoinStruct) Allowance(client *ethclient.Client, owner common.Address, spender common.Address) (*big.Int, error) {
	tokenManager := UtilsInterface.GetTokenManager(client)
	opts := UtilsInterface.GetOptions()
	returnedValues := InvokeFunctionWithTimeout(tokenManager, "Allowance", &opts, owner, spender)
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

func (r RetryStruct) RetryAttempts(numberOfAttempts uint) retry.Option {
	return retry.Attempts(numberOfAttempts)
}

func (f FlagSetStruct) GetLogFileName(flagSet *pflag.FlagSet) (string, error) {
	return flagSet.GetString("logFile")
}
