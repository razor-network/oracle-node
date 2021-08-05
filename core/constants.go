package core

import "math/big"

//TODO: Change these addresses to the deployed address.
var StakeManagerAddress = "0x809A7d7cA60A3a0C3dD7dbf2Ba6f56e9b09156d7"
var RAZORAddress = "0x9CfA4A2D0E705B1C9dA418EB901f51817512D2dd"
var ParametersAddress = "0xf85B5Af15a9a4c417603Afb5809F308aFb484339"
var AssetManagerAddress = "0x2494FcCada1B6532a9f705155fDC02b1BEfcb13a"
var VoteManagerAddress = "0x9561518cA5a293Ea636a40148e76e9F11dAAeab5"
var RandomClientAddress = "0xb3A395A36c469A33d9F078EB1A1c602A11Dad7a7"
var BlockManagerAddress = "0x099272194FD75656dd013BD21d1eF6E987a14C13"

var StateLength uint64 = 75
var EpochLength int64 = 300
var NumberOfStates int64 = 4
var DecimalsMultiplier int64 = 100000000
var NumberOfBlocks = 10
var ChainId = big.NewInt(80001)
