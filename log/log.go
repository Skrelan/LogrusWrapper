package log

import (
	"github.com/sirupsen/logrus"
)

var muteable = logrus.New()
var immuteable = logrus.New()

func init() {
	s := "debug" //you can use the Viper pkg to read from an config file and assign it here.
	muteable.SetLevel(setLogLevel(s))
}

func setLogLevel(s string) logrus.Level {
	switch s {
	case "debug":
		return logrus.DebugLevel
	case "info":
		return logrus.InfoLevel
	case "warn":
		return logrus.WarnLevel
	case "error":
		return logrus.ErrorLevel
	default:
		return logrus.DebugLevel
	}
}

//Debug is a wrapper for Debug function
func Debug(s ...interface{}) {
	muteable.Debug(s)
}

//Debugf is a wrapper for Debugf function
func Debugf(format string, args ...interface{}) {
	muteable.Debugf(format, args)
}

//Info is a wrapper for Info function
func Info(s ...interface{}) {
	muteable.Info(s)
}

//Infof is a wrapper for Infof function
func Infof(format string, args ...interface{}) {
	muteable.Infof(format, args)
}

//Warn is a wrapper for Warn function
func Warn(s ...interface{}) {
	muteable.Warn(s)
}

//Warnf is a wrapper for Warnf function
func Warnf(format string, args ...interface{}) {
	muteable.Warnf(format, args)
}

//Error is a wrapper for Error function
func Error(s ...interface{}) {
	muteable.Error(s)
}

//Errorf is a wrapper for Errorf function
func Errorf(format string, args ...interface{}) {
	muteable.Errorf(format, args)
}

//Panic is a wrapper for Panic function
func Panic(s ...interface{}) {
	muteable.Panic(s)
}

//Panicf is a wrapper for Panicf function
func Panicf(format string, args ...interface{}) {
	muteable.Panicf(format, args)
}

//AlwaysLog is will always log data passed to this
func AlwaysLog(s ...interface{}) {
	immuteable.Info(s)
}

//AlwaysLogf will always log the data passed to this
func AlwaysLogf(format string, args ...interface{}) {
	immuteable.Infof(format, args)
}
