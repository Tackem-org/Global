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

//https://stackoverflow.com/questions/44119951/how-to-check-a-log-output-in-go-test
var (
	mu              sync.Mutex
	i               *log.Logger
	d               *log.Logger
	e               *log.Logger
	w               *log.Logger
	f               *log.Logger
	t               *log.Logger
	file            *os.File
	filePath        string
	restrictLogSize bool  = true
	dontPanic       bool  = false
	maxSize         int64 = 50 * 1024
	fileCountLimit  uint8 = 5
	logVerbose      bool
	mw              io.Writer
	dm              debug.Mask
	logSettings     int = log.Ldate | log.Ltime | log.Lmicroseconds
)

func Setup(logFile string, verbose bool, debugMask debug.Mask) {
	mu.Lock()
	defer mu.Unlock()
	filePath = logFile
	logVerbose = verbose
	restrictLogSize = true
	dm = debugMask

	openFile()
	if logVerbose {
		mw = io.MultiWriter(os.Stdout, file)
	} else {
		mw = file
	}
	setupBackend()
}

func Shutdown() {
	mu.Lock()
	defer mu.Unlock()
	file.Close()
	file = nil
}

func openFile() {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	file = f
}
func setupBackend() {
	i = log.New(mw, "INFO: ", logSettings)
	d = log.New(mw, "DEBUG: ", logSettings)
	e = log.New(mw, "ERROR: ", logSettings)
	w = log.New(mw, "WARNING: ", logSettings)
	f = log.New(mw, "FATAL: ", logSettings)
	t = log.New(mw, "TODO: ", logSettings)
}

func checkLogSize() {
	if !restrictLogSize {
		return
	}

	fhandler, _ := os.Stat(filePath)
	size := fhandler.Size()
	if maxSize < size {
		file.Close()
		moveBackupLogFiles(0)
		os.Rename(filePath, filePath+".0.bak")
		openFile()
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

func Custom(prefix string, message string, values ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	t := log.New(mw, prefix+": ", logSettings)
	t.Println(fmt.Sprintf(message, values...))
	checkLogSize()
}

func Info(message string, values ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	i.Println(fmt.Sprintf(message, values...))
	checkLogSize()
}

func Debug(debugMask debug.Mask, message string, values ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	if dm == 0 {
		return
	}
	if !dm.HasAny(debugMask) {
		return
	}
	d.Println(fmt.Sprintf(message, values...))
	checkLogSize()
}

func Warning(message string, values ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	w.Println(fmt.Sprintf(message, values...))
	checkLogSize()
}

func Error(message string, values ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	e.Println(fmt.Sprintf(message, values...))
	checkLogSize()
}

func Todo(message string, values ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	t.Println(fmt.Sprintf(message, values...))
	checkLogSize()
}

func Fatal(message string, values ...interface{}) {
	mu.Lock()
	defer mu.Unlock()
	f.Println(fmt.Sprintf(message, values...))
	if !dontPanic {
		panic(fmt.Errorf(message, values...))
	}
}
