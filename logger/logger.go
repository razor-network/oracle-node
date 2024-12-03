package logger

import (
	"errors"
	"fmt"
	"io"
	"math/big"
	"os"
	"razor/block"
	"razor/core"
	"razor/core/types"
	"razor/path"
	"runtime"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/razor-network/goInfo"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger encapsulates logrus and its dependencies for contextual logging.
type Logger struct {
	LogrusInstance *logrus.Logger
	address        string
	client         *ethclient.Client
	blockMonitor   *block.BlockMonitor
}

// Global logger instance
var globalLogger = NewLogger("", nil, nil)

// GetLogger returns the global logger instance.
func GetLogger() *Logger {
	return globalLogger
}

// UpdateLogger updates the global logger instance with new parameters.
func UpdateLogger(address string, client *ethclient.Client, blockMonitor *block.BlockMonitor) {
	globalLogger.address = address
	globalLogger.client = client
	globalLogger.blockMonitor = blockMonitor
}

func init() {
	InitializeLogger("", types.Configurations{})

	osInfo := goInfo.GetInfo()
	globalLogger.LogrusInstance.WithFields(logrus.Fields{
		"Operating System": osInfo.OS,
		"Core":             osInfo.Core,
		"Platform":         osInfo.Platform,
		"CPUs":             osInfo.CPUs,
		"razor-go version": core.VersionWithMeta,
		"go version":       runtime.Version(),
	}).Info()

}

//NewLogger initializes a Logger instance with given parameters.
func NewLogger(address string, client *ethclient.Client, blockMonitor *block.BlockMonitor) *Logger {
	logger := &Logger{
		LogrusInstance: logrus.New(),
		address:        address,
		client:         client,
		blockMonitor:   blockMonitor,
	}

	return logger
}

func InitializeLogger(fileName string, config types.Configurations) {
	path.PathUtilsInterface = &path.PathUtils{}
	path.OSUtilsInterface = &path.OSUtils{}
	if fileName != "" {
		logFilePath, err := path.PathUtilsInterface.GetLogFilePath(fileName)
		if err != nil {
			globalLogger.Fatal("Error in fetching log file path: ", err)
		}

		lumberJackLogger := &lumberjack.Logger{
			Filename:   logFilePath,
			MaxSize:    config.LogFileMaxSize,
			MaxBackups: config.LogFileMaxBackups,
			MaxAge:     config.LogFileMaxAge,
		}

		mw := io.MultiWriter(os.Stderr, lumberJackLogger)
		globalLogger.LogrusInstance.SetOutput(mw)
	}
	globalLogger.LogrusInstance.Formatter = &logrus.JSONFormatter{}
}

// Error logs a simple error message.
func (l *Logger) Error(args ...interface{}) {
	epoch, blockNumber := l.updateBlockInfo()
	logFields := logrus.Fields{
		"address":     l.address,
		"epoch":       epoch,
		"blockNumber": blockNumber,
		"version":     core.VersionWithMeta,
	}
	l.LogrusInstance.WithFields(logFields).Errorln(args...)
}

// Info logs a simple informational message.
func (l *Logger) Info(args ...interface{}) {
	epoch, blockNumber := l.updateBlockInfo()
	logFields := logrus.Fields{
		"address":     l.address,
		"epoch":       epoch,
		"blockNumber": blockNumber,
		"version":     core.VersionWithMeta,
	}
	l.LogrusInstance.WithFields(logFields).Infoln(args...)
}

// Debug logs a simple debug message.
func (l *Logger) Debug(args ...interface{}) {
	epoch, blockNumber := l.updateBlockInfo()
	logFields := logrus.Fields{
		"address":     l.address,
		"epoch":       epoch,
		"blockNumber": blockNumber,
		"version":     core.VersionWithMeta,
	}
	l.LogrusInstance.WithFields(logFields).Debugln(args...)
}

// Fatal logs a fatal error message and exits the application.
func (l *Logger) Fatal(args ...interface{}) {
	epoch, blockNumber := l.updateBlockInfo()
	logFields := logrus.Fields{
		"address":     l.address,
		"epoch":       epoch,
		"blockNumber": blockNumber,
		"version":     core.VersionWithMeta,
	}
	errMsg := joinString(args)
	err := errors.New(errMsg)
	l.LogrusInstance.WithFields(logFields).Fatalln(err)
}

// Warn logs a simple warning message.
func (l *Logger) Warn(args ...interface{}) {
	epoch, blockNumber := l.updateBlockInfo()
	logFields := logrus.Fields{
		"address":     l.address,
		"epoch":       epoch,
		"blockNumber": blockNumber,
		"version":     core.VersionWithMeta,
	}
	l.LogrusInstance.WithFields(logFields).Warnln(args...)
}

// Errorf logs a formatted error message.
func (l *Logger) Errorf(format string, args ...interface{}) {
	epoch, blockNumber := l.updateBlockInfo()
	logFields := logrus.Fields{
		"address":     l.address,
		"epoch":       epoch,
		"blockNumber": blockNumber,
		"version":     core.VersionWithMeta,
	}
	l.LogrusInstance.WithFields(logFields).Errorf(format, args...)
}

// Infof logs a formatted informational message.
func (l *Logger) Infof(format string, args ...interface{}) {
	epoch, blockNumber := l.updateBlockInfo()
	logFields := logrus.Fields{
		"address":     l.address,
		"epoch":       epoch,
		"blockNumber": blockNumber,
		"version":     core.VersionWithMeta,
	}
	l.LogrusInstance.WithFields(logFields).Infof(format, args...)
}

// Debugf logs a formatted debug message.
func (l *Logger) Debugf(format string, args ...interface{}) {
	epoch, blockNumber := l.updateBlockInfo()
	logFields := logrus.Fields{
		"address":     l.address,
		"epoch":       epoch,
		"blockNumber": blockNumber,
		"version":     core.VersionWithMeta,
	}
	l.LogrusInstance.WithFields(logFields).Debugf(format, args...)
}

// Fatalf logs a formatted fatal error message and exits the application.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	epoch, blockNumber := l.updateBlockInfo()
	logFields := logrus.Fields{
		"address":     l.address,
		"epoch":       epoch,
		"blockNumber": blockNumber,
		"version":     core.VersionWithMeta,
	}
	errMsg := joinString(args...)
	err := errors.New(errMsg)
	l.LogrusInstance.WithFields(logFields).Fatalf(format, err)
}

