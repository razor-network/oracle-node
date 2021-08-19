package core

import "math/big"

//TODO: Change these addresses to the deployed address.
var StakeManagerAddress = "0x5B4a0dF703d1D5Ea5c1605f24a8572449b4F22B5"
var RAZORAddress = "0xB3c9650357B358ddf8d60167b4B305D7fd0676c3"
var ParametersAddress = "0x4F629Aa4CB67fbDbcDDB508fc16ED4BD7f25E751"
var AssetManagerAddress = "0xA43AfAA0F3B4D4109Cc0949E95C57111F5fE6993"
var VoteManagerAddress = "0x76CFd554C5ed222bea7A2c0d23631a47fa3bc128"
var BlockManagerAddress = "0xD63c084F2B2eC5688fbc0A42D3AbfdB0cd7A2535"

var StateLength uint64 = 75
var EpochLength int64 = 300
var NumberOfStates int64 = 4
var DecimalsMultiplier int64 = 100000000
var NumberOfBlocks = 10
var ChainId = big.NewInt(80001)
