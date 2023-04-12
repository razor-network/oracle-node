package client

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"razor/logger"
	"reflect"
	"time"
)

var log = logger.NewLogger()

var SwitchClientToAlternateClient bool
var AlternateProvider string

func StartTimerForAlternateClient(switchClientAfterTime uint64) {
	log.Infof("StartTimerForAlternateClient: Alternate client will be switched back to primary client in %v seconds!", switchClientAfterTime)
	time.Sleep(time.Duration(switchClientAfterTime) * time.Second)
	log.Info("Switching back to primary RPC..")
	SwitchClientToAlternateClient = false
}

//ReplaceClientWithAlternateClient will replace the primary client(client from primary RPC) with secondary client which would be created using alternate RPC
func ReplaceClientWithAlternateClient(arguments []reflect.Value) []reflect.Value {
	clientDataType := reflect.TypeOf((*ethclient.Client)(nil)).Elem()
	for i := range arguments {
		argument := arguments[i]
		argumentDataType := reflect.TypeOf(argument.Interface()).Elem()
		if argumentDataType != nil {
			if argumentDataType == clientDataType {
				alternateClient, dialErr := ethclient.Dial(AlternateProvider)
				if dialErr != nil {
					log.Errorf("Error in connecting using alternate RPC %v: %v", AlternateProvider, dialErr)
					return arguments
				}
				arguments[i] = reflect.ValueOf(alternateClient)
				return arguments
			}
		}
	}
	return arguments
}
