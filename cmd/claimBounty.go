//Package cmd provides all functions related to command line
package cmd

import (
	"errors"
	"math/big"
	"os"
	"razor/core"
	"razor/core/types"
	"razor/logger"
	"razor/path"
	"razor/pkg/bindings"
	"razor/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var claimBountyCmd = &cobra.Command{
	Use:   "claimBounty",
	Short: "claim earned bounty",
	Long: `ClaimBounty allows the users who are bountyHunter to redeem their bounty in razor network

Example:
  ./razor claimBounty --address 0x5a0b54d5dc17e0aadc383d2db43b0a0d3e029c4c --bountyId 2 --logFile claimBounty`,
	Run: initialiseClaimBounty,
}

//This function initialises the ExecuteClaimBounty function
func initialiseClaimBounty(cmd *cobra.Command, args []string) {
	cmdUtils.ExecuteClaimBounty(cmd.Flags())
}

//This function sets the flags appropriately and executes the ClaimBounty function
func (*UtilsStruct) ExecuteClaimBounty(flagSet *pflag.FlagSet) {
	config, err := cmdUtils.GetConfigData()
	utils.CheckError("Error in getting config: ", err)
	log.Debugf("ExecuteClaimBounty: config: %+v", config)

	client := razorUtils.ConnectToClient(config.Provider)

	address, err := flagSetUtils.GetStringAddress(flagSet)
	utils.CheckError("Error in getting address: ", err)
	log.Debug("ExecuteClaimBounty: Address: ", address)

	logger.SetLoggerParameters(client, address)

	log.Debug("Checking to assign log file...")
	fileUtils.AssignLogFile(flagSet, config)

	log.Debug("Getting password...")
	password := razorUtils.AssignPassword(flagSet)

	err = utils.CheckPassword(address, password)
	utils.CheckError("Error in fetching private key from given password: ", err)

	if razorUtils.IsFlagPassed("bountyId") {
		bountyId, err := flagSetUtils.GetUint32BountyId(flagSet)
		utils.CheckError("Error in getting bountyId: ", err)
		log.Debug("ExecuteClaimBounty: BountyId: ", bountyId)

		redeemBountyInput := types.RedeemBountyInput{
			Address:  address,
			Password: password,
			BountyId: bountyId,
		}

		log.Debugf("ExecuteClaimBounty: Calling ClaimBounty() with arguments redeem bounty input: %+v", redeemBountyInput)
		txn, err := cmdUtils.ClaimBounty(config, client, redeemBountyInput)
		utils.CheckError("ClaimBounty error: ", err)

		if txn != core.NilHash {
			err = razorUtils.WaitForBlockCompletion(client, txn.Hex())
			utils.CheckError("Error in WaitForBlockCompletion for claimBounty: ", err)
		}
	} else {
		log.Debug("ExecuteClaimBounty: Calling HandleClaimBounty()")
		err := cmdUtils.HandleClaimBounty(client, config, types.Account{
			Address:  address,
			Password: password,
		})
		utils.CheckError("HandleClaimBounty error: ", err)
	}

}

//This function handles claimBounty by picking bountyid's from disputeData file and if there is any error it returns the error
func (*UtilsStruct) HandleClaimBounty(client *ethclient.Client, config types.Configurations, account types.Account) error {
	disputeFilePath, err := pathUtils.GetDisputeDataFileName(account.Address)
	if err != nil {
		return err
	}
	log.Debug("HandleClaimBounty: Dispute data file path: ", disputeFilePath)
	if _, err := path.OSUtilsInterface.Stat(disputeFilePath); !errors.Is(err, os.ErrNotExist) {
		log.Debug("Fetching the dispute data from dispute data file...")
		disputeData, err = fileUtils.ReadFromDisputeJsonFile(disputeFilePath)
		if err != nil {
			return err
		}

		log.Debugf("HandleClaimBounty: DisputeData: %+v", disputeData)
	}

	if disputeData.BountyIdQueue == nil {
		log.Error("No bounty id's present")
		return errors.New("no bounty earned")
	}

	if disputeData.BountyIdQueue != nil {
		log.Info("Bounty ids that needs be claimed: ", disputeData.BountyIdQueue)
		length := len(disputeData.BountyIdQueue)
		log.Info("Claiming bounty for bountyId ", disputeData.BountyIdQueue[length-1])
		redeemBountyInput := types.RedeemBountyInput{
			BountyId: disputeData.BountyIdQueue[length-1],
			Address:  account.Address,
			Password: account.Password,
		}
		log.Debugf("HandleClaimBounty: Calling ClaimBounty() with arguments redeemBountyInput: %+v", redeemBountyInput)
		claimBountyTxn, err := cmdUtils.ClaimBounty(config, client, redeemBountyInput)
		if err != nil {
			return err
		}
		if claimBountyTxn != core.NilHash {
			claimBountyErr := razorUtils.WaitForBlockCompletion(client, claimBountyTxn.Hex())
			if claimBountyErr == nil {
				if len(disputeData.BountyIdQueue) > 1 {
					//Removing the bountyId from the queue as the bounty is being claimed
					disputeData.BountyIdQueue = disputeData.BountyIdQueue[:length-1]
				} else {
					disputeData.BountyIdQueue = nil
				}
			}
		}
	}

	log.Debug("Saving the updated dispute data to dispute data file...")
	err = fileUtils.SaveDataToDisputeJsonFile(disputeFilePath, disputeData.BountyIdQueue)
	if err != nil {
		return err
	}
	return nil
}

