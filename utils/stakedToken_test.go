package utils

import (
	"errors"
	"github.com/avast/retry-go"
	"github.com/stretchr/testify/mock"
	"math/big"
	"razor/pkg/bindings"
	"razor/utils/mocks"
	"reflect"
	"testing"
)

func TestGetStakerSRZRBalance(t *testing.T) {
	var staker bindings.StructsStaker
	type args struct {
		sRZR    *big.Int
		sRZRErr error
	}
	tests := []struct {
		name    string
		args    args
		want    *big.Int
		wantErr bool
	}{
		{
			name: "Test 1: When GetStakerSRZRBalance executes successfully",
			args: args{
				sRZR:    big.NewInt(2000),
				sRZRErr: nil,
			},
			want:    big.NewInt(2000),
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error from BalanceOf()",
			args: args{
				sRZRErr: errors.New("sRZR error"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(mocks.Utils)
			stakedTokenMock := new(mocks.StakedTokenUtils)
			retryMock := new(mocks.RetryUtils)

			utils := StartRazor(OptionsPackageStruct{
				UtilsInterface:       utilsMock,
				StakedTokenInterface: stakedTokenMock,
				RetryInterface:       retryMock,
			})

			stakedTokenMock.On("BalanceOf", mock.Anything, mock.Anything, mock.Anything).Return(tt.args.sRZR, tt.args.sRZRErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetStakerSRZRBalance(rpcParameters, staker)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetStakerSRZRBalance() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetStakerSRZRBalance() got = %v, want %v", got, tt.want)
			}
		})
	}
}
