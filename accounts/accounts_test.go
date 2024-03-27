package accounts

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts"
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

func Test_getPrivateKeyFromKeystore(t *testing.T) {
	password := "Razor@123"

	type args struct {
		keystoreFilePath string
		password         string
	}
	tests := []struct {
		name    string
		args    args
		want    *ecdsa.PrivateKey
		wantErr bool
	}{
		{
			name: "Test 1: When keystore file is present and getPrivateKeyFromKeystore function executes successfully",
			args: args{
				keystoreFilePath: "test_accounts/UTC--2024-03-20T07-03-56.358521000Z--911654feb423363fb771e04e18d1e7325ae10a91",
				password:         password,
			},
			wantErr: false,
		},
		{
			name: "Test 2: When there is no keystore file present at the desired path",
			args: args{
				keystoreFilePath: "test_accounts/UTC--2024-03-20T07-03-56.358521000Z--211654feb423363fb771e04e18d1e7325ae10a91",
				password:         password,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Test 3: When password is incorrect for the desired keystore file",
			args: args{
				keystoreFilePath: "test_accounts/UTC--2024-03-20T07-03-56.358521000Z--911654feb423363fb771e04e18d1e7325ae10a91",
				password:         "Razor@456",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := getPrivateKeyFromKeystore(tt.args.keystoreFilePath, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPrivateKeyFromKeystore() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetPrivateKey(t *testing.T) {
	password := "Razor@123"
	keystoreDirPath := "test_accounts"

	type args struct {
		address         string
		password        string
		keystoreDirPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test 1: When input address with correct password is present in keystore directory",
			args: args{
				address:         "0x911654feb423363fb771e04e18d1e7325ae10a91",
				password:        password,
				keystoreDirPath: keystoreDirPath,
			},
			wantErr: false,
		},
		{
			name: "Test 2: When input upper case address with correct password is present in keystore directory",
			args: args{
				address:         "0x2F5F59615689B706B6AD13FD03343DCA28784989",
				password:        password,
				keystoreDirPath: keystoreDirPath,
			},
			wantErr: false,
		},
		{
			name: "Test 3: When provided address is not present in keystore directory",
			args: args{
				address:         "0x911654feb423363fb771e04e18d1e7325ae10a91_not_present",
				keystoreDirPath: keystoreDirPath,
			},
			wantErr: true,
		},
		{
			name: "Test 4: When input address with incorrect password is present in keystore directory",
			args: args{
				address:         "0x911654feb423363fb771e04e18d1e7325ae10a91",
				password:        "incorrect password",
				keystoreDirPath: keystoreDirPath,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			accountUtils := &AccountUtils{}
			_, err := accountUtils.GetPrivateKey(tt.args.address, tt.args.password, tt.args.keystoreDirPath)
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
