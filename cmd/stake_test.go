package cmd

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/utils"
	"testing"
)

func getTransactionArgs() types.TransactionOptions {
	password := "12345"
	address := "0x53baCf1E1C6E14BB5700178Acf426dF6f0d5A949" // Provide your address for testing
	provider := "http://127.0.0.1:8545/"
	client, err := ethclient.Dial(provider)
	if err != nil {
		log.Fatal("Error in connecting...\n", err)
	}
	log.Info("Connected to: ", provider)
	balance, balanceErr := utils.FetchBalance(client, address)
	if balanceErr != nil {
		log.Fatalf("Error in fetching balance: %e", balanceErr)
	}
	amount := "1000"
	amountInWei := utils.GetAmountWithChecks(amount, balance)

	txnArgs := types.TransactionOptions{
		Client:         client,
		AccountAddress: address,
		Amount:         amountInWei,
		Password:       password,
		ChainId:        big.NewInt(1337),
		GasMultiplier:  10,
	}
	return txnArgs
}
func getTransactionArgsMinStake() types.TransactionOptions {
	password := "12345"
	address := "0x53baCf1E1C6E14BB5700178Acf426dF6f0d5A949" // Provide your address for testing
	provider := "http://127.0.0.1:8545/"
	client, err := ethclient.Dial(provider)
	if err != nil {
		log.Fatal("Error in connecting...\n", err)
	}
	log.Info("Connected to: ", provider)
	balance, balanceErr := utils.FetchBalance(client, address)
	if balanceErr != nil {
		log.Fatalf("Error in fetching balance: %e", balanceErr)
	}
	amount := "10"
	amountInWei := utils.GetAmountWithChecks(amount, balance)

	txnArgs := types.TransactionOptions{
		Client:         client,
		AccountAddress: address,
		Amount:         amountInWei,
		Password:       password,
		ChainId:        big.NewInt(1337),
		GasMultiplier:  10,
	}
	return txnArgs
}

func Test_approve(t *testing.T) {
	txnArgs := getTransactionArgs()

	t.Run("Test1: Staker is able to approve", func(t *testing.T) {
		approvedStatus := approve(txnArgs)
		if got := approvedStatus; got != nil {
			t.Errorf("Staker is not able to approve")
		}
	})
	t.Run("Test2: Sufficient allowance, no need to increase approved amount ", func(t *testing.T) {
		approvedStatus := approve(txnArgs)
		if got := approvedStatus; got != nil {
			t.Errorf("No Sufficient allowance, increase the amount ")
		}
	})
	t.Run("Test3: Allowance of number of razors is set correctly after approve", func(t *testing.T) {
		tokenManager := utils.GetTokenManager(txnArgs.Client)
		opts := utils.GetOptions(false, txnArgs.AccountAddress, "")
		allowance, allowanceErr := tokenManager.Allowance(&opts, common.HexToAddress(txnArgs.AccountAddress), common.HexToAddress(core.StakeManagerAddress))
		if allowanceErr != nil {
			utils.CheckError("Error in getting allowance amount", allowanceErr)
		}
		if got := allowance; got.Cmp(txnArgs.Amount) != 1 && got.Cmp(txnArgs.Amount) != 0 {
			t.Errorf("Amount of razor's approved are less than the required amount of razor's to be sent")
		}
	})
}

func Test_stakeCoins(t *testing.T) {
	txnArgs := getTransactionArgs()
	opts := utils.GetOptions(false, txnArgs.AccountAddress, "")
	stakeManager := utils.GetStakeManager(txnArgs.Client)

	type args struct {
		txnOpts types.TransactionOptions
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test1: Staker is able to stake for the first time.",
			args: args{
				txnOpts: getTransactionArgs(),
			},
		},
		{
			name: "Test2: Staker is able to stake again",
			args: args{
				txnOpts: getTransactionArgs(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stakerId, err := stakeManager.GetStakerId(&opts, common.HexToAddress(txnArgs.AccountAddress))
			utils.CheckError("Error in getting staker ID: ", err)

			if stakerId.Cmp(big.NewInt(0)) == 0 {
				approveErr := approve(txnArgs)
				utils.CheckError("Approve error: ", approveErr)
				stakeErr := stakeCoins(txnArgs)
				utils.CheckError("Stake error: ", stakeErr)

				newStakerID, stakerIdErr := stakeManager.GetStakerId(&opts, common.HexToAddress(txnArgs.AccountAddress))
				utils.CheckError("Error in getting staker ID: ", stakerIdErr)
				staker, stakerErr := utils.GetStaker(txnArgs.Client, txnArgs.AccountAddress, newStakerID)
				utils.CheckError("Error in getting staker: ", stakerErr)
				expectedStake := big.NewInt(0)
				expectedStake.Add(expectedStake, txnArgs.Amount)
				newStake := staker.Stake

				if got := newStake; got.Cmp(expectedStake) != 0 {
					t.Errorf("Error in Staking for the first time")
				}
			}
			if stakerId.Cmp(big.NewInt(0)) != 0 {
				staker, stakerErr := utils.GetStaker(txnArgs.Client, txnArgs.AccountAddress, stakerId)
				utils.CheckError("Error in getting staker: ", stakerErr)
				previousStake := staker.Stake

				approveErr := approve(txnArgs)
				utils.CheckError("Approve error: ", approveErr)
				stakeErr := stakeCoins(txnArgs)
				utils.CheckError("Stake error: ", stakeErr)
				updatedStaker, updatedStakerErr := utils.GetStaker(txnArgs.Client, txnArgs.AccountAddress, stakerId)
				utils.CheckError("Error in getting staker: ", updatedStakerErr)
				newStake := updatedStaker.Stake

				expectedStake := big.NewInt(0)
				expectedStake.Add(previousStake, txnArgs.Amount)

				if got := newStake; got.Cmp(expectedStake) != 0 {
					t.Errorf("Error in Staking the stake amount")
				}
			}
		})
	}

	t.Run("Test3: Stake should be greater than minimum stake", func(t *testing.T) {
		expectedErr := errors.New("stake amount is less than minimum stake")
		txnArgsMinStake := getTransactionArgsMinStake()
		approveErr := approve(txnArgsMinStake)
		utils.CheckError("Approve error: ", approveErr)
		err := stakeCoins(txnArgsMinStake)
		if expectedErr.Error() != err.Error() {
			t.Errorf("Expected error :%v and got error :%v", expectedErr, err)
		}
	})
}
