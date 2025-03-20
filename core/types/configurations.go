package types

type Configurations struct {
	Provider           string
	GasMultiplier      float32
	BufferPercent      int32
	WaitTime           int32
	GasPrice           int32
	LogLevel           string
	GasLimitOverride   uint64
	GasLimitMultiplier float32
	RPCTimeout         int64
	HTTPTimeout        int64
	LogFileMaxSize     int
	LogFileMaxBackups  int
	LogFileMaxAge      int
}

type ConfigDetail struct {
	FlagName     string
	Key          string
	DefaultValue interface{}
}
