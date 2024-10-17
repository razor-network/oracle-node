package types

import (
	"math/big"
	"net/http"
	"razor/cache"
)

type ElectedProposer struct {
	Iteration       int
	Stake           *big.Int
	StakerId        uint32
	BiggestStake    *big.Int
	NumberOfStakers uint32
	Salt            [32]byte
	Epoch           uint32
}

type Commitment struct {
	Epoch          uint32
	CommitmentHash [32]byte
}

type Rogue struct {
	IsRogue   bool
	RogueMode []string
}

type CommitData struct {
	AssignedCollections    map[int]bool
	SeqAllottedCollections []*big.Int
	Leaves                 []*big.Int
}

type RevealedStruct struct {
	RevealedValues []AssignedAsset
	Influence      *big.Int
}

type RevealedDataMaps struct {
	SortedRevealedValues map[uint16][]*big.Int
	VoteWeights          map[string]*big.Int
	InfluenceSum         map[uint16]*big.Int
}

type DisputeFileData struct {
	BountyIdQueue []uint32
}

type ProposeData struct {
	MediansData           []*big.Int
	RevealedCollectionIds []uint16
	RevealedDataMaps      *RevealedDataMaps
}

type CommitFileData struct {
	Epoch                  uint32
	AssignedCollections    map[int]bool
	SeqAllottedCollections []*big.Int
	Leaves                 []*big.Int
	Commitment             [32]byte
}

type ProposeFileData struct {
	Epoch                 uint32
	MediansData           []*big.Int
	RevealedCollectionIds []uint16
	RevealedDataMaps      *RevealedDataMaps
}

type CommitParams struct {
	JobsCache                 *cache.JobsCache
	CollectionsCache          *cache.CollectionsCache
	LocalCache                *cache.LocalCache
	HttpClient                *http.Client
	FromBlockToCheckForEvents *big.Int
}
