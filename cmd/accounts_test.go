package cmd

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/magiconair/properties/assert"
	"razor/core/types"
	"reflect"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	var path string
	var password string

	utilsStruct := UtilsStruct{
		keystoreUtils: KeystoreMock{},
	}
	type args struct {
		account    accounts.Account
		accountErr error
	}
	tests := []struct {
		name          string
		args          args
		want          accounts.Account
		expectedFatal bool
	}{
		{
			name: "Test 1: When NewAccounts executes successfully",
			args: args{
				account: accounts.Account{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
					URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
				},
			},
			want: accounts.Account{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
				URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in getting new account",
			args: args{
				accountErr: errors.New("account error"),
			},
			want:          accounts.Account{Address: common.HexToAddress("0x00")},
			expectedFatal: true,
		},
	}

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			NewAccountMock = func(string, string) (accounts.Account, error) {
				return tt.args.account, tt.args.accountErr
			}

			fatal = false
			got := CreateAccount(path, password, utilsStruct)
			if tt.expectedFatal {
				assert.Equal(t, tt.expectedFatal, fatal)
			}
			if got.Address != tt.want.Address {
				t.Errorf("New address created, got = %v, want %v", got, tt.want.Address)
			}
		})
	}
}

func TestGetPrivateKey(t *testing.T) {

}

func TestSign(t *testing.T) {
	type args struct {
		hash        []byte
		account     types.Account
		defaultPath string
		utilsStruct UtilsStruct
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
			got, err := Sign(tt.args.hash, tt.args.account, tt.args.defaultPath, tt.args.utilsStruct)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sign() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPrivateKeyFromKeystore(t *testing.T) {
	var password string
	var keystorePath string
	var privateKey *ecdsa.PrivateKey
	var jsonBytes []byte

	utilsStruct := UtilsStruct{
		razorUtils:    UtilsMock{},
		keystoreUtils: KeystoreMock{},
	}

	type args struct {
		jsonBytes    []byte
		jsonBytesErr error
		key          *keystore.Key
		keyErr       error
	}
	tests := []struct {
		name          string
		args          args
		want          *ecdsa.PrivateKey
		expectedFatal bool
	}{
		{
			name: "Test 1: When GetPrivateKey function executes successfully",
			args: args{
				jsonBytes: jsonBytes,
				key: &keystore.Key{
					PrivateKey: privateKey,
				},
			},
			want:          privateKey,
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in reading data from file",
			args: args{
				jsonBytesErr: errors.New("error in reading data"),
				key: &keystore.Key{
					PrivateKey: nil,
				},
			},
			want:          nil,
			expectedFatal: true,
		},
		{
			name: "Test 3: When there is an error in fetching private key",
			args: args{
				jsonBytes: jsonBytes,
				key: &keystore.Key{
					PrivateKey: nil,
				},
				keyErr: errors.New("private key error"),
			},
			want:          privateKey,
			expectedFatal: true,
		},
	}

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ReadFileMock = func(string) ([]byte, error) {
				return tt.args.jsonBytes, tt.args.jsonBytesErr
			}

			DecryptKeyMock = func([]byte, string) (*keystore.Key, error) {
				return tt.args.key, tt.args.keyErr
			}

			fatal = false
			got := getPrivateKeyFromKeystore(keystorePath, password, utilsStruct)
			if tt.expectedFatal {
				assert.Equal(t, tt.expectedFatal, fatal)
			}
			if got != tt.want {
				t.Errorf("Private key from GetPrivateKey, got = %v, want %v", got, tt.want)
			}
		})
	}
}
