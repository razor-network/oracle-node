package types

type Configurations struct {
	Provider           string
	GasMultiplier      float32
	BufferPercent      int32
	WaitTime           int32
	GasPrice           int32
	LogLevel           string
	GasLimitMultiplier float32
	RPCTimeout         int64
	HTTPTimeout        int64
	LogFileMaxSize     int
	LogFileMaxBackups  int
	LogFileMaxAge      int
}
