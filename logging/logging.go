package logging

import "log"

type LoggingInterface interface {
	Setup(logFile string, verbose bool)
	Shutdown()
	CustomLogger(prefix string) *log.Logger
	Custom(prefix string, message string, values ...interface{})
	Info(message string, values ...interface{})
	Warning(message string, values ...interface{})
	Error(message string, values ...interface{})
	Todo(message string, values ...interface{})
	Fatal(message string, values ...interface{}) error
}

var (
	LI LoggingInterface
)

func Setup(logFile string, verbose bool) {
	if LI == nil {
		LI = DefaultLogging()
	}
	LI.Setup(logFile, verbose)
}

func Shutdown() {
	if LI == nil {
		return
	}
	LI.Shutdown()
}

func CustomLogger(prefix string) *log.Logger {
	if LI == nil {
		return nil
	}
	return LI.CustomLogger(prefix)
}

func Custom(prefix string, message string, values ...interface{}) {
	if LI == nil {
		return
	}
	LI.Custom(prefix, message, values...)
}

func Info(message string, values ...interface{}) {
	if LI == nil {
		return
	}
	LI.Info(message, values...)
}

func Warning(message string, values ...interface{}) {
	if LI == nil {
		return
	}
	LI.Warning(message, values...)
}

func Error(message string, values ...interface{}) {
	if LI == nil {
		return
	}
	LI.Error(message, values...)
}

func Todo(message string, values ...interface{}) {
	if LI == nil {
		return
	}
	LI.Todo(message, values...)
}

func Fatal(message string, values ...interface{}) {
	if LI == nil {
		return
	}
	panic(LI.Fatal(message, values...))

}
