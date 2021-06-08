package core

import "math/big"

var StakeManagerAddress = "0xC8C86b586551ED4F1c1A1B2617B305be7637e5B2"
var SchellingCoinAddress = "0x8b4F6FE4E110C9509d77A7205Ad85a224d9c3E68"
var StateManagerAddress = "0xDdBb10a49A54172FdBa02b93D6D1dDfC12A838ba"
var ConstantsAddress = "0xeE7Dc215BEAEcf2c6f4B30E19b89bb580E404EDa"
var JobManagerAddress = "0x28c2c5725ED40D5faD7bdA2419A4753BC6946A9b"
var VoteManagerAddress = "0x246b61043839479E9d070506766bFbfBe659e0b4"
var RandomClientAddress = "0xcf8F791945346E6daB85d2C140911A99E409768C"
var BlockManagerAddress = "0x99E7Afeb0215D2E80088b8C7Bde54BD9162c3A35"

var EpochLength uint64 = 75
var NumberOfStates int64 = 4
var DecimalsMultiplier int64 = 100000000
var NumberOfBlocks = 10
var ChainId = big.NewInt(80001)