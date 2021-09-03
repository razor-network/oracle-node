package core

import "math/big"

//TODO: Change these addresses to the deployed address.
var StakeManagerAddress = "0x92d5fB140fB93E0790baD45f945677ABBC0239c3"
var RAZORAddress = "0xA02C76D746877d3bB5df1fD3c4130D642c16ab49"
var ParametersAddress = "0x31E62E022f241bfE8c3Bc83DE8c090af7a67DdD6"
var AssetManagerAddress = "0x80280d1D25558Bd23712720266428Ba3AfaD2e0B"
var VoteManagerAddress = "0xD0dcD208Cf140cC160E8e5423a9d76458c58AC6E"
var BlockManagerAddress = "0x8eAC317759167b0901cf0b2548a91321F62B3f2B"

var StateLength uint64 = 75
var EpochLength int64 = 375
var NumberOfStates int64 = 5
var ChainId = big.NewInt(80001)
