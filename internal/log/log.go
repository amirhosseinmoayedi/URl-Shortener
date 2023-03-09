package log

import (
	"github.com/sirupsen/logrus"
	"os"
)

var Logger *logrus.Logger

func InitLogger() {
	level, err := logrus.ParseLevel("debug")
	if err != nil {
		logrus.Fatalf("Error in initiating logger")
	}

	Logger = &logrus.Logger{
		Out:   os.Stderr,
		Level: level,
		Hooks: make(logrus.LevelHooks),
		// ReportCaller: true,
		Formatter: &logrus.TextFormatter{
			ForceColors:            true,
			DisableColors:          false,
			DisableLevelTruncation: true,
			TimestampFormat:        "2006-01-02 15:04:05",
			// LogFormat:       "[%time%] %lvl%, %msg%",
		},
	}
}
