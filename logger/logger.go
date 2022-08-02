//Package loggers provides function for logging messages for a specific system or application component
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
	"strconv"
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

// InitializeLogger function initializes the logger
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
		compress := viper.GetString("compress")
		if compress == "" {
			compress = "true"
		}
		compressBool, err := strconv.ParseBool(compress)
		if err != nil {
			return
		}

		lumberJackLogger := &lumberjack.Logger{
			Filename:   logFilePath,
			MaxSize:    logFileMaxSize,
			MaxBackups: logFileMaxBackups,
			MaxAge:     logFileMaxAge,
			Compress:   compressBool,
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

//This function joins the string
func joinString(args ...interface{}) string {
	str := ""
	for index := 0; index < len(args); index++ {
		msg := fmt.Sprintf("%v", args[index])
		str += " " + msg
	}
	return str
}

// Error function handles the errors in logs
func (logger *StandardLogger) Error(args ...interface{}) {
	var addressLogField = logrus.Fields{
		"address": Address,
	}
	logger.WithFields(addressLogField).Errorln(args...)
}

// Info function handles the info in logs
func (logger *StandardLogger) Info(args ...interface{}) {
	var addressLogField = logrus.Fields{
		"address": Address,
	}
	logger.WithFields(addressLogField).Infoln(args...)
}

// Debug function helps in debugging from logs
func (logger *StandardLogger) Debug(args ...interface{}) {
	var addressLogField = logrus.Fields{
		"address": Address,
	}
	logger.WithFields(addressLogField).Debugln(args...)
}

// Fatal function handles the fatal messages from logs
func (logger *StandardLogger) Fatal(args ...interface{}) {
	var addressLogField = logrus.Fields{
		"address": Address,
	}
	errMsg := joinString(args)
	err := errors.New(errMsg)
	logger.WithFields(addressLogField).Fatalln(err)
}

// Errorf function allows us to use formatting features to create descriptive error messages
func (logger *StandardLogger) Errorf(format string, args ...interface{}) {
	var addressLogField = logrus.Fields{
		"address": Address,
	}
	logger.WithFields(addressLogField).Errorf(format, args...)
}

// Infof function allows us to use formatting features to create descriptive info messages
func (logger *StandardLogger) Infof(format string, args ...interface{}) {
	var addressLogField = logrus.Fields{
		"address": Address,
	}
	logger.WithFields(addressLogField).Infof(format, args...)
}

// Debugf function allows us to use formatting features to create descriptive debug messages
func (logger *StandardLogger) Debugf(format string, args ...interface{}) {
	var addressLogField = logrus.Fields{
		"address": Address,
	}
	logger.WithFields(addressLogField).Debugf(format, args...)
}

// Fatalf function allows us to use formatting features to create descriptive fatal messages
func (logger *StandardLogger) Fatalf(format string, args ...interface{}) {
	var addressLogField = logrus.Fields{
		"address": Address,
	}
	errMsg := joinString(args)
	err := errors.New(errMsg)
	logger.WithFields(addressLogField).Fatalf(format, err)
}
