package utils

import (
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"

	"github.com/avast/retry-go"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func (*UtilsStruct) GetVoteManagerWithOpts(client *ethclient.Client) (*bindings.VoteManager, bind.CallOpts) {
	return UtilsInterface.GetVoteManager(client), UtilsInterface.GetOptions()
}

func (*UtilsStruct) GetCommitments(client *ethclient.Client, address string) ([32]byte, error) {
	stakerId, err := UtilsInterface.GetStakerId(client, address)
	if err != nil {
		return [32]byte{}, err
	}
	var (
		commitments   types.Commitment
		commitmentErr error
	)
	commitmentErr = retry.Do(
		func() error {
			commitments, commitmentErr = VoteManagerInterface.Commitments(client, stakerId)
			if commitmentErr != nil {
				log.Error("Error in fetching commitment....Retrying")
				return commitmentErr
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if commitmentErr != nil {
		return [32]byte{}, err
	}
	return commitments.CommitmentHash, nil
}

func (*UtilsStruct) GetVoteValue(client *ethclient.Client, epoch uint32, stakerId uint32, medianIndex uint16) (*big.Int, error) {
	var (
		voteValue    uint32
		voteValueErr error
	)
	voteValueErr = retry.Do(
		func() error {
			voteValue, voteValueErr = VoteManagerInterface.GetVoteValue(client, epoch, stakerId, medianIndex)
			if voteValueErr != nil {
				log.Error("Error in fetching last vote value....Retrying")
				return voteValueErr
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if voteValueErr != nil {
		return big.NewInt(0), voteValueErr
	}
	return voteValue, nil
}

func (*UtilsStruct) GetInfluenceSnapshot(client *ethclient.Client, stakerId uint32, epoch uint32) (*big.Int, error) {
	var (
		influenceSnapshot *big.Int
		influenceErr      error
	)
	influenceErr = retry.Do(
		func() error {
			influenceSnapshot, influenceErr = VoteManagerInterface.GetInfluenceSnapshot(client, epoch, stakerId)
			if influenceErr != nil {
				log.Error("Error in fetching influence snapshot....Retrying")
				return influenceErr
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if influenceErr != nil {
		return nil, influenceErr
	}
	return influenceSnapshot, nil
}

func (*UtilsStruct) GetStakeSnapshot(client *ethclient.Client, stakerId uint32, epoch uint32) (*big.Int, error) {
	var (
		stakeSnapshot *big.Int
		snapshotErr   error
	)
	snapshotErr = retry.Do(
		func() error {
			stakeSnapshot, snapshotErr = VoteManagerInterface.GetStakeSnapshot(client, epoch, stakerId)
			if snapshotErr != nil {
				log.Error("Error in fetching stake snapshot....Retrying")
				return snapshotErr
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if snapshotErr != nil {
		return nil, snapshotErr
	}
	return stakeSnapshot, nil
}

func (*UtilsStruct) GetTotalInfluenceRevealed(client *ethclient.Client, epoch uint32, medianIndex uint16) (*big.Int, error) {
	var (
		totalInfluenceRevealed *big.Int
		influenceErr           error
	)
	influenceErr = retry.Do(
		func() error {
			totalInfluenceRevealed, influenceErr = VoteManagerInterface.GetTotalInfluenceRevealed(client, epoch, medianIndex)
			if influenceErr != nil {
				log.Error("Error in fetching total influence revealed....Retrying")
				return influenceErr
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if influenceErr != nil {
		return nil, influenceErr
	}
	return totalInfluenceRevealed, nil
}

func (*UtilsStruct) GetEpochLastCommitted(client *ethclient.Client, stakerId uint32) (uint32, error) {
	var (
		epochLastCommitted uint32
		err                error
	)
	err = retry.Do(
		func() error {
			epochLastCommitted, err = VoteManagerInterface.GetEpochLastCommitted(client, stakerId)
			if err != nil {
				log.Error("Error in fetching epoch last committed....Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return epochLastCommitted, nil
}

func (*UtilsStruct) GetEpochLastRevealed(client *ethclient.Client, stakerId uint32) (uint32, error) {
	var (
		epochLastRevealed uint32
		err               error
	)
	err = retry.Do(
		func() error {
			epochLastRevealed, err = VoteManagerInterface.GetEpochLastRevealed(client, stakerId)
			if err != nil {
				log.Error("Error in fetching epoch last revealed....Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return epochLastRevealed, nil
}

func (*UtilsStruct) ToAssign(client *ethclient.Client) (uint16, error) {
	var (
		toAssign uint16
		err      error
	)
	err = retry.Do(
		func() error {
			toAssign, err = VoteManagerInterface.ToAssign(client)
			if err != nil {
				log.Error("Error in fetching toAssign....Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return toAssign, nil
}

func (*UtilsStruct) GetSaltFromBlockchain(client *ethclient.Client) ([32]byte, error) {
	var (
		salt [32]byte
		err  error
	)
	err = retry.Do(
		func() error {
			salt, err = VoteManagerInterface.GetSaltFromBlockchain(client)
			if err != nil {
				log.Error("Error in fetching salt....Retrying")
				return err
			}
			return nil
		}, RetryInterface.RetryAttempts(core.MaxRetries))
	if err != nil {
		return [32]byte{}, err
	}
	return salt, nil
}
