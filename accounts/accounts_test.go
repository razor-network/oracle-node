package accounts

import (
	"crypto/ecdsa"
	"encoding/hex"
	"reflect"
	"testing"
)

//func TestCreateAccount(t *testing.T) {
//	var keystorePath string
//	var password string
//	var fileInfo fs.FileInfo
//
//	type args struct {
//		account    accounts.Account
//		accountErr error
//		statErr    error
//		isNotExist bool
//		mkdirErr   error
//	}
//	tests := []struct {
//		name          string
//		args          args
//		want          accounts.Account
//		expectedFatal bool
//	}{
//		{
//			name: "Test 1: When NewAccounts executes successfully",
//			args: args{
//				account: accounts.Account{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
//					URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
//				},
//			},
//			want: accounts.Account{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
//				URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
//			},
//			expectedFatal: false,
//		},
//		{
//			name: "Test 2: When there is an error in getting new account",
//			args: args{
//				accountErr: errors.New("account error"),
//			},
//			want:          accounts.Account{Address: common.HexToAddress("0x00")},
//			expectedFatal: true,
//		},
//		{
//			name: "Test 3: When keystore directory does not exists and mkdir creates it",
//			args: args{
//				account: accounts.Account{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
//					URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
//				},
//				statErr:    errors.New("not exists"),
//				isNotExist: true,
//			},
//			want: accounts.Account{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
//				URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
//			},
//			expectedFatal: false,
//		},
//		{
//			name: "Test 4: When keystore directory does not exists and there an error creating new one",
//			args: args{
//				account: accounts.Account{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
//					URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
//				},
//				statErr:    errors.New("not exists"),
//				isNotExist: true,
//				mkdirErr:   errors.New("mkdir error"),
//			},
//			want: accounts.Account{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
//				URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
//			},
//			expectedFatal: true,
//		},
//	}
//
//	defer func() { log.ExitFunc = nil }()
//	var fatal bool
//	log.ExitFunc = func(int) { fatal = true }
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			accountsMock := new(mocks.AccountManagerInterface)
//			osMock := new(mocks1.OSInterface)
//
//			path.OSUtilsInterface = osMock
//
//			accountsMock.On("NewAccount", mock.AnythingOfType("string"), mock.AnythingOfType("string")).Return(tt.args.account, tt.args.accountErr)
//			osMock.On("Stat", mock.AnythingOfType("string")).Return(fileInfo, tt.args.statErr)
//			osMock.On("IsNotExist", mock.Anything).Return(tt.args.isNotExist)
//			osMock.On("Mkdir", mock.Anything, mock.Anything).Return(tt.args.mkdirErr)
//
//			fatal = false
//
//			got := accountsMock.CreateAccount(keystorePath, password)
//			if tt.expectedFatal {
//				assert.Equal(t, tt.expectedFatal, fatal)
//			}
//			if got.Address != tt.want.Address {
//				t.Errorf("New address created, got = %v, want %v", got, tt.want.Address)
//			}
//		})
//	}
//}

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
		address  string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test 1: When input address with correct password is present in keystore directory",
			args: args{
				address:  "0x911654feb423363fb771e04e18d1e7325ae10a91",
				password: password,
			},
			wantErr: false,
		},
		{
			name: "Test 2: When input upper case address with correct password is present in keystore directory",
			args: args{
				address:  "0x2F5F59615689B706B6AD13FD03343DCA28784989",
				password: password,
			},
			wantErr: false,
		},
		{
			name: "Test 3: When provided address is not present in keystore directory",
			args: args{
				address: "0x911654feb423363fb771e04e18d1e7325ae10a91_not_present",
			},
			wantErr: true,
		},
		{
			name: "Test 4: When input address with incorrect password is present in keystore directory",
			args: args{
				address:  "0x911654feb423363fb771e04e18d1e7325ae10a91",
				password: "incorrect password",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			am := NewAccountManager(keystoreDirPath)
			_, err := am.GetPrivateKey(tt.args.address, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetPrivateKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestSignData(t *testing.T) {
	password := "Razor@123"

	hexHash := "a3b8d42c7015c1e9354f8b9c2161d9b2e1ad89e6b6c7a9610e029fd7afec27ae"
	hashBytes, err := hex.DecodeString(hexHash)
	if err != nil {
		log.Fatal("Failed to decode hex string")
	}

	type args struct {
		address  string
		password string
		hash     []byte
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Test 1: When Sign function returns no error",
			args: args{
				address:  "0x911654feb423363fb771e04e18d1e7325ae10a91",
				password: password,
				hash:     hashBytes,
			},
			want:    "f14cf1b8c9486777e4280b4da6cb7c314d5c19b7e30d32a46f83767a1946e35a39f6941df71375d7ffceaddac81e2454e9a129896803d02f633eb78ab7883ff200",
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting private key",
			args: args{
				address:  "0x_invalid_address",
				password: password,
				hash:     hashBytes,
			},
			want:    hex.EncodeToString(nil),
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			am := NewAccountManager("test_accounts")
			got, err := am.SignData(tt.args.hash, tt.args.address, tt.args.password)
			if !reflect.DeepEqual(hex.EncodeToString(got), tt.want) {
				t.Errorf("Sign() got = %v, want %v", hex.EncodeToString(got), tt.want)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Sign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
