package cmd

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/pflag"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"
	"razor/utils"
	"time"

	"github.com/spf13/cobra"
)

var unstakeCmd = &cobra.Command{
	Use:   "unstake",
	Short: "Unstake your razors",
	Long: `unstake allows user to unstake their sRzrs in the razor network

Example:	
  ./razor unstake --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --value 1000 --autoWithdraw
	`,
	Run: initialiseUnstake,
}

func initialiseUnstake(cmd *cobra.Command, args []string) {
	cmdUtilsMockery.ExecuteUnstake(cmd.Flags())
}

func (*UtilsStructMockery) ExecuteUnstake(flagSet *pflag.FlagSet) {
	config, err := cmdUtilsMockery.GetConfigData()
	utils.CheckError("Error in getting config: ", err)

	password := razorUtilsMockery.AssignPassword(flagSet)
	address, err := flagSetUtilsMockery.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)

	autoWithdraw, err := flagSetUtilsMockery.GetBoolAutoWithdraw(flagSet)
	utils.CheckError("Error in getting autoWithdraw status: ", err)

	client := razorUtilsMockery.ConnectToClient(config.Provider)

	valueInWei, err := cmdUtilsMockery.AssignAmountInWei(flagSet)
	utils.CheckError("Error in getting amountInWei: ", err)

	razorUtilsMockery.CheckEthBalanceIsZero(client, address)

	stakerId, err := razorUtilsMockery.AssignStakerId(flagSet, client, address)
	utils.CheckError("StakerId error: ", err)

	lock, err := razorUtilsMockery.GetLock(client, address, stakerId)
	utils.CheckError("Error in getting lock: ", err)

	if lock.Amount.Cmp(big.NewInt(0)) != 0 {
		err = errors.New("existing lock")
		log.Fatal(err)
	}

	unstakeInput := types.UnstakeInput{
		Address:    address,
		Password:   password,
		ValueInWei: valueInWei,
		StakerId:   stakerId,
	}

	txnOptions, err := cmdUtilsMockery.Unstake(config, client, unstakeInput)
	utils.CheckError("Unstake Error: ", err)
	if autoWithdraw {
		err = cmdUtilsMockery.AutoWithdraw(txnOptions, stakerId)
		utils.CheckError("AutoWithdraw Error: ", err)
	}
}

func (*UtilsStructMockery) Unstake(config types.Configurations, client *ethclient.Client, input types.UnstakeInput) (types.TransactionOptions, error) {
	txnArgs := types.TransactionOptions{
		Client:          client,
		Password:        input.Password,
		AccountAddress:  input.Address,
		Amount:          input.ValueInWei,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.StakeManagerAddress,
		MethodName:      "unstake",
		ABI:             bindings.StakeManagerABI,
	}
	stakerId := input.StakerId
	lock, err := razorUtilsMockery.GetLock(txnArgs.Client, txnArgs.AccountAddress, stakerId)
	if err != nil {
		log.Error("Error in getting lock: ", err)
		return txnArgs, err
	}

	if lock.Amount.Cmp(big.NewInt(0)) != 0 {
		err := errors.New("existing lock")
		log.Error(err)
		return txnArgs, err
	}

	staker, err := razorUtilsMockery.GetStaker(client, txnArgs.AccountAddress, stakerId)
	if err != nil {
		log.Error("Error in getting staker: ", err)
		return txnArgs, err
	}

	sAmount, err := cmdUtilsMockery.GetAmountInSRZRs(client, txnArgs.AccountAddress, staker, txnArgs.Amount)
	if err != nil {
		log.Error("Error in getting sRZR amount: ", err)
		return txnArgs, err
	}

	epoch, err := cmdUtilsMockery.WaitForAppropriateState(txnArgs.Client, "unstake", 0, 1, 4)
	if err != nil {
		log.Error("Error in fetching epoch: ", err)
		return txnArgs, err
	}
	txnArgs.Parameters = []interface{}{epoch, stakerId, txnArgs.Amount}
	txnOpts := razorUtilsMockery.GetTxnOpts(txnArgs)
	log.Info("Unstaking coins")
	txn, err := stakeManagerUtilsMockery.Unstake(txnArgs.Client, txnOpts, stakerId, sAmount)
	if err != nil {
		log.Error("Error in un-staking: ", err)
		return txnArgs, err
	}
	log.Info("Transaction hash: ", transactionUtilsMockery.Hash(txn))
	razorUtilsMockery.WaitForBlockCompletion(txnArgs.Client, transactionUtilsMockery.Hash(txn).String())
	return txnArgs, nil
}

