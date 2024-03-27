package accounts

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/mock"
	"io/fs"
	"razor/accounts/mocks"
	"razor/core/types"
	"razor/path"
	mocks1 "razor/path/mocks"
	"reflect"
	"testing"
)

func TestCreateAccount(t *testing.T) {
	var keystorePath string
	var password string
	var fileInfo fs.FileInfo

	type args struct {
		account    accounts.Account
		accountErr error
		statErr    error
		isNotExist bool
		mkdirErr   error
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
		{
			name: "Test 3: When keystore directory does not exists and mkdir creates it",
			args: args{
				account: accounts.Account{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
					URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
				},
				statErr:    errors.New("not exists"),
				isNotExist: true,
			},
			want: accounts.Account{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
				URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
			},
			expectedFatal: false,
		},
		{
			name: "Test 4: When keystore directory does not exists and there an error creating new one",
			args: args{
				account: accounts.Account{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
					URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
				},
				statErr:    errors.New("not exists"),
				isNotExist: true,
				mkdirErr:   errors.New("mkdir error"),
			},
			want: accounts.Account{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
				URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
			},
			expectedFatal: true,
		},
	}

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accountsMock := new(mocks.AccountInterface)
			osMock := new(mocks1.OSInterface)

			path.OSUtilsInterface = osMock
			AccountUtilsInterface = accountsMock

			accountsMock.On("NewAccount", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(tt.args.account, tt.args.accountErr)
			osMock.On("Stat", mock.AnythingOfType("string")).Return(fileInfo, tt.args.statErr)
			osMock.On("IsNotExist", mock.Anything).Return(tt.args.isNotExist)
			osMock.On("Mkdir", mock.Anything, mock.Anything).Return(tt.args.mkdirErr)

			accountUtils := AccountUtils{}
			fatal = false
			got := accountUtils.CreateAccount(keystorePath, password)
			if tt.expectedFatal {
				assert.Equal(t, tt.expectedFatal, fatal)
			}
			if got.Address != tt.want.Address {
				t.Errorf("New address created, got = %v, want %v", got, tt.want.Address)
			}
		})
	}
}

func TestGetPrivateKeyFromKeystore(t *testing.T) {
	var password string
	var keystorePath string
	var privateKey *ecdsa.PrivateKey
	var jsonBytes []byte

	type args struct {
		jsonBytes    []byte
		jsonBytesErr error
		key          *keystore.Key
		keyErr       error
	}
	tests := []struct {
		name    string
		args    args
		want    *ecdsa.PrivateKey
		wantErr bool
	}{
		{
			name: "Test 1: When GetPrivateKey function executes successfully",
			args: args{
				jsonBytes: jsonBytes,
				key: &keystore.Key{
					PrivateKey: privateKey,
				},
			},
			want:    privateKey,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in reading data from file",
			args: args{
				jsonBytesErr: errors.New("error in reading data"),
				key: &keystore.Key{
					PrivateKey: nil,
				},
			},
			want:    nil,
			wantErr: true,
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
			want:    privateKey,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accountsMock := new(mocks.AccountInterface)
			AccountUtilsInterface = accountsMock

			accountsMock.On("ReadFile", mock.AnythingOfType("string")).Return(tt.args.jsonBytes, tt.args.jsonBytesErr)
			accountsMock.On("DecryptKey", mock.Anything, mock.AnythingOfType("string")).Return(tt.args.key, tt.args.keyErr)

			accountUtils := &AccountUtils{}
			got, err := accountUtils.GetPrivateKeyFromKeystore(keystorePath, password)
			if got != tt.want {
				t.Errorf("Private key from GetPrivateKey, got = %v, want %v", got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPrivateKeyFromKeystore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetPrivateKey(t *testing.T) {
	var password string
	var keystorePath string
	var privateKey *ecdsa.PrivateKey

	accountsList := []accounts.Account{
		{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
			URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
		},
		{Address: common.HexToAddress("0x000000000000000000000000000000000000dea2"),
			URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
		},
	}

	type args struct {
		address    string
		accounts   []accounts.Account
		privateKey *ecdsa.PrivateKey
	}
	tests := []struct {
		name    string
		args    args
		want    *ecdsa.PrivateKey
		wantErr bool
	}{
		{
			name: "Test 1: When input address is present in accountsList",
			args: args{
				address:    "0x000000000000000000000000000000000000dea1",
				accounts:   accountsList,
				privateKey: privateKey,
			},
			want:    privateKey,
			wantErr: false,
		},
		{
			name: "Test 2: When input address is not present in accountsList",
			args: args{
				address:    "0x000000000000000000000000000000000000dea3",
				accounts:   accountsList,
				privateKey: privateKey,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accountsMock := new(mocks.AccountInterface)
			AccountUtilsInterface = accountsMock

			accountsMock.On("Accounts", mock.AnythingOfType("string")).Return(tt.args.accounts)
			accountsMock.On("GetPrivateKeyFromKeystore", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(tt.args.privateKey, nil)

			accountUtils := &AccountUtils{}
			got, err := accountUtils.GetPrivateKey(tt.args.address, password, keystorePath)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetPrivateKey() got = %v, want %v", got, tt.want)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPrivateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSignData(t *testing.T) {
	var hash []byte
	var account types.Account
	var defaultPath string
	var privateKey *ecdsa.PrivateKey
	var signature []byte

	type args struct {
		privateKey    *ecdsa.PrivateKey
		privateKeyErr error
		signature     []byte
		signatureErr  error
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{
			name: "Test 1: When Sign function returns no error",
			args: args{
				privateKey:   privateKey,
				signature:    signature,
				signatureErr: nil,
			},
			want:    signature,
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting private key",
			args: args{
				privateKeyErr: errors.New("privateKey"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accountsMock := new(mocks.AccountInterface)
			AccountUtilsInterface = accountsMock

			accountsMock.On("GetPrivateKey", mock.AnythingOfType("string"), mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(tt.args.privateKey, tt.args.privateKeyErr)
			accountsMock.On("Sign", mock.Anything, mock.Anything).Return(tt.args.signature, tt.args.signatureErr)

			accountUtils := &AccountUtils{}

			got, err := accountUtils.SignData(hash, account, defaultPath)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sign() got = %v, want %v", got, tt.want)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
