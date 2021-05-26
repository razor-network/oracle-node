package utils

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/wealdtech/go-merkletree"
	"math/big"
	"razor/core/types"
	"razor/pkg/bindings"
	"reflect"
	"testing"
)

func TestConnectToClient(t *testing.T) {
	type args struct {
		provider string
	}
	tests := []struct {
		name string
		args args
		want *ethclient.Client
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConnectToClient(tt.args.provider); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConnectToClient() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFetchBalance(t *testing.T) {
	type args struct {
		client         *ethclient.Client
		accountAddress string
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FetchBalance(tt.args.client, tt.args.accountAddress)
			if (err != nil) != tt.wantErr {
				t.Errorf("FetchBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FetchBalance() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetActiveJobs(t *testing.T) {
	type args struct {
		client  *ethclient.Client
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    []types.Job
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetActiveJobs(tt.args.client, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetActiveJobs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetActiveJobs() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetCommitments(t *testing.T) {
	type args struct {
		client  *ethclient.Client
		address string
		epoch   *big.Int
	}
	tests := []struct {
		name    string
		args    args
		want    [32]byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetCommitments(tt.args.client, tt.args.address, tt.args.epoch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCommitments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCommitments() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDataInBytes(t *testing.T) {
	type args struct {
		data []*big.Int
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDataInBytes(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDataInBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDefaultPath(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDefaultPath(); got != tt.want {
				t.Errorf("GetDefaultPath() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDelayedState(t *testing.T) {
	type args struct {
		client *ethclient.Client
	}
	tests := []struct {
		name    string
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetDelayedState(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDelayedState() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetDelayedState() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetEpoch(t *testing.T) {
	type args struct {
		client  *ethclient.Client
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetEpoch(tt.args.client, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEpoch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetEpoch() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMerkleTree(t *testing.T) {
	type args struct {
		data []*big.Int
	}
	tests := []struct {
		name    string
		args    args
		want    *merkletree.MerkleTree
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMerkleTree(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMerkleTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMerkleTree() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMerkleTreeRoot(t *testing.T) {
	type args struct {
		data []*big.Int
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMerkleTreeRoot(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMerkleTreeRoot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMerkleTreeRoot() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetMinStakeAmount(t *testing.T) {
	type args struct {
		client  *ethclient.Client
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetMinStakeAmount(tt.args.client, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetMinStakeAmount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMinStakeAmount() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStake(t *testing.T) {
	type args struct {
		client   *ethclient.Client
		address  string
		stakerId *big.Int
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStake(tt.args.client, tt.args.address, tt.args.stakerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStake() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStake() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStaker(t *testing.T) {
	type args struct {
		client   *ethclient.Client
		address  string
		stakerId *big.Int
	}
	tests := []struct {
		name    string
		args    args
		want    bindings.StructsStaker
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStaker(tt.args.client, tt.args.address, tt.args.stakerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStaker() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStaker() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetStakerId(t *testing.T) {
	type args struct {
		client  *ethclient.Client
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetStakerId(tt.args.client, tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStakerId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStakerId() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetVotes(t *testing.T) {
	type args struct {
		client  *ethclient.Client
		address string
		epoch   *big.Int
	}
	tests := []struct {
		name string
		args args
		want struct {
			Value  *big.Int
			Weight *big.Int
		}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetVotes(tt.args.client, tt.args.address, tt.args.epoch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVotes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWaitForBlockCompletion(t *testing.T) {
	type args struct {
		client     *ethclient.Client
		hashToRead string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := WaitForBlockCompletion(tt.args.client, tt.args.hashToRead); got != tt.want {
				t.Errorf("WaitForBlockCompletion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkTransactionReceipt(t *testing.T) {
	type args struct {
		client  *ethclient.Client
		_txHash string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkTransactionReceipt(tt.args.client, tt.args._txHash); got != tt.want {
				t.Errorf("checkTransactionReceipt() = %v, want %v", got, tt.want)
			}
		})
	}
}
