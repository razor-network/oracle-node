package cmd

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"testing"
)

func Test_importAccount(t *testing.T) {

	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)

	account := accounts.Account{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
		URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
	}

	razorUtils := UtilsMock{}
	keystoreUtils := KeystoreMock{}
	cryptoUtils := CryptoMock{}

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

			PrivateKeyPromptMock = func() string {
				return tt.args.privateKey
			}

			PasswordPromptMock = func() string {
				return tt.args.password
			}

			GetDefaultPathMock = func() (string, error) {
				return tt.args.path, tt.args.pathErr
			}

			HexToECDSAMock = func(string) (*ecdsa.PrivateKey, error) {
				return tt.args.ecdsaPrivateKey, tt.args.ecdsaPrivateKeyErr
			}

			ImportECDSAMock = func(string, *ecdsa.PrivateKey, string) (accounts.Account, error) {
				return tt.args.importAccount, tt.args.importAccountErr
			}

			got, err := importAccount(razorUtils, keystoreUtils, cryptoUtils)
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