//This function allows the users who are bountyHunter to redeem their bounty in razor network
func (*UtilsStruct) ClaimBounty(config types.Configurations, client *ethclient.Client, redeemBountyInput types.RedeemBountyInput) (common.Hash, error) {
	txnArgs := types.TransactionOptions{
		Client:          client,
		AccountAddress:  redeemBountyInput.Address,
		Password:        redeemBountyInput.Password,
		ChainId:         core.ChainId,
		Config:          config,
		ContractAddress: core.StakeManagerAddress,
		ABI:             bindings.StakeManagerMetaData.ABI,
		MethodName:      "redeemBounty",
		Parameters:      []interface{}{redeemBountyInput.BountyId},
	}
	epoch, err := razorUtils.GetEpoch(txnArgs.Client)
	if err != nil {
		log.Error("Error in getting epoch: ", err)
		return core.NilHash, err
	}
	log.Debug("ClaimBounty: Epoch: ", epoch)

	callOpts := razorUtils.GetOptions()
	bountyLock, err := stakeManagerUtils.GetBountyLock(txnArgs.Client, &callOpts, redeemBountyInput.BountyId)
	if err != nil {
		log.Error("Error in getting bounty lock: ", err)
		return core.NilHash, err
	}
	log.Debugf("ClaimBounty: Bounty lock: %+v", bountyLock)

	if bountyLock.Amount.Cmp(big.NewInt(0)) == 0 {
		err = errors.New("bounty amount is 0")
		log.Error(err)
		return core.NilHash, err
	}

	log.Info("Claiming bounty transaction...")
	waitFor := int32(bountyLock.RedeemAfter) - int32(epoch)
	if waitFor > 0 {
		log.Debug("Waiting for lock period to get over....")

		timeRemaining := uint64(waitFor) * core.EpochLength
		if waitFor == 1 {
			log.Infof("Cannot claim bounty now. Please wait for %d epoch! (approximately %s)", waitFor, razorUtils.SecondsToReadableTime(int(timeRemaining)))
		} else {
			log.Infof("Cannot claim bounty now. Please wait for %d epochs! (approximately %s)", waitFor, razorUtils.SecondsToReadableTime(int(timeRemaining)))
		}
		return core.NilHash, nil
	}

	txnOpts := razorUtils.GetTxnOpts(txnArgs)

	log.Debug("Executing RedeemBounty transaction with bountyId: ", redeemBountyInput.BountyId)
	tx, err := stakeManagerUtils.RedeemBounty(txnArgs.Client, txnOpts, redeemBountyInput.BountyId)
	if err != nil {
		return core.NilHash, err
	}
	txnHash := transactionUtils.Hash(tx)
	log.Info("Txn Hash: ", txnHash.Hex())
	return txnHash, nil
}

func init() {
	rootCmd.AddCommand(claimBountyCmd)
	var (
		Address  string
		Password string
		BountyId uint32
	)

	claimBountyCmd.Flags().StringVarP(&Address, "address", "a", "", "address of the staker")
	claimBountyCmd.Flags().StringVarP(&Password, "password", "", "", "password path of staker to protect the keystore")
	claimBountyCmd.Flags().Uint32VarP(&BountyId, "bountyId", "", 0, "bountyId of the bounty hunter")

	addrErr := claimBountyCmd.MarkFlagRequired("address")
	utils.CheckError("Address error: ", addrErr)

}
