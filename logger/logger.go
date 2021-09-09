package logger

import (
	"github.com/razor-network/goInfo"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
	"razor/path"
)

type StandardLogger struct {
	*logrus.Logger
}

var standardLogger = &StandardLogger{logrus.New()}

func init() {

	logFilePath, err := path.GetLogFilePath()
	if err != nil {
		standardLogger.Fatal("Error in fetching log file path: ", err)
	}

	standardLogger.Formatter = &logrus.JSONFormatter{}
	standardLogger.Out = &lumberjack.Logger{
		Filename:   logFilePath,
		MaxSize:    5,
		MaxBackups: 10,
		MaxAge:     30,
	}

	osInfo := goInfo.GetInfo()

	standardLogger.WithFields(logrus.Fields{
		"Operating System": osInfo.OS,
		"Core":             osInfo.Core,
		"Platform":         osInfo.Platform,
		"CPUs":             osInfo.CPUs,
	}).Info()

}

func NewLogger() *StandardLogger {
	return standardLogger
}
