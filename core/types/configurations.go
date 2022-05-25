//Package types include the different user defined items of possible different types in a single type
package types

type Configurations struct {
	Provider           string
	ChainId            int64
	GasMultiplier      float32
	BufferPercent      int32
	WaitTime           int32
	GasPrice           int32
	LogLevel           string
	GasLimitMultiplier float32
}
