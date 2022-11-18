package logger

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/razor-network/goInfo"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"math/big"
	"os"
	"razor/block"
	"razor/core"
	"razor/path"
	"runtime"
)

type StandardLogger struct {
	*logrus.Logger
}

var standardLogger = &StandardLogger{logrus.New()}

var Address string
var Epoch uint32
var BlockNumber *big.Int
var FileName string
var Client *ethclient.Client

func init() {
	path.PathUtilsInterface = &path.PathUtils{}
	path.OSUtilsInterface = &path.OSUtils{}
	InitializeLogger(FileName)

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

func InitializeLogger(fileName string) {
	if fileName != "" {
		logFilePath, err := path.PathUtilsInterface.GetLogFilePath(fileName)
		if err != nil {
			standardLogger.Fatal("Error in fetching log file path: ", err)
		}

		lumberJackLogger := &lumberjack.Logger{
			Filename:   logFilePath,
			MaxSize:    182, // Maximum Size of a log file
			MaxBackups: 52,  // Maximum number of log files
			MaxAge:     365, // Maximum number of days to retain olf files
		}

		out := os.Stderr
		mw := io.MultiWriter(out, lumberJackLogger)
		standardLogger.Formatter = &logrus.JSONFormatter{}
		standardLogger.SetOutput(mw)

	} else {
		standardLogger.Formatter = &logrus.JSONFormatter{}
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

func (logger *StandardLogger) Error(args ...interface{}) {
	SetEpochAndBlockNumber(Client)
	var logFields = logrus.Fields{
		"address":     Address,
		"epoch":       Epoch,
		"blockNumber": BlockNumber,
	}
	logger.WithFields(logFields).Errorln(args...)
}

func (logger *StandardLogger) Info(args ...interface{}) {
	SetEpochAndBlockNumber(Client)
	var logFields = logrus.Fields{
		"address":     Address,
		"epoch":       Epoch,
		"blockNumber": BlockNumber,
	}
	logger.WithFields(logFields).Infoln(args...)
}

func (logger *StandardLogger) Debug(args ...interface{}) {
	SetEpochAndBlockNumber(Client)
	var logFields = logrus.Fields{
		"address":     Address,
		"epoch":       Epoch,
		"blockNumber": BlockNumber,
	}
	logger.WithFields(logFields).Debugln(args...)
}

func (logger *StandardLogger) Fatal(args ...interface{}) {
	SetEpochAndBlockNumber(Client)
	var logFields = logrus.Fields{
		"address":     Address,
		"epoch":       Epoch,
		"blockNumber": BlockNumber,
	}
	errMsg := joinString(args)
	err := errors.New(errMsg)
	logger.WithFields(logFields).Fatalln(err)
}

func (logger *StandardLogger) Errorf(format string, args ...interface{}) {
	SetEpochAndBlockNumber(Client)
	var logFields = logrus.Fields{
		"address":     Address,
		"epoch":       Epoch,
		"blockNumber": BlockNumber,
	}
	logger.WithFields(logFields).Errorf(format, args...)
}

func (logger *StandardLogger) Infof(format string, args ...interface{}) {
	SetEpochAndBlockNumber(Client)
	var logFields = logrus.Fields{
		"address":     Address,
		"epoch":       Epoch,
		"blockNumber": BlockNumber,
	}
	logger.WithFields(logFields).Infof(format, args...)
}

func (logger *StandardLogger) Debugf(format string, args ...interface{}) {
	SetEpochAndBlockNumber(Client)
	var logFields = logrus.Fields{
		"address":     Address,
		"epoch":       Epoch,
		"blockNumber": BlockNumber,
	}
	logger.WithFields(logFields).Debugf(format, args...)
}

func (logger *StandardLogger) Fatalf(format string, args ...interface{}) {
	SetEpochAndBlockNumber(Client)
	var logFields = logrus.Fields{
		"address":     Address,
		"epoch":       Epoch,
		"blockNumber": BlockNumber,
	}
	errMsg := joinString(args)
	err := errors.New(errMsg)
	logger.WithFields(logFields).Fatalf(format, err)
}

func SetEpochAndBlockNumber(client *ethclient.Client) {
	if client != nil {
		latestBlock := block.GetLatestBlock()
		if latestBlock != nil {
			BlockNumber = latestBlock.Number
			epoch := latestBlock.Time / uint64(core.EpochLength)
			Epoch = uint32(epoch)
		}
	}
}

func SetLoggerParameters(client *ethclient.Client, address string) {
	Address = address
	Client = client
	go block.CalculateLatestBlock(client)
}
