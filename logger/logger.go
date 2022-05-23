//Packages loggers provides function for logging messages for a specific system or application component
package logger

import (
	"errors"
	"fmt"
	"github.com/razor-network/goInfo"
	"github.com/sirupsen/logrus"
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

//This function initializes the logger
func InitializeLogger(fileName string) {
	if fileName != "" {
		logFilePath, err := path.PathUtilsInterface.GetLogFilePath(fileName)
		if err != nil {
			standardLogger.Fatal("Error in fetching log file path: ", err)
		}

		lumberJackLogger := &lumberjack.Logger{
			Filename:   logFilePath,
			MaxSize:    5,
			MaxBackups: 10,
			MaxAge:     30,
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

//This function joins the string
func joinString(args ...interface{}) string {
	str := ""
	for index := 0; index < len(args); index++ {
		msg := fmt.Sprintf("%v", args[index])
		str += " " + msg
	}
	return str
}

//This function handles the errors in logs
func (logger *StandardLogger) Error(args ...interface{}) {
	var addressLogField = logrus.Fields{
		"address": Address,
	}
	logger.WithFields(addressLogField).Errorln(args...)
}

//This function handles the info in logs
func (logger *StandardLogger) Info(args ...interface{}) {
	var addressLogField = logrus.Fields{
		"address": Address,
	}
	logger.WithFields(addressLogField).Infoln(args...)
}

//This function helps in debugging from logs
func (logger *StandardLogger) Debug(args ...interface{}) {
	var addressLogField = logrus.Fields{
		"address": Address,
	}
	logger.WithFields(addressLogField).Debugln(args...)
}

//This function handles the fatal messages from logs
func (logger *StandardLogger) Fatal(args ...interface{}) {
	var addressLogField = logrus.Fields{
		"address": Address,
	}
	errMsg := joinString(args)
	err := errors.New(errMsg)
	logger.WithFields(addressLogField).Fatalln(err)
}

//This function allows us to use formatting features to create descriptive error messages
func (logger *StandardLogger) Errorf(format string, args ...interface{}) {
	var addressLogField = logrus.Fields{
		"address": Address,
	}
	logger.WithFields(addressLogField).Errorf(format, args...)
}

//This function allows us to use formatting features to create descriptive info messages
func (logger *StandardLogger) Infof(format string, args ...interface{}) {
	var addressLogField = logrus.Fields{
		"address": Address,
	}
	logger.WithFields(addressLogField).Infof(format, args...)
}

//This function allows us to use formatting features to create descriptive debug messages
func (logger *StandardLogger) Debugf(format string, args ...interface{}) {
	var addressLogField = logrus.Fields{
		"address": Address,
	}
	logger.WithFields(addressLogField).Debugf(format, args...)
}

//This function allows us to use formatting features to create descriptive fatal messages
func (logger *StandardLogger) Fatalf(format string, args ...interface{}) {
	var addressLogField = logrus.Fields{
		"address": Address,
	}
	errMsg := joinString(args)
	err := errors.New(errMsg)
	logger.WithFields(addressLogField).Fatalf(format, err)
}
