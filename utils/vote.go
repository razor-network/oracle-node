package utils

import (
	"github.com/avast/retry-go"
	"math/big"
	"razor/core"
	"razor/core/types"
	"razor/pkg/bindings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

func (*UtilsStruct) GetVoteManagerWithOpts(client *ethclient.Client) (*bindings.VoteManager, bind.CallOpts) {
	return UtilsInterface.GetVoteManager(client), UtilsInterface.GetOptions()
}

func (*UtilsStruct) GetCommitments(client *ethclient.Client, address string) ([32]byte, error) {
	callOpts := UtilsInterface.GetOptions()
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
			commitments, commitmentErr = Options.Commitments(client, &callOpts, stakerId)
			if commitmentErr != nil {
				log.Error("Error in fetching commitment....Retrying")
				return commitmentErr
			}
			return nil
		}, Options.RetryAttempts(core.MaxRetries))
	if commitmentErr != nil {
		return [32]byte{}, err
	}
	return commitments.CommitmentHash, nil
}

func (*UtilsStruct) GetVoteValue(client *ethclient.Client, assetId uint16, stakerId uint32) (*big.Int, error) {
	callOpts := UtilsInterface.GetOptions()
	var (
		voteValue    *big.Int
		voteValueErr error
	)
	voteValueErr = retry.Do(
		func() error {
			voteValue, voteValueErr = Options.GetVoteValue(client, &callOpts, assetId-1, stakerId)
			if voteValueErr != nil {
				log.Error("Error in fetching last vote value....Retrying")
				return voteValueErr
			}
			return nil
		}, Options.RetryAttempts(core.MaxRetries))
	if voteValueErr != nil {
		return nil, voteValueErr
	}
	return voteValue, nil
}

func (*UtilsStruct) GetInfluenceSnapshot(client *ethclient.Client, stakerId uint32, epoch uint32) (*big.Int, error) {
	callOpts := UtilsInterface.GetOptions()
	var (
		influenceSnapshot *big.Int
		influenceErr      error
	)
	influenceErr = retry.Do(
		func() error {
			influenceSnapshot, influenceErr = Options.GetInfluenceSnapshot(client, &callOpts, epoch, stakerId)
			if influenceErr != nil {
				log.Error("Error in fetching influence snapshot....Retrying")
				return influenceErr
			}
			return nil
		}, Options.RetryAttempts(core.MaxRetries))
	if influenceErr != nil {
		return nil, influenceErr
	}
	return influenceSnapshot, nil
}

func (*UtilsStruct) GetStakeSnapshot(client *ethclient.Client, stakerId uint32, epoch uint32) (*big.Int, error) {
	callOpts := UtilsInterface.GetOptions()
	var (
		stakeSnapshot *big.Int
		snapshotErr   error
	)
	snapshotErr = retry.Do(
		func() error {
			stakeSnapshot, snapshotErr = Options.GetStakeSnapshot(client, &callOpts, epoch, stakerId)
			if snapshotErr != nil {
				log.Error("Error in fetching stake snapshot....Retrying")
				return snapshotErr
			}
			return nil
		}, Options.RetryAttempts(core.MaxRetries))
	if snapshotErr != nil {
		return nil, snapshotErr
	}
	return stakeSnapshot, nil
}

func (*UtilsStruct) GetTotalInfluenceRevealed(client *ethclient.Client, epoch uint32) (*big.Int, error) {
	callOpts := UtilsInterface.GetOptions()
	var (
		totalInfluenceRevealed *big.Int
		influenceErr           error
	)
	influenceErr = retry.Do(
		func() error {
			totalInfluenceRevealed, influenceErr = Options.GetTotalInfluenceRevealed(client, &callOpts, epoch)
			if influenceErr != nil {
				log.Error("Error in fetching total influence revealed....Retrying")
				return influenceErr
			}
			return nil
		}, Options.RetryAttempts(core.MaxRetries))
	if influenceErr != nil {
		return nil, influenceErr
	}
	return totalInfluenceRevealed, nil
}

func (*UtilsStruct) GetEpochLastCommitted(client *ethclient.Client, stakerId uint32) (uint32, error) {
	callOpts := UtilsInterface.GetOptions()
	var (
		epochLastCommitted uint32
		err                error
	)
	err = retry.Do(
		func() error {
			epochLastCommitted, err = Options.GetEpochLastCommitted(client, &callOpts, stakerId)
			if err != nil {
				log.Error("Error in fetching epoch last committed....Retrying")
				return err
			}
			return nil
		}, Options.RetryAttempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return epochLastCommitted, nil
}

func (*UtilsStruct) GetEpochLastRevealed(client *ethclient.Client, stakerId uint32) (uint32, error) {
	callOpts := UtilsInterface.GetOptions()
	var (
		epochLastRevealed uint32
		err               error
	)
	err = retry.Do(
		func() error {
			epochLastRevealed, err = Options.GetEpochLastRevealed(client, &callOpts, stakerId)
			if err != nil {
				log.Error("Error in fetching epoch last revealed....Retrying")
				return err
			}
			return nil
		}, Options.RetryAttempts(core.MaxRetries))
	if err != nil {
		return 0, err
	}
	return epochLastRevealed, nil
}
