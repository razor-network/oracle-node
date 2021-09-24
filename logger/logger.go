package logger

import (
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

func init() {

	logFilePath, err := path.GetLogFilePath()
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
