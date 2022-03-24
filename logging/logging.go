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
	I LoggingInterface = DefaultLogging()
)

func Setup(logFile string, verbose bool) {
	I.Setup(logFile, verbose)
}

func Shutdown() {
	I.Shutdown()
}

func CustomLogger(prefix string) *log.Logger {
	return I.CustomLogger(prefix)
}

func Custom(prefix string, message string, values ...interface{}) {
	I.Custom(prefix, message, values...)
}

func Info(message string, values ...interface{}) {
	I.Info(message, values...)
}

func Warning(message string, values ...interface{}) {
	I.Warning(message, values...)
}

func Error(message string, values ...interface{}) {
	I.Error(message, values...)
}

func Todo(message string, values ...interface{}) {
	I.Todo(message, values...)
}

func Fatal(message string, values ...interface{}) {
	panic(I.Fatal(message, values...))
}
