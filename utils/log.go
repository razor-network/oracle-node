package utils

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"razor/block"
	"razor/core/types"
	"razor/logger"
)

var log = logger.NewLogger("", &ethclient.Client{}, "", types.Configurations{}, &block.BlockMonitor{})

func UpdateUtilsLogger(updatedLogger *logger.Logger) {
	log = updatedLogger
}
