package logging

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/Tackem-org/Global/logging/debug"
)

var (
	mu             sync.Mutex
	i              *log.Logger
	d              *log.Logger
	e              *log.Logger
	w              *log.Logger
	f              *log.Logger
	t              *log.Logger
	file           *os.File
	filePath       string
	maxSize        int64 = 50 * 1024
	fileCountLimit uint8 = 5
	logVerbose     bool
	mw             io.Writer
	dm             debug.Mask

	//Only one of the following
	//Debug Log Settings
	// logSettings = log.Ldate|log.Ltime|log.Lshortfile
	//Production Settings
	logSettings = log.Ldate | log.Ltime
)

func Setup(logFile string, verbose bool, debugMask debug.Mask) {
	mu.Lock()
	defer mu.Unlock()
	filePath = logFile
	logVerbose = verbose
	dm = debugMask
	setupBackend()
}

func Shutdown() {
	mu.Lock()
	defer mu.Unlock()
	file.Close()
}

func setupBackend() {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}
	if logVerbose {
		mw = io.MultiWriter(os.Stdout, file)
	} else {
		mw = file
	}
	i = log.New(mw, "INFO: ", logSettings)
	d = log.New(mw, "DEBUG: ", logSettings)
	e = log.New(mw, "ERROR: ", logSettings)
	w = log.New(mw, "WARNING: ", logSettings)
	f = log.New(mw, "FATAL: ", logSettings)
	t = log.New(mw, "TODO: ", logSettings)
}

func checkLogSize() {
	fhandler, _ := os.Stat(filePath)
	size := fhandler.Size()
	if maxSize < size {
		file.Close()
		moveBackupLogFiles(0)
		os.Rename(filePath, filePath+".0.bak")
		setupBackend()
	}
}

func moveBackupLogFiles(i uint8) {
	if !fileExists(fmt.Sprintf("%s.%d.bak", filePath, i)) {
		return
	}
	if fileCountLimit > 0 && i >= fileCountLimit {
		os.Remove(fmt.Sprintf("%s.%d.bak", filePath, i))
		return
	}

	moveBackupLogFiles(i + 1)
	os.Rename(fmt.Sprintf("%s.%d.bak", filePath, i), fmt.Sprintf("%s.%d.bak", filePath, i+1))
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func CustomLogger(prefix string) *log.Logger {
	mu.Lock()
	defer mu.Unlock()
	return log.New(mw, prefix+": ", logSettings)
}

func Custom(prefix string, message string) {
	mu.Lock()
	defer mu.Unlock()
	custom(prefix, message)
}

func Customf(prefix string, message string, values ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	custom(prefix, fmt.Sprintf(message, values...))
}

func custom(prefix string, message string) {
	t := log.New(mw, prefix+": ", logSettings)
	t.Println(message)
}

func Info(message string) {
	mu.Lock()
	defer mu.Unlock()
	info(message)
}

func Infof(message string, values ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	info(fmt.Sprintf(message, values...))
}

func info(message string) {
	i.Println(message)
	checkLogSize()
}

func Debug(debugMask debug.Mask, message string) {
	mu.Lock()
	defer mu.Unlock()
	debugBack(debugMask, message)
}

func Debugf(debugMask debug.Mask, message string, values ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	debugBack(debugMask, fmt.Sprintf(message, values...))
}

func debugBack(debugMask debug.Mask, message string) {
	if dm == 0 {
		return
	}
	if !dm.HasAny(debugMask) {
		return
	}
	d.Println(message)
	checkLogSize()
}

func Warning(message string) {
	mu.Lock()
	defer mu.Unlock()
	warning(message)
}

func Warningf(message string, values ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	warning(fmt.Sprintf(message, values...))
}

func warning(message string) {
	w.Println(message)
	checkLogSize()
}

func Error(message string) {
	mu.Lock()
	defer mu.Unlock()
	error(message)
}

func Errorf(message string, values ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	error(fmt.Sprintf(message, values...))
}

func error(message string) {
	e.Println(message)
	checkLogSize()
}

func Todo(message string) {
	mu.Lock()
	defer mu.Unlock()
	todo(message)
}

func Todof(message string, values ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	todo(fmt.Sprintf(message, values...))
}

func todo(message string) {
	t.Println(message)
	checkLogSize()
}

func Fatal(message string) {
	mu.Lock()
	defer mu.Unlock()
	fatal(message)
}

func Fatalf(message string, values ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	fatal(fmt.Sprintf(message, values...))
}

func fatal(message string) {
	f.Println(message)
	panic(errors.New(message))
}
