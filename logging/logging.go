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
	I LoggingInterface
)

func Setup(logFile string, verbose bool) {
	if I == nil {
		I = DefaultLogging()
	}
	I.Setup(logFile, verbose)
}

func Shutdown() {
	if I == nil {
		return
	}
	I.Shutdown()
}

func CustomLogger(prefix string) *log.Logger {
	if I == nil {
		return nil
	}
	return I.CustomLogger(prefix)
}

func Custom(prefix string, message string, values ...interface{}) {
	if I == nil {
		return
	}
	I.Custom(prefix, message, values...)
}

func Info(message string, values ...interface{}) {
	if I == nil {
		return
	}
	I.Info(message, values...)
}

func Warning(message string, values ...interface{}) {
	if I == nil {
		return
	}
	I.Warning(message, values...)
}

func Error(message string, values ...interface{}) {
	if I == nil {
		return
	}
	I.Error(message, values...)
}

func Todo(message string, values ...interface{}) {
	if I == nil {
		return
	}
	I.Todo(message, values...)
}

func Fatal(message string, values ...interface{}) {
	if I == nil {
		return
	}
	panic(I.Fatal(message, values...))

}
