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

type StandardLogger struct {
	*logrus.Logger
	address      string
	epoch        uint32
	blockNumber  *big.Int
	client       *ethclient.Client
	blockManager *block.BlockManager
}

var standardLogger *StandardLogger

func init() {
	path.PathUtilsInterface = &path.PathUtils{}
	path.OSUtilsInterface = &path.OSUtils{}
	standardLogger = initLogger()

	osInfo := goInfo.GetInfo()
	standardLogger.WithFields(logrus.Fields{
		"Operating System": osInfo.OS,
		"Core":             osInfo.Core,
		"Platform":         osInfo.Platform,
		"CPUs":             osInfo.CPUs,
		"razor-go version": core.VersionWithMeta,
		"go version":       runtime.Version(),
	}).Info()

}

func initLogger() *StandardLogger {
	logger := &StandardLogger{
		Logger: logrus.New(),
	}
	logger.Formatter = &logrus.JSONFormatter{}
	return logger
}

func SetupLogFile(fileName string, config types.Configurations) {
	if fileName != "" {
		logFilePath, err := path.PathUtilsInterface.GetLogFilePath(fileName)
		if err != nil {
			standardLogger.Fatal("Error in fetching log file path: ", err)
		}

		lumberJackLogger := &lumberjack.Logger{
			Filename:   logFilePath,
			MaxSize:    config.LogFileMaxSize,
			MaxBackups: config.LogFileMaxBackups,
			MaxAge:     config.LogFileMaxAge,
		}

		out := os.Stderr
		mw := io.MultiWriter(out, lumberJackLogger)
		standardLogger.SetOutput(mw)
	}
}

func NewLogger() *StandardLogger {
	return standardLogger
}

func joinString(args ...interface{}) string {
	str := ""
	for index := 0; index < len(args); index++ {
		msg := fmt.Sprintf("%v", args[index])
		str += " " + msg
	}
	return str
}

// Helper function to prepare log fields
func (logger *StandardLogger) prepareLogFields() logrus.Fields {
	logger.SetEpochAndBlockNumber(logger.client)
	return logrus.Fields{
		"address":     logger.address,
		"epoch":       logger.epoch,
		"blockNumber": logger.blockNumber,
		"version":     core.VersionWithMeta,
	}
}

func (logger *StandardLogger) Error(args ...interface{}) {
	logger.WithFields(logger.prepareLogFields()).Errorln(args...)
}

func (logger *StandardLogger) Info(args ...interface{}) {
	logger.WithFields(logger.prepareLogFields()).Infoln(args...)
}

func (logger *StandardLogger) Debug(args ...interface{}) {
	logger.WithFields(logger.prepareLogFields()).Debugln(args...)
}

func (logger *StandardLogger) Warn(args ...interface{}) {
	logger.WithFields(logger.prepareLogFields()).Warnln(args...)
}

func (logger *StandardLogger) Fatal(args ...interface{}) {
	errMsg := joinString(args)
	err := errors.New(errMsg)
	logger.WithFields(logger.prepareLogFields()).Fatalln(err)
}

func (logger *StandardLogger) Errorf(format string, args ...interface{}) {
	logger.WithFields(logger.prepareLogFields()).Errorf(format, args...)
}

func (logger *StandardLogger) Infof(format string, args ...interface{}) {
	logger.WithFields(logger.prepareLogFields()).Infof(format, args...)
}

func (logger *StandardLogger) Debugf(format string, args ...interface{}) {
	logger.WithFields(logger.prepareLogFields()).Debugf(format, args...)
}

func (logger *StandardLogger) Warnf(format string, args ...interface{}) {
	logger.WithFields(logger.prepareLogFields()).Warnf(format, args...)
}

func (logger *StandardLogger) Fatalf(format string, args ...interface{}) {
	errMsg := joinString(args)
	err := errors.New(errMsg)
	logger.WithFields(logger.prepareLogFields()).Fatalf(format, err)
}

func (logger *StandardLogger) SetEpochAndBlockNumber(client *ethclient.Client) {
	if client != nil {
		latestBlock := logger.blockManager.GetLatestBlock()
		if latestBlock != nil {
			logger.blockNumber = latestBlock.Number
			epoch := latestBlock.Time / core.EpochLength
			logger.epoch = uint32(epoch)
		}
	}
}

func (logger *StandardLogger) SetLoggerParameters(client *ethclient.Client, address string) {
	logger.address = address
	logger.client = client
	blockManager := block.NewBlockManager(client)
	logger.blockManager = blockManager
	go blockManager.CalculateLatestBlock()
}
