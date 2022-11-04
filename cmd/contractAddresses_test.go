package cmd

import (
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/mock"
	"razor/cmd/mocks"
	utilsPkgMocks "razor/utils/mocks"
	"testing"
)

func TestContractAddresses(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "When ContractAddresses() executes successfully",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ut := &UtilsStruct{}
			ut.ContractAddresses()
		})
	}
}

func TestExecuteContractAddresses(t *testing.T) {
	var flagSet *pflag.FlagSet
	tests := []struct {
		name          string
		expectedFatal bool
	}{
		{
			name:          "When ExecuteContractAddresses() executes successfully",
			expectedFatal: false,
		},
	}

	defer func() { log.ExitFunc = nil }()
	var fatal bool
	log.ExitFunc = func(int) { fatal = true }
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			utilsMock := new(utilsPkgMocks.Utils)
			cmdUtilsMock := new(mocks.UtilsCmdInterface)

			razorUtils = utilsMock
			cmdUtils = cmdUtilsMock

			utilsMock.On("AssignLogFile", mock.AnythingOfType("*pflag.FlagSet"), mock.Anything)
			cmdUtilsMock.On("ContractAddresses")

			utils := &UtilsStruct{}
			fatal = false

			utils.ExecuteContractAddresses(flagSet)
			if fatal != tt.expectedFatal {
				t.Error("The ExecuteContractAddresses function didn't execute as expected")
			}
		})
	}
}
