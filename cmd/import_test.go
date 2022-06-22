package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"github.com/awnumar/memguard"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
	"io/fs"
	"razor/cmd/mocks"
	"razor/path"
	mocks1 "razor/path/mocks"
	mocks2 "razor/utils/mocks"
	"testing"
)

func TestImportAccount(t *testing.T) {

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	var fileInfo fs.FileInfo

	account := accounts.Account{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
		URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
	}

	type args struct {
		privateKey         string
		password           string
		keyBuffer          *memguard.LockedBuffer
		keyBufferBytes     []byte
		decryptData        []byte
		path               string
		pathErr            error
		ecdsaPrivateKey    *ecdsa.PrivateKey
		ecdsaPrivateKeyErr error
		importAccount      accounts.Account
		importAccountErr   error
		statErr            error
		isNotExist         bool
		mkdirErr           error
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
				keyBufferBytes:     []byte("test"),
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
		{
			name: "Test 5: When keystore directory is not present and mkdir creates it",
			args: args{
				privateKey:         "0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d",
				password:           "test",
				path:               "/home/local",
				pathErr:            nil,
				ecdsaPrivateKey:    privateKey,
				ecdsaPrivateKeyErr: nil,
				importAccount:      account,
				importAccountErr:   nil,
				statErr:            errors.New("not exists"),
				isNotExist:         true,
			},
			want:    account,
			wantErr: nil,
		},
		{
			name: "Test 5: When keystore directory is not present and there is an error creating new one",
			args: args{
				privateKey:         "0x4f3edf983ac636a65a842ce7c78d9aa706d3b113bce9c46f30d7d21715b23b1d",
				password:           "test",
				path:               "/home/local",
				pathErr:            nil,
				ecdsaPrivateKey:    privateKey,
				ecdsaPrivateKeyErr: nil,
				importAccount:      account,
				importAccountErr:   nil,
				statErr:            errors.New("not exists"),
				isNotExist:         true,
				mkdirErr:           errors.New("mkdir error"),
			},
			want: accounts.Account{
				Address: common.Address{0x00},
			},
			wantErr: errors.New("mkdir error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterface)
			keystoreUtilsMock := new(mocks.KeystoreInterface)
			cryptoUtilsMock := new(mocks.CryptoInterface)
			osMock := new(mocks1.OSInterface)
			utilsPkgMock := new(mocks2.Utils)

			path.OSUtilsInterface = osMock
			razorUtils = utilsMock
			keystoreUtils = keystoreUtilsMock
			cryptoUtils = cryptoUtilsMock
			utilsInterface = utilsPkgMock

			utilsMock.On("PrivateKeyPrompt").Return(tt.args.privateKey)
			utilsPkgMock.On("InterruptAndPurge")
			utilsMock.On("PasswordPrompt").Return(tt.args.password)
			utilsPkgMock.On("KeyBuffer", mock.Anything).Return(tt.args.keyBuffer)
			utilsPkgMock.On("KeyBufferBytes", mock.Anything).Return(tt.args.keyBufferBytes)
			utilsPkgMock.On("Decrypt", mock.Anything).Return(tt.args.decryptData)
			utilsMock.On("GetDefaultPath").Return(tt.args.path, tt.args.pathErr)
			cryptoUtilsMock.On("HexToECDSA", mock.AnythingOfType("string")).Return(tt.args.ecdsaPrivateKey, tt.args.ecdsaPrivateKeyErr)
			keystoreUtilsMock.On("ImportECDSA", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.importAccount, tt.args.importAccountErr)
			osMock.On("Stat", mock.AnythingOfType("string")).Return(fileInfo, tt.args.statErr)
			osMock.On("IsNotExist", mock.Anything).Return(tt.args.isNotExist)
			osMock.On("Mkdir", mock.Anything, mock.Anything).Return(tt.args.mkdirErr)
			utilsPkgMock.On("DestroyKeyBuffer", mock.Anything)

			utils := &UtilsStruct{}

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
	var flagSet *pflag.FlagSet

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
			utilsMock := new(mocks.UtilsInterface)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)

			cmdUtils = cmdUtilsMock
			razorUtils = utilsMock

			utilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"))
			cmdUtilsMock.On("ImportAccount").Return(tt.args.account, tt.args.accountErr)

			utils := &UtilsStruct{}
			utils.ExecuteImport(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The executeImport function didn't execute as expected")
			}
		})
	}
}
