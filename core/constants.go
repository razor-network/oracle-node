package core

import "math/big"

//TODO: Change these addresses to the deployed address.
var StakeManagerAddress = "0x70C55e9f008f66fB67F508b02fdcac81916826e6"
var RAZORAddress = "0x34A8759E4024A9A217014DB7Ee046498A16f08f0"
var ParametersAddress = "0xcCb4043db41D6aDaFc770FBC706c9E3194524287"
var AssetManagerAddress = "0xD62408a3aECD9aC97878B216979790b3C3b543c4"
var VoteManagerAddress = "0x6e83e9C80B34eB41e8FE252c075B5713558Cc0bc"
var RandomClientAddress = "0x2CCc74005be81213c9322EE56AAE2273525101c2"
var BlockManagerAddress = "0x2e19aB0b134Ae1917A44C3C3338fA774cCB4c079"

var StateLength uint64 = 75
var EpochLength int64 = 300
var NumberOfStates int64 = 4
var DecimalsMultiplier int64 = 100000000
var NumberOfBlocks = 10
var ChainId = big.NewInt(80001)
