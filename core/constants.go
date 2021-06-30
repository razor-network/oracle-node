package core

import "math/big"

//TODO: Change these addresses to the deployed address.
var StakeManagerAddress = "0xBfBA8b5F2CaD3Fdd82893BECDBfb0193C878CED6"
var SchellingCoinAddress = "0x1Ca31A8832c532DD22AcFB0a3cf7b9F87F4574e6"
var ParametersAddress = "0x1B08F228697A195fB0406A8EB24C3A4E35209c1A"
var AssetManagerAddress = "0xC5edba1DdA6FF3D7682284Da6eeb65b793704a5c"
var VoteManagerAddress = "0xA47b2F152628E2865A42153d59dFa4336E94be50"
var RandomClientAddress = "0xc470b4254550700eA75dbdf945eaf4f5c571da04"
var BlockManagerAddress = "0x36CD33dbC84920afB057DE9CD9fC9edB446d338D"

var StateLength uint64 = 75
var EpochLength int64 = 300
var NumberOfStates int64 = 4
var DecimalsMultiplier int64 = 100000000
var NumberOfBlocks = 10
var ChainId = big.NewInt(80001)