func (*UtilsStructMockery) AutoWithdraw(txnArgs types.TransactionOptions, stakerId uint32) error {
	log.Info("Starting withdrawal now...")
	razorUtilsMockery.Sleep(time.Duration(core.EpochLength) * time.Second)
	txn, err := cmdUtilsMockery.WithdrawFunds(txnArgs.Client, types.Account{
		Address:  txnArgs.AccountAddress,
		Password: txnArgs.Password,
	}, txnArgs.Config, stakerId)
	if err != nil {
		log.Error("WithdrawFunds error ", err)
		return err
	}
	if txn != core.NilHash {
		razorUtilsMockery.WaitForBlockCompletion(txnArgs.Client, txn.String())
	}
	return nil
}

func (*UtilsStructMockery) GetAmountInSRZRs(client *ethclient.Client, address string, staker bindings.StructsStaker, amount *big.Int) (*big.Int, error) {
	stakedToken := razorUtilsMockery.GetStakedToken(client, staker.TokenAddress)
	callOpts := razorUtilsMockery.GetOptions()

	sRZRBalance, err := stakeManagerUtilsMockery.BalanceOf(stakedToken, &callOpts, common.HexToAddress(address))
	if err != nil {
		log.Error("Error in getting sRZRBalance: ", err)
		return nil, err
	}

	totalSupply, err := stakeManagerUtilsMockery.GetTotalSupply(stakedToken, &callOpts)
	if err != nil {
		log.Error("Error in getting total supply: ", err)
		return nil, err
	}

	maxUnstake := razorUtilsMockery.ConvertSRZRToRZR(sRZRBalance, staker.Stake, totalSupply)
	log.Debugf("The maximum RZRs you can unstake: %g RZRs", razorUtilsMockery.GetAmountInDecimal(maxUnstake))

	if maxUnstake.Cmp(amount) < 0 {
		log.Error("Amount exceeds maximum unstake amount")
		return nil, errors.New("invalid amount")
	}

	sAmount, err := razorUtilsMockery.ConvertRZRToSRZR(amount, staker.Stake, totalSupply)
	if err != nil {
		log.Error("Error in getting sAmount: ", err)
		return nil, err
	}
	return sAmount, nil
}

func init() {

	cmdUtilsMockery = &UtilsStructMockery{}
	razorUtilsMockery = UtilsMockery{}
	transactionUtilsMockery = TransactionUtilsMockery{}
	stakeManagerUtilsMockery = StakeManagerUtilsMockery{}
	flagSetUtilsMockery = FLagSetUtilsMockery{}
	utils.Options = &utils.OptionsStruct{}
	utils.UtilsInterface = &utils.UtilsStruct{}

	rootCmd.AddCommand(unstakeCmd)

	var (
		Address               string
		AmountToUnStake       string
		WithdrawAutomatically bool
		Password              string
		Power                 string
		StakerId              uint32
	)

	unstakeCmd.Flags().StringVarP(&Address, "address", "a", "", "user's address")
	unstakeCmd.Flags().StringVarP(&AmountToUnStake, "value", "v", "0", "value of sRazors to un-stake")
	unstakeCmd.Flags().BoolVarP(&WithdrawAutomatically, "autoWithdraw", "", false, "withdraw after un-stake automatically")
	unstakeCmd.Flags().StringVarP(&Password, "password", "", "", "password path to protect the keystore")
	unstakeCmd.Flags().StringVarP(&Power, "pow", "", "", "power of 10")
	unstakeCmd.Flags().Uint32VarP(&StakerId, "stakerId", "", 0, "staker id")

	addrErr := unstakeCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)
	valueErr := unstakeCmd.MarkFlagRequired("value")
	utils.CheckError("Value error: ", valueErr)

}
