package logger

import (
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
	logrusInstance *logrus.Logger
	address        string
	client         *ethclient.Client
	blockMonitor   *block.BlockMonitor
}

// NewLogger initializes a Logger instance with given parameters.
func NewLogger(address string, client *ethclient.Client, fileName string, config types.Configurations, blockMonitor *block.BlockMonitor) *Logger {
	logger := &Logger{
		logrusInstance: logrus.New(),
		address:        address,
		client:         client,
		blockMonitor:   blockMonitor,
	}

	logger.setupLogger(fileName, config)
	logger.logSystemInfo()

	return logger
}

// setupLogger configures the logger's output and formatting.
func (l *Logger) setupLogger(fileName string, config types.Configurations) {
	path.PathUtilsInterface = &path.PathUtils{}
	path.OSUtilsInterface = &path.OSUtils{}
	if fileName != "" {
		logFilePath, err := path.PathUtilsInterface.GetLogFilePath(fileName)
		if err != nil {
			l.logrusInstance.Fatal("Error in fetching log file path: ", err)
		}

		lumberJackLogger := &lumberjack.Logger{
			Filename:   logFilePath,
			MaxSize:    config.LogFileMaxSize,
			MaxBackups: config.LogFileMaxBackups,
			MaxAge:     config.LogFileMaxAge,
		}

		mw := io.MultiWriter(os.Stderr, lumberJackLogger)
		l.logrusInstance.SetOutput(mw)
	}
	l.logrusInstance.Formatter = &logrus.JSONFormatter{}
}

// logSystemInfo logs system-related details at initialization.
func (l *Logger) logSystemInfo() {
	osInfo := goInfo.GetInfo()
	l.logrusInstance.WithFields(logrus.Fields{
		"Operating System": osInfo.OS,
		"Core":             osInfo.Core,
		"Platform":         osInfo.Platform,
		"CPUs":             osInfo.CPUs,
		"razor-go version": core.VersionWithMeta,
		"go version":       runtime.Version(),
	}).Info("Logger initialized")
}

// logWithFields adds context and logs a message with the specified level and format.
func (l *Logger) logWithFields(level logrus.Level, format string, args ...interface{}) {
	epoch, blockNumber := l.updateBlockInfo()
	fields := logrus.Fields{
		"address":     l.address,
		"epoch":       epoch,
		"blockNumber": blockNumber,
		"version":     core.VersionWithMeta,
	}
	entry := l.logrusInstance.WithFields(fields)

	message := fmt.Sprintf(format, args...)
	switch level {
	case logrus.InfoLevel:
		entry.Info(message)
	case logrus.DebugLevel:
		entry.Debug(message)
	case logrus.WarnLevel:
		entry.Warn(message)
	case logrus.ErrorLevel:
		entry.Error(message)
	case logrus.FatalLevel:
		entry.Fatal(message)
		os.Exit(1)
	default:
		entry.Print(message)
	}
}

// Info logs a simple informational message.
func (l *Logger) Info(args ...interface{}) {
	l.logWithFields(logrus.InfoLevel, joinString(args...), args...)
}

// Infof logs a formatted informational message.
func (l *Logger) Infof(format string, args ...interface{}) {
	l.logWithFields(logrus.InfoLevel, format, args...)
}

// Debug logs a simple debug message.
func (l *Logger) Debug(args ...interface{}) {
	l.logWithFields(logrus.DebugLevel, joinString(args...), args...)
}

// Debugf logs a formatted debug message.
func (l *Logger) Debugf(format string, args ...interface{}) {
	l.logWithFields(logrus.DebugLevel, format, args...)
}

// Warn logs a simple warning message.
func (l *Logger) Warn(args ...interface{}) {
	l.logWithFields(logrus.WarnLevel, joinString(args...), args...)
}

// Warnf logs a formatted warning message.
func (l *Logger) Warnf(format string, args ...interface{}) {
	l.logWithFields(logrus.WarnLevel, format, args...)
}

// Error logs a simple error message.
func (l *Logger) Error(args ...interface{}) {
	l.logWithFields(logrus.ErrorLevel, joinString(args...), args...)
}

// Errorf logs a formatted error message.
func (l *Logger) Errorf(format string, args ...interface{}) {
	l.logWithFields(logrus.ErrorLevel, format, args...)
}

// Fatal logs a simple fatal error message and exits the application.
func (l *Logger) Fatal(args ...interface{}) {
	l.logWithFields(logrus.FatalLevel, joinString(args...), args...)
}

// Fatalf logs a formatted fatal error message and exits the application.
func (l *Logger) Fatalf(format string, args ...interface{}) {
	l.logWithFields(logrus.FatalLevel, format, args...)
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

// SetAddress updates the logger's address field.
func (l *Logger) SetAddress(address string) {
	l.address = address
}

// SetClient updates the logger's Ethereum client.
func (l *Logger) SetClient(client *ethclient.Client) {
	l.client = client
}

// SetLogLevel sets the log level for the logger instance.
func (l *Logger) SetLogLevel(level logrus.Level) {
	l.logrusInstance.SetLevel(level)
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
