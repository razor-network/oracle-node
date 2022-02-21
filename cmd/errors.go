package cmd

import "fmt"

//Flag error codes
const ConfigErrorCode = 1001
const AddressErrorCode = 1002
const FetchBalanceErrorCode = 1003
const AmountErrorCode = 1004
const RogueStatusErrorCode = 1007
const RogueModeErrorCode = 1008
const AutoVoteStatusErrorCode = 1009
const BountyIdErrorCode = 1010

//Functionality error codes
const ApproveErrorCode = 1020
const StakeErrorCode = 1021
const ClaimBountyErrorCode = 1022
const FetchEpochErrorCode = 1023

//Flag error messages
const ConfigErrorMessage = "Error in getting config:"
const AddressErrorMessage = "Error in fetching address:"
const FetchBalanceErrorMessage = "Error in fetching balance for account:"
const AmountErrorMessage = "Error in getting amount:"
const RogueStatusErrorMessage = "Error in getting rogue status:"
const RogueModeErrorMessage = "Error in getting rogue mode:"
const AutoVoteStatusErrorMessage = "Error in getting autoVote status:"
const BountyIdErrorMessage = "Error in getting bountyId:"

//Functionality error messages
const ApproveErrorMessage = "Approve error:"
const StakeErrorMessage = "Stake error:"
const ClaimBountyErrorMessage = "ClaimBounty error:"
const FetchEpochErrorMessage = "Error in getting epoch:"

var GetError = func(errCode int, errMessage string) string {
	return fmt.Sprintf("ERR: %d\t%s ", errCode, errMessage)
}
