package core

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

//TODO: Change these addresses to the deployed address.
var StakeManagerAddress = "0x6133d9f3Dd76C2D405652bF284728b8AAB665d1b"
var RAZORAddress = "0xc89FadB7A04813f4d68D39363ECa86cb232897c3"
var ParametersAddress = "0xC463712DF07c4401665Fd7765915007C08d4182a"
var AssetManagerAddress = "0x4b0e123589bAaD54D776aE68d6A9357804e92e70"
var VoteManagerAddress = "0xbD7945A5B9Fb1697995f715bE82930977707eeA8"
var BlockManagerAddress = "0xbDa74791465aEA005eF001b74521867d61606D5d"

var StateLength uint64 = 60
var EpochLength int64 = 300

var NumberOfStates int64 = 5
var ChainId = big.NewInt(80001)
var MaxRetries = 3
var NilHash = common.Hash{0x00}
