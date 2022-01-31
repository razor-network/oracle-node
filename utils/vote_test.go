package utils

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"razor/pkg/bindings"
	"reflect"
	"testing"
)

func TestUtilsStruct_GetCommitments(t *testing.T) {
	type args struct {
		client  *ethclient.Client
		address string
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
			ut := &UtilsStruct{}
			got, err := ut.GetCommitments(tt.args.client, tt.args.address)
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

func TestUtilsStruct_GetEpochLastCommitted(t *testing.T) {
	type args struct {
		client   *ethclient.Client
		stakerId uint32
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ut := &UtilsStruct{}
			got, err := ut.GetEpochLastCommitted(tt.args.client, tt.args.stakerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEpochLastCommitted() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetEpochLastCommitted() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUtilsStruct_GetEpochLastRevealed(t *testing.T) {
	type args struct {
		client   *ethclient.Client
		stakerId uint32
	}
	tests := []struct {
		name    string
		args    args
		want    uint32
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ut := &UtilsStruct{}
			got, err := ut.GetEpochLastRevealed(tt.args.client, tt.args.stakerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetEpochLastRevealed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetEpochLastRevealed() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUtilsStruct_GetInfluenceSnapshot(t *testing.T) {
	type args struct {
		client   *ethclient.Client
		stakerId uint32
		epoch    uint32
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
			ut := &UtilsStruct{}
			got, err := ut.GetInfluenceSnapshot(tt.args.client, tt.args.stakerId, tt.args.epoch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInfluenceSnapshot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInfluenceSnapshot() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUtilsStruct_GetRandaoHash(t *testing.T) {
	type args struct {
		client *ethclient.Client
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
			ut := &UtilsStruct{}
			got, err := ut.GetRandaoHash(tt.args.client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRandaoHash() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRandaoHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUtilsStruct_GetStakeSnapshot(t *testing.T) {
	type args struct {
		client   *ethclient.Client
		stakerId uint32
		epoch    uint32
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
			ut := &UtilsStruct{}
			got, err := ut.GetStakeSnapshot(tt.args.client, tt.args.stakerId, tt.args.epoch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStakeSnapshot() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStakeSnapshot() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUtilsStruct_GetTotalInfluenceRevealed(t *testing.T) {
	type args struct {
		client *ethclient.Client
		epoch  uint32
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
			ut := &UtilsStruct{}
			got, err := ut.GetTotalInfluenceRevealed(tt.args.client, tt.args.epoch)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTotalInfluenceRevealed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTotalInfluenceRevealed() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUtilsStruct_GetVoteManagerWithOpts(t *testing.T) {
	type args struct {
		client *ethclient.Client
	}
	tests := []struct {
		name  string
		args  args
		want  *bindings.VoteManager
		want1 bind.CallOpts
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ut := &UtilsStruct{}
			got, got1 := ut.GetVoteManagerWithOpts(tt.args.client)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVoteManagerWithOpts() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("GetVoteManagerWithOpts() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestUtilsStruct_GetVoteValue(t *testing.T) {
	type args struct {
		client   *ethclient.Client
		assetId  uint16
		stakerId uint32
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
			ut := &UtilsStruct{}
			got, err := ut.GetVoteValue(tt.args.client, tt.args.assetId, tt.args.stakerId)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetVoteValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetVoteValue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUtilsStruct_GetVotes(t *testing.T) {
	type args struct {
		client   *ethclient.Client
		stakerId uint32
	}
	tests := []struct {
		name    string
		args    args
		want    bindings.StructsVote
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ut := &UtilsStruct{}
			got, err := ut.GetVotes(tt.args.client, tt.args.stakerId)
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
