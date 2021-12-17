package logging

import (
	"io"
	"log"
	"os"
)

var (
	i    *log.Logger
	e    *log.Logger
	w    *log.Logger
	f    *log.Logger
	file *os.File
	mw   io.Writer

	//Only one of the following
	//Debug Log Settings
	// logSettings = log.Ldate|log.Ltime|log.Lshortfile
	//Production Settings
	logSettings = log.Ldate | log.Ltime
)

func Setup(logFile string, verbose bool) {
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	if verbose {
		mw = io.MultiWriter(os.Stdout, file)
	} else {
		mw = file
	}
	i = log.New(mw, "INFO: ", logSettings)
	e = log.New(mw, "ERROR: ", logSettings)
	w = log.New(mw, "WARNING: ", logSettings)
	f = log.New(mw, "FATAL: ", logSettings)
}

func Shutdown() {
	file.Close()

}

func CustomLogger(prefix string) *log.Logger {
	return log.New(mw, prefix+": ", logSettings)
}

func Custom(prefix string, message string) {
	t := log.New(mw, prefix+": ", logSettings)
	t.Println(message)
}

func Info(message string) {
	i.Println(message)
}

func Warning(message string) {
	w.Println(message)
}

func Error(message string) {
	e.Println(message)
}

func Fatal(err error) {
	f.Println(err.Error())
	panic(err)
}
