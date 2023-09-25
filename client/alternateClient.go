package client

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"razor/logger"
	"reflect"
	"time"
)

var (
	log                   = logger.NewLogger()
	alternateClientStruct AlternateClientStruct
)

type AlternateClientStruct struct {
	switchToAlternateClient bool
	alternateProvider       string
}

func StartTimerForAlternateClient(switchClientAfterTime uint64) {
	log.Infof("StartTimerForAlternateClient: Alternate client will be switched back to primary client in %v seconds!", switchClientAfterTime)
	time.Sleep(time.Duration(switchClientAfterTime) * time.Second)
	log.Info("Switching back to primary RPC..")
	SetSwitchToAlternateClientStatus(false)
}

// ReplaceClientWithAlternateClient will replace the primary client(client from primary RPC) with secondary client which would be created using alternate RPC
func ReplaceClientWithAlternateClient(arguments []reflect.Value) []reflect.Value {
	clientDataType := reflect.TypeOf((*ethclient.Client)(nil)).Elem()
	for i := range arguments {
		argument := arguments[i]
		argumentDataType := reflect.TypeOf(argument.Interface()).Elem()
		if argumentDataType != nil {
			if argumentDataType == clientDataType {
				alternateProvider := GetAlternateProvider()
				alternateClient, dialErr := ethclient.Dial(alternateProvider)
				if dialErr != nil {
					log.Errorf("Error in connecting using alternate RPC %v: %v", alternateProvider, dialErr)
					return arguments
				}
				arguments[i] = reflect.ValueOf(alternateClient)
				return arguments
			}
		}
	}
	return arguments
}

func GetSwitchToAlternateClientStatus() bool {
	return alternateClientStruct.switchToAlternateClient
}

func SetSwitchToAlternateClientStatus(status bool) {
	alternateClientStruct.switchToAlternateClient = status
}

func GetAlternateProvider() string {
	return alternateClientStruct.alternateProvider
}

func SetAlternateProvider(alternateProvider string) {
	alternateClientStruct.alternateProvider = alternateProvider
}
