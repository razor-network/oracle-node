package core

import "math/big"

//TODO: Change these addresses to the deployed address.
var StakeManagerAddress = "0x58bA893aEc9fcbd8A2e992027bDC643041d83CF7"
var RAZORAddress = "0x1aD3fDE97626d13d36b85f7E16b2CbB887787Aec"
var ParametersAddress = "0xa4833a98Fd501C17042D0C2f0fE314e7Ea4C07bb"
var AssetManagerAddress = "0x76c5b2bA88697f22e02398a5d8474f5bfe452b3F"
var VoteManagerAddress = "0xA9120C9b39ACbEA0e016cA536DCFE28Ac3ABf40D"
var RandomClientAddress = "0x65df259491AF2049f148cfc4267a3EA3fB7fCd23"
var BlockManagerAddress = "0x3c55ED741cB7DE68a05AFa3Dca5112653dd18455"

var StateLength uint64 = 75
var EpochLength int64 = 300
var NumberOfStates int64 = 4
var DecimalsMultiplier int64 = 100000000
var NumberOfBlocks = 10
var ChainId = big.NewInt(80001)
