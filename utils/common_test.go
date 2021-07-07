package utils

import (
	"github.com/wealdtech/go-merkletree"
	"math/big"
	"reflect"
	"testing"
)

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
