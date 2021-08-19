package utils

import "razor/core"

func GetStakeManagerAddress() string {
	addresses := core.AssignAddressesFromJSON()
	return addresses.StakeManagerAddress
}

func GetAssetManagerAddress() string {
	addresses := core.AssignAddressesFromJSON()
	return addresses.AssetManagerAddress
}

func GetVoteManagerAddress() string {
	addresses := core.AssignAddressesFromJSON()
	return addresses.VoteManagerAddress
}

func GetRAZORAddress() string {
	addresses := core.AssignAddressesFromJSON()
	return addresses.RAZORAddress
}

func GetBlockManagerAddress() string {
	addresses := core.AssignAddressesFromJSON()
	return addresses.BlockManagerAddress
}

func GetParametersAddress() string {
	addresses := core.AssignAddressesFromJSON()
	return addresses.ParametersAddress
}
