package services

import (
	"github.com/sirupsen/logrus"
)

func ContextFields() logrus.Fields {
	return logrus.Fields{}
}

func ConfigureLogging() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
}

func InfoFormat(format string, args ...interface{}) {
	logrus.WithFields(ContextFields()).Infof(format, args...)
}

func DebugFormat(format string, args ...interface{}) {
	logrus.WithFields(ContextFields()).Debugf(format, args...)
}

func ErrorLine(args ...interface{}) {
	logrus.WithFields(ContextFields()).Errorln(args...)
}

func ErrorFormat(format string, args ...interface{}) {
	logrus.WithFields(ContextFields()).Errorf(format, args...)
}
