package utils

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	solsha3 "github.com/miguelmota/go-solidity-sha3"
	"math/big"
	"os"
	"path"
	"razor/accounts"
	"razor/core/types"
	"reflect"
	"testing"
)

func TestEcRecover(t *testing.T) {
	dir, _ := os.Getwd()
	razorPath := path.Dir(dir)
	testKeystorePath := path.Join(razorPath, "utils/test_accounts")

	accounts.AccountUtilsInterface = accounts.AccountUtils{}
	hash := solsha3.SoliditySHA3([]string{"address", "uint32", "uint256", "string"}, []interface{}{common.HexToAddress("0x57Baf83BAD5bee0F7F44d84669A50C35c57E3576"), 9201, big.NewInt(31337), "razororacle"})
	ethHash := SignHash(hash)
	signedData, _ := accounts.AccountUtilsInterface.SignData(ethHash, types.Account{Address: "0x57Baf83BAD5bee0F7F44d84669A50C35c57E3576",
		Password: "Test@123"}, testKeystorePath)
	modifiedSignedData, _ := accounts.AccountUtilsInterface.SignData(ethHash, types.Account{Address: "0x57Baf83BAD5bee0F7F44d84669A50C35c57E3576",
		Password: "Test@123"}, testKeystorePath)
	modifiedSignedData[64] = 27
	type args struct {
		data hexutil.Bytes
		sig  hexutil.Bytes
	}
	tests := []struct {
		name    string
		args    args
		want    common.Address
		wantErr bool
	}{
		{
			name: "Test 1: Recovers address successfully",
			args: args{
				data: hash,
				sig:  signedData,
			},
			want:    common.HexToAddress("0x57Baf83BAD5bee0F7F44d84669A50C35c57E3576"),
			wantErr: false,
		},
		{
			name: "Test 2: Signature length is not equal to 65",
			args: args{
				data: nil,
				sig:  nil,
			},
			want:    common.HexToAddress(""),
			wantErr: true,
		},
		{
			name: "Test 3: Cannot recover address due to error in getting public key from signature",
			args: args{
				data: hash,
				sig:  modifiedSignedData,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EcRecover(tt.args.data, tt.args.sig)
			if (err != nil) != tt.wantErr {
				t.Errorf("EcRecover() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("EcRecover() got = %v, want %v", got, tt.want)
			}
		})
	}
}
