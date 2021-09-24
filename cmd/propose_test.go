package cmd

import (
	"math/big"
	"reflect"
	"testing"
)

func Test_influencedMedian(t *testing.T) {
	type args struct {
		sortedVotes            []*big.Int
		totalInfluenceRevealed *big.Int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		{
			name: "Test if sortedVotes is empty",
			args: args{
				sortedVotes:            []*big.Int{},
				totalInfluenceRevealed: big.NewInt(100000000),
			},
			want: big.NewInt(0),
		},
		{
			name: "Test if totalInfluenceRevealed is 0",
			args: args{
				sortedVotes:            []*big.Int{big.NewInt(100)},
				totalInfluenceRevealed: big.NewInt(0),
			},
			want: big.NewInt(100),
		},
		{
			name: "Test if all the values are present",
			args: args{
				sortedVotes:            []*big.Int{big.NewInt(100), big.NewInt(110), big.NewInt(115), big.NewInt(118)},
				totalInfluenceRevealed: big.NewInt(400),
			},
			want: big.NewInt(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := influencedMedian(tt.args.sortedVotes, tt.args.totalInfluenceRevealed); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("influencedMedian() = %v, want %v", got, tt.want)
			}
		})
	}
}
