package common

import "github.com/sirupsen/logrus"

var (
	// JSONLog logs with json formatter
	JSONLog *logrus.Logger

	// TextLog logs with text formatter
	TextLog *logrus.Logger
)

// InitLog .
func InitLog() {
	JSONLog = logrus.New()
	JSONLog.SetFormatter(&logrus.JSONFormatter{})

	TextLog = logrus.New()
	TextLog.SetFormatter(&logrus.TextFormatter{})
}
