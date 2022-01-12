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
	mu             sync.RWMutex
	dmu            sync.RWMutex
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
		e := os.Rename(filePath, filePath+".0.bak")
		if e != nil {
			Fatal(e)
		}
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
	e := os.Rename(fmt.Sprintf("%s.%d.bak", filePath, i), fmt.Sprintf("%s.%d.bak", filePath, i+1))
	if e != nil {
		Fatal(e)
	}

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
	t := log.New(mw, prefix+": ", logSettings)
	t.Println(message)
}

func Info(message string) {
	mu.Lock()
	defer mu.Unlock()
	i.Println(message)
	checkLogSize()
}

func Debug(debugMask debug.Mask, message string) {
	if dm == 0 {
		return
	}
	dmu.Lock()
	defer dmu.Unlock()
	if !dm.HasAny(debugMask) {
		return
	}
	mu.Lock()
	defer mu.Unlock()
	d.Println(message)
	checkLogSize()
}

func Warning(message string) {
	mu.Lock()
	defer mu.Unlock()
	w.Println(message)
	checkLogSize()
}

func Error(message string) {
	mu.Lock()
	defer mu.Unlock()
	e.Println(message)
	checkLogSize()
}

func Fatal(err error) {
	mu.Lock()
	defer mu.Unlock()
	f.Println(err.Error())
	panic(err)
}

func Todo(message string) {
	mu.Lock()
	defer mu.Unlock()
	t.Println(message)
	checkLogSize()
}
