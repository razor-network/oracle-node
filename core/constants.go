package core

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/big"
)

type Addresses struct {
	StakeManagerAddress string `json:"StakeManager"`
	RAZORAddress        string `json:"RAZOR"`
	ParametersAddress   string `json:"Parameters"`
	AssetManagerAddress string `json:"AssetManager"`
	VoteManagerAddress  string `json:"VoteManager"`
	BlockManagerAddress string `json:"BlockManager"`
}

func AssignAddressesFromJSON() Addresses {
	var addresses Addresses
	data, err := ioutil.ReadFile("./../deployed_addresses/addresses.json")
	if err != nil {
		log.Fatal(err)
	}

	unmarshalErr := json.Unmarshal(data, &addresses)
	if unmarshalErr != nil {
		log.Fatal(unmarshalErr)
	}
	return addresses
}

var StateLength uint64 = 75
var EpochLength int64 = 300
var NumberOfStates int64 = 4
var DecimalsMultiplier int64 = 100000000
var NumberOfBlocks = 10
var ChainId = big.NewInt(80001)