// Warnf logs a formatted warning message.
func (l *Logger) Warnf(format string, args ...interface{}) {
	epoch, blockNumber := l.updateBlockInfo()
	logFields := logrus.Fields{
		"address":     l.address,
		"epoch":       epoch,
		"blockNumber": blockNumber,
		"version":     core.VersionWithMeta,
	}
	l.LogrusInstance.WithFields(logFields).Warnf(format, args...)
}

// joinString concatenates multiple arguments into a single string.
func joinString(args ...interface{}) string {
	str := ""
	for index := 0; index < len(args); index++ {
		msg := fmt.Sprintf("%v", args[index])
		str += " " + msg
	}
	return str
}

// SetLogLevel sets the log level for the logger instance.
func (l *Logger) SetLogLevel(level logrus.Level) {
	l.LogrusInstance.SetLevel(level)
}

// updateBlockInfo fetches block info from the BlockMonitor.
func (l *Logger) updateBlockInfo() (uint32, *big.Int) {
	if l.blockMonitor == nil {
		return 0, nil
	}
	latestBlock := l.blockMonitor.GetLatestBlock()
	if latestBlock != nil {
		epoch := uint32(latestBlock.Time / core.EpochLength)
		return epoch, latestBlock.Number
	}
	return 0, nil
}
