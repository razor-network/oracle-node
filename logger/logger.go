package logger

import (
	"errors"
	"fmt"
	"github.com/getsentry/sentry-go"
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

func init() {
	path.PathUtilsInterface = &path.PathUtils{}
	path.OSUtilsInterface = &path.OSUtils{}
	logFilePath, err := path.PathUtilsInterface.GetLogFilePath()
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
		"Address": Address,
	}
	errMsg := joinString(args)
	err := errors.New(errMsg)
	sentry.CaptureException(err)
	logger.WithFields(addressLogField).Errorln(args...)
}

func (logger *StandardLogger) Info(args ...interface{}) {
	var addressLogField = logrus.Fields{
		"Address": Address,
	}
	msg := joinString(args)
	sentry.CaptureMessage(msg)
	logger.WithFields(addressLogField).Infoln(args...)
}

func (logger *StandardLogger) Debug(args ...interface{}) {
	var addressLogField = logrus.Fields{
		"Address": Address,
	}
	msg := joinString(args)
	sentry.CaptureMessage(msg)
	logger.WithFields(addressLogField).Debugln(args...)
}

func (logger *StandardLogger) Fatal(args ...interface{}) {
	var addressLogField = logrus.Fields{
		"Address": Address,
	}
	defer sentry.Recover()
	errMsg := joinString(args)
	err := errors.New(errMsg)
	sentry.WithScope(func(scope *sentry.Scope) {
		scope.SetLevel(sentry.LevelFatal)
		sentry.CaptureException(err)
	})
	logger.WithFields(addressLogField).Fatalln(err)
}
