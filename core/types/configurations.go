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
	LogFileMaxSize     int
	LogFileMaxBackups  int
	LogFileMaxAge      int
}
