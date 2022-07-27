package logger

import (
	"errors"
	"fmt"
	"github.com/razor-network/goInfo"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"os"
	"razor/core"
	"razor/path"
	"runtime"
)

type StandardLogger struct {
	*logrus.Logger
}

var standardLogger = &StandardLogger{logrus.New()}

var Address string
var FileName string

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

		logFileMaxSize := viper.GetInt("logFileMaxSize")
		if logFileMaxSize == 0 {
			logFileMaxSize = 5
		}
		logFileMaxBackups := viper.GetInt("logFileMaxBackups")
		if logFileMaxBackups == 0 {
			logFileMaxBackups = 10
		}
		logFileMaxAge := viper.GetInt("logFileMaxAge")
		if logFileMaxAge == 0 {
			logFileMaxAge = 30
		}

		lumberJackLogger := &lumberjack.Logger{
			Filename:   logFilePath,
			MaxSize:    logFileMaxSize,
			MaxBackups: logFileMaxBackups,
			MaxAge:     logFileMaxAge,
		}

		stderr := os.Stderr
		stdin := os.Stdin
		stdout := os.Stdout
		mw := io.MultiWriter(stderr, stdin, stdout, lumberJackLogger)
		standardLogger.SetOutput(mw)
	}
	standardLogger.Formatter = &logrus.JSONFormatter{}
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
	var addressLogField = logrus.Fields{
		"address": Address,
	}
	logger.WithFields(addressLogField).Errorln(args...)
}

func (logger *StandardLogger) Info(args ...interface{}) {
	var addressLogField = logrus.Fields{
		"address": Address,
	}
	logger.WithFields(addressLogField).Infoln(args...)
}

func (logger *StandardLogger) Debug(args ...interface{}) {
	var addressLogField = logrus.Fields{
		"address": Address,
	}
	logger.WithFields(addressLogField).Debugln(args...)
}

func (logger *StandardLogger) Fatal(args ...interface{}) {
	var addressLogField = logrus.Fields{
		"address": Address,
	}
	errMsg := joinString(args)
	err := errors.New(errMsg)
	logger.WithFields(addressLogField).Fatalln(err)
}

func (logger *StandardLogger) Errorf(format string, args ...interface{}) {
	var addressLogField = logrus.Fields{
		"address": Address,
	}
	logger.WithFields(addressLogField).Errorf(format, args...)
}

func (logger *StandardLogger) Infof(format string, args ...interface{}) {
	var addressLogField = logrus.Fields{
		"address": Address,
	}
	logger.WithFields(addressLogField).Infof(format, args...)
}

func (logger *StandardLogger) Debugf(format string, args ...interface{}) {
	var addressLogField = logrus.Fields{
		"address": Address,
	}
	logger.WithFields(addressLogField).Debugf(format, args...)
}

func (logger *StandardLogger) Fatalf(format string, args ...interface{}) {
	var addressLogField = logrus.Fields{
		"address": Address,
	}
	errMsg := joinString(args)
	err := errors.New(errMsg)
	logger.WithFields(addressLogField).Fatalf(format, err)
}
