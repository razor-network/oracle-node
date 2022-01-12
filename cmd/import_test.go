package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
	"razor/cmd/mocks"
	"testing"
)

func TestImportAccount(t *testing.T) {

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	account := accounts.Account{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
		URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
	}

	type args struct {
		privateKey         string
		password           string
		path               string
		pathErr            error
		ecdsaPrivateKey    *ecdsa.PrivateKey
		ecdsaPrivateKeyErr error
		importAccount      accounts.Account
		importAccountErr   error
	}
	tests := []struct {
		name    string
		args    args
		want    accounts.Account
		wantErr error
	}{
		{
			name: "Test 1: When importAccount executes successfully",
			args: args{
				privateKey:         "0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d",
				password:           "test",
				path:               "/home/local",
				pathErr:            nil,
				ecdsaPrivateKey:    privateKey,
				ecdsaPrivateKeyErr: nil,
				importAccount:      account,
				importAccountErr:   nil,
			},
			want:    account,
			wantErr: nil,
		},
		{
			name: "Test 2: When importAccount fails due to path error",
			args: args{
				privateKey:         "0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d",
				password:           "test",
				path:               "",
				pathErr:            errors.New("path error"),
				ecdsaPrivateKey:    privateKey,
				ecdsaPrivateKeyErr: nil,
				importAccount:      account,
				importAccountErr:   nil,
			},
			want: accounts.Account{Address: common.Address{0x00},
				URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
			},
			wantErr: errors.New("path error"),
		},
		{
			name: "Test 3: When importAccount fails due to parsing privateKey error",
			args: args{
				privateKey:         "0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d",
				password:           "test",
				path:               "/home/local",
				pathErr:            nil,
				ecdsaPrivateKeyErr: errors.New("parsing private key error"),
				importAccount:      account,
				importAccountErr:   nil,
			},
			want: accounts.Account{Address: common.Address{0x00},
				URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
			},
			wantErr: errors.New("parsing private key error"),
		},
		{
			name: "Test 4: When importAccount fails due ImportECDSA error",
			args: args{
				privateKey:         "0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d",
				password:           "test",
				path:               "/home/local",
				pathErr:            nil,
				ecdsaPrivateKey:    privateKey,
				ecdsaPrivateKeyErr: nil,
				importAccount:      account,
				importAccountErr:   errors.New("import error"),
			},
			want: accounts.Account{Address: common.Address{0x00},
				URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
			},
			wantErr: errors.New("import error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterfaceMockery)
			keystoreUtilsMock := new(mocks.KeystoreInterfaceMockery)
			cryptoUtilsMock := new(mocks.CryptoInterface)

			razorUtilsMockery = utilsMock
			keystoreUtilsMockery = keystoreUtilsMock
			cryptoUtils = cryptoUtilsMock

			utilsMock.On("PrivateKeyPrompt").Return(tt.args.privateKey)
			utilsMock.On("PasswordPrompt").Return(tt.args.password)
			utilsMock.On("GetDefaultPath").Return(tt.args.path, tt.args.pathErr)
			cryptoUtilsMock.On("HexToECDSA", mock.AnythingOfType("string")).Return(tt.args.ecdsaPrivateKey, tt.args.ecdsaPrivateKeyErr)
			keystoreUtilsMock.On("ImportECDSA", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.importAccount, tt.args.importAccountErr)

			utils := &UtilsStructMockery{}

			got, err := utils.ImportAccount()
			if got.Address != tt.want.Address {
				t.Errorf("New address imported, got = %v, want %v", got, tt.want.Address)
			}

			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for importAccount function, got = %v, want %v", got, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for importAccount function, got = %v, want %v", got, tt.wantErr)
				}
			}
		})
	}
}

func TestExecuteImport(t *testing.T) {

	type args struct {
		account    accounts.Account
		accountErr error
	}
	tests := []struct {
		name          string
		args          args
		expectedFatal bool
	}{
		{
			name: "Test 1: When ExecuteImport execites successfully",
			args: args{
				account: accounts.Account{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
					URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
				},
			},
			expectedFatal: false,
		},
		{
			name: "Test 2: When there is an error in ImportAccount",
			args: args{
				account:    accounts.Account{},
				accountErr: errors.New("account error"),
			},
			expectedFatal: true,
		},
	}
	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			cmdUtilMock := new(mocks.UtilsCmdInterfaceMockery)
			cmdUtilsMockery = cmdUtilMock

			cmdUtilMock.On("ImportAccount").Return(tt.args.account, tt.args.accountErr)

			utils := &UtilsStructMockery{}
			utils.ExecuteImport()
			if fatal != tt.expectedFatal {
				t.Error("The executeImport function didn't execute as expected")
			}
		})
	}
}
