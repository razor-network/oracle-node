package cmd

import (
	"errors"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/mock"
	"razor/cmd/mocks"
	"reflect"
	"testing"
)

func Test_listAccounts(t *testing.T) {

	accountsList := []accounts.Account{
		{Address: common.HexToAddress("0x000000000000000000000000000000000000dea1"),
			URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
		},
		{Address: common.HexToAddress("0x000000000000000000000000000000000000dea2"),
			URL: accounts.URL{Scheme: "TestKeyScheme", Path: "test/key/path"},
		},
	}

	type args struct {
		path     string
		pathErr  error
		accounts []accounts.Account
	}
	tests := []struct {
		name    string
		args    args
		want    []accounts.Account
		wantErr error
	}{
		{
			name: "When listAccounts executes successfully",
			args: args{
				path:     "test/key/path",
				pathErr:  nil,
				accounts: accountsList,
			},
			want:    accountsList,
			wantErr: nil,
		},
		{
			name: "When listAccounts fails due to path error",
			args: args{
				path:     "test/key/",
				pathErr:  errors.New("path error"),
				accounts: accountsList,
			},
			want:    nil,
			wantErr: errors.New("path error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			utilsMock := new(mocks.UtilsInterfaceMockery)
			keystoreUtilsMock := new(mocks.KeystoreInterfaceMockery)

			razorUtilsMockery = utilsMock
			keystoreUtilsMockery = keystoreUtilsMock

			utilsMock.On("GetDefaultPath").Return(tt.args.path, tt.args.pathErr)
			keystoreUtilsMock.On("Accounts", mock.AnythingOfType("string")).Return(tt.args.accounts)
			utils := &UtilsStructMockery{}
			got, err := utils.ListAccounts()

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("List of accounts , got = %v, want %v", got, tt.want)
			}

			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for listAccounts function, got = %v, want = %v", got, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for listAccounts function, got = %v, want = %v", got, tt.wantErr)
				}
			}

		})
	}
}
