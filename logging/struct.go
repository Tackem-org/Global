package logging

import (
	"fmt"
	"io"
	"log"
	"os"
	"sync"

	"github.com/Tackem-org/Global/file"
)

type Logging struct {
	mu              sync.Mutex
	i               *log.Logger
	e               *log.Logger
	w               *log.Logger
	f               *log.Logger
	t               *log.Logger
	file            *os.File
	screen          *os.File
	filePath        string
	restrictLogSize bool
	maxSize         int64
	fileCountLimit  uint8
	logVerbose      bool
	mw              io.Writer
	logSettings     int
}

func DefaultLogging() *Logging {
	return &Logging{
		mu:              sync.Mutex{},
		i:               nil,
		e:               nil,
		w:               nil,
		f:               nil,
		t:               nil,
		file:            nil,
		screen:          os.Stdout,
		filePath:        "",
		restrictLogSize: true,
		maxSize:         50 * 1024,
		fileCountLimit:  5,
		logVerbose:      false,
		mw:              nil,
		logSettings:     log.Ldate | log.Ltime | log.Lmicroseconds,
	}
}

func (l *Logging) Setup(logFile string, verbose bool) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.filePath = logFile
	l.logVerbose = verbose
	l.setupBackend()

}

func (l *Logging) Shutdown() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.file.Close()
	l.file = nil
}

func (l *Logging) FileCountLimit() uint8 {
	return l.fileCountLimit
}

func (l *Logging) FileClosed() bool {
	return l.file == nil
}

func (l *Logging) LogSettings(logSettings int) {
	l.logSettings = logSettings
}

func (l *Logging) Writer(mw io.Writer) {
	l.mw = mw
}

func (l *Logging) setupBackend() {
	f, err := os.OpenFile(l.filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		panic(err)
	}

	l.file = f

	if l.logVerbose {
		l.mw = io.MultiWriter(l.screen, l.file)
	} else {
		l.mw = l.file
	}

	l.SetupLoggers()
}

func (l *Logging) SetupLoggers() {
	l.i = log.New(l.mw, "INFO: ", l.logSettings)
	l.e = log.New(l.mw, "ERROR: ", l.logSettings)
	l.w = log.New(l.mw, "WARNING: ", l.logSettings)
	l.f = log.New(l.mw, "FATAL: ", l.logSettings)
	l.t = log.New(l.mw, "TODO: ", l.logSettings)
}

func checkLogSize(l *Logging) {
	if l.filePath == "" {
		return
	}
	fhandler, _ := os.Stat(l.filePath)
	size := fhandler.Size()
	if l.maxSize < size {
		l.file.Close()
		moveBackupLogFiles(l, 0)
		os.Rename(l.filePath, l.filePath+".0.bak")
		l.setupBackend()
	}
}

func moveBackupLogFiles(l *Logging, i uint8) {
	if !file.FileExists(fmt.Sprintf("%s.%d.bak", l.filePath, i)) {
		return
	}
	if l.fileCountLimit > 0 && i >= l.fileCountLimit {
		os.Remove(fmt.Sprintf("%s.%d.bak", l.filePath, i))
		return
	}

	moveBackupLogFiles(l, i+1)
	os.Rename(fmt.Sprintf("%s.%d.bak", l.filePath, i), fmt.Sprintf("%s.%d.bak", l.filePath, i+1))
}

func (l *Logging) CustomLogger(prefix string) *log.Logger {
	l.mu.Lock()
	defer l.mu.Unlock()
	return log.New(l.mw, prefix+": ", l.logSettings)
}

func (l *Logging) Custom(prefix string, message string, values ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	t := log.New(l.mw, prefix+": ", l.logSettings)
	t.Println(fmt.Sprintf(message, values...))
	checkLogSize(l)
}

func (l *Logging) Info(message string, values ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.i.Println(fmt.Sprintf(message, values...))
	checkLogSize(l)
}

func (l *Logging) Warning(message string, values ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.w.Println(fmt.Sprintf(message, values...))
	checkLogSize(l)
}

func (l *Logging) Error(message string, values ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.e.Println(fmt.Sprintf(message, values...))
	checkLogSize(l)
}

func (l *Logging) Todo(message string, values ...interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.t.Println(fmt.Sprintf(message, values...))
	checkLogSize(l)
}

func (l *Logging) Fatal(message string, values ...interface{}) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.f.Println(fmt.Sprintf(message, values...))
	return fmt.Errorf(message, values...)
}
