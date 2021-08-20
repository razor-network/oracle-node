package core

import "math/big"

//TODO: Change these addresses to the deployed address.
var StakeManagerAddress = "0x4E648ec2d2e248c6621F95973780d24d3D9F525A"
var RAZORAddress = "0x7DfD56634D999bC28092Fe35De027c5B65958aF1"
var ParametersAddress = "0xc51621ee6e7613Ea6B67034eB780b16f75FABD8E"
var AssetManagerAddress = "0x15613697F76E78aFcb435aDa9De62469C2a5c9c3"
var VoteManagerAddress = "0xa64B36d852d7B62074d8824E8DC7412acACFC1f2"
var BlockManagerAddress = "0x2F3ACb16397aa8662874F8bD115E7B47E86730b3"

var StateLength uint64 = 75
var EpochLength int64 = 300
var NumberOfStates int64 = 4
var DecimalsMultiplier int64 = 100000000
var NumberOfBlocks = 10
var ChainId = big.NewInt(80001)
