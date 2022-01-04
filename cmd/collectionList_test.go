package cmd

import (
	"errors"
	"github.com/ethereum/go-ethereum/ethclient"
	"razor/pkg/bindings"
	"testing"
)

func TestUtilsStruct_GetCollectionList(t *testing.T) {
	var client *ethclient.Client
	type fields struct {
		razorUtils UtilsMock
	}
	testUtils := fields{
		razorUtils: UtilsMock{},
	}

	collectionListArray := []bindings.StructsCollection{
		{Active: true, Id: 7, AssetIndex: 1, Power: 2,
			AggregationMethod: 2, JobIDs: []uint16{1, 2, 3}, Name: "ethCollectionMean",
		},
		{Active: true, Id: 8, AssetIndex: 2, Power: 2,
			AggregationMethod: 2, JobIDs: []uint16{4, 5, 6}, Name: "btcCollectionMean",
		},
	}

	type args struct {
		client            *ethclient.Client
		collectionList    []bindings.StructsCollection
		collectionListErr error
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name:   "Test 1: When collectionList executes properly",
			fields: testUtils,
			args: args{
				client:            client,
				collectionList:    collectionListArray,
				collectionListErr: nil,
			},

			wantErr: nil,
		},
		{
			name:   "Test 2: When there is a error fetching collection list ",
			fields: testUtils,
			args: args{
				client:            client,
				collectionListErr: errors.New("error in fetching collection list"),
			},
			wantErr: errors.New("error in fetching collection list"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsStruct := &UtilsStruct{
				razorUtils: tt.fields.razorUtils,
			}

			GetCollectionsMock = func(*ethclient.Client) ([]bindings.StructsCollection, error) {
				return tt.args.collectionList, tt.args.collectionListErr
			}
			err := utilsStruct.GetCollectionList(tt.args.client)

			if err == nil || tt.wantErr == nil {
				if err != tt.wantErr {
					t.Errorf("Error for collectionList function, got = %v, want = %v", err, tt.wantErr)
				}
			} else {
				if err.Error() != tt.wantErr.Error() {
					t.Errorf("Error for collectionList function, got = %v, want = %v", err, tt.wantErr)
				}
			}
		})
	}
}
