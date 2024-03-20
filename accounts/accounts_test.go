package accounts

import (
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/magiconair/properties/assert"
	"github.com/stretchr/testify/mock"
	"io/fs"
	"path/filepath"
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
			name: "Test 2: When another input address with correct password is present in keystore directory",
			args: args{
				address:         "0x2f5f59615689b706b6ad13fd03343dca28784989",
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
		{
			name: "Test 5: When a keystore file is renamed differently from the address to which it belonged",
			args: args{
				address:         "0x811654feb423363fb771e04e18d1e7325ae10a91",
				password:        password,
				keystoreDirPath: "test_accounts/incorrect_test_accounts",
			},
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

func TestFindKeystoreFileForAddress(t *testing.T) {
	testAccountsKeystorePath := "test_accounts"

	tests := []struct {
		name         string
		keystoreDir  string
		address      string
		expectedFile string
		expectErr    bool
	}{
		{
			name:         "Test 1: Matching file exists for an address",
			keystoreDir:  testAccountsKeystorePath,
			address:      "0x911654feb423363fb771e04e18d1e7325ae10a91",
			expectedFile: filepath.Join(testAccountsKeystorePath, "UTC--2024-03-20T07-03-56.358521000Z--911654feb423363fb771e04e18d1e7325ae10a91"),
			expectErr:    false,
		},
		{
			name:         "Test 2: Matching file exists for another address",
			keystoreDir:  testAccountsKeystorePath,
			address:      "0x2f5f59615689b706b6ad13fd03343dca28784989",
			expectedFile: filepath.Join(testAccountsKeystorePath, "UTC--2024-03-20T07-04-11.601622000Z--2f5f59615689b706b6ad13fd03343dca28784989"),
			expectErr:    false,
		},
		{
			name:        "Test 3: No matching file",
			keystoreDir: testAccountsKeystorePath,
			address:     "nonexistentaddress",
			expectErr:   true,
		},
		{
			name:        "Test 4: When keystore directory doesnt exists",
			keystoreDir: "test_accounts_invalid",
			address:     "0x2f5f59615689b706b6ad13fd03343dca28784989",
			expectErr:   true,
		},
		{
			name:         "Test 5: When multiple files for same account is present in the keystore directory",
			keystoreDir:  "test_accounts/incorrect_test_accounts",
			address:      "0x811654feb423363fb771e04e18d1e7325ae10a91",
			expectedFile: filepath.Join("test_accounts/incorrect_test_accounts", "UTC--2024-03-20T07-04-56.358521000Z--811654feb423363fb771e04e18d1e7325ae10a91"),
			expectErr:    true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := FindKeystoreFileForAddress(tc.keystoreDir, tc.address)
			if tc.expectErr {
				if err == nil {
					t.Errorf("Expected an error but got none")
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect an error but got one: %v", err)
				}
				if got != tc.expectedFile {
					t.Errorf("Expected file %v, got %v", tc.expectedFile, got)
				}
			}
		})
	}
}
