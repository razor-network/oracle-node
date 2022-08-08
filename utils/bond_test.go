package utils

import (
	"errors"
	"github.com/avast/retry-go"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/stretchr/testify/mock"
	"razor/pkg/bindings"
	"razor/utils/mocks"
	"reflect"
	"testing"
)

func TestGetBondManagerWithOpts(t *testing.T) {
	var callOpts bind.CallOpts
	var bondManager *bindings.BondManager
	var client *ethclient.Client

	utilsMock := new(mocks.Utils)

	optionsPackageStruct := OptionsPackageStruct{
		UtilsInterface: utilsMock,
	}
	utils := StartRazor(optionsPackageStruct)

	utilsMock.On("GetOptions").Return(callOpts)
	utilsMock.On("GetBondManager", mock.AnythingOfType("*ethclient.Client")).Return(bondManager)

	gotBondManager, gotCallOpts := utils.GetBondManagerWithOpts(client)
	if !reflect.DeepEqual(gotCallOpts, callOpts) {
		t.Errorf("GetBondManagerWithOpts() got callopts = %v, want %v", gotCallOpts, callOpts)
	}
	if !reflect.DeepEqual(gotBondManager, bondManager) {
		t.Errorf("GetBondManagerWithOpts() got bondManager = %v, want %v", gotBondManager, bondManager)
	}
}

func TestGetDataBondCollections(t *testing.T) {
	var client *ethclient.Client

	type args struct {
		dataBondCollections    []uint16
		dataBondCollectionsErr error
	}
	tests := []struct {
		name    string
		args    args
		want    []uint16
		wantErr bool
	}{
		{
			name: "Test 1: When GetDataBondCollections() executes successfully",
			args: args{
				dataBondCollections: []uint16{1, 2, 3},
			},
			want:    []uint16{1, 2, 3},
			wantErr: false,
		},
		{
			name: "Test 2: When there is an error in getting dataBondCollections",
			args: args{
				dataBondCollectionsErr: errors.New("dataBondCollections error"),
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			retryMock := new(mocks.RetryUtils)
			bondManagerMock := new(mocks.BondManagerUtils)

			optionsPackageStruct := OptionsPackageStruct{
				RetryInterface:       retryMock,
				BondManagerInterface: bondManagerMock,
			}
			utils := StartRazor(optionsPackageStruct)

			bondManagerMock.On("GetDataBondCollections", mock.AnythingOfType("*ethclient.Client")).Return(tt.args.dataBondCollections, tt.args.dataBondCollectionsErr)
			retryMock.On("RetryAttempts", mock.AnythingOfType("uint")).Return(retry.Attempts(1))

			got, err := utils.GetDataBondCollections(client)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDataBondCollections() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDataBondCollections() got = %v, want %v", got, tt.want)
			}
		})
	}
}
