package logging

import (
	"bufio"
	"os"
	"testing"

	"github.com/Tackem-org/Global/logging/debug"
	"github.com/stretchr/testify/assert"
)

var CheckLogSize = checkLogSize

func SetupTest(t *testing.T) (*bufio.Scanner, *os.File, *os.File) {
	r, w, err := os.Pipe()
	if err != nil {
		assert.Fail(t, "couldn't get os Pipe: %v", err)
	}
	mw = w
	logSettings = 0
	restrictLogSize = false
	dontPanic = true
	setupBackend()
	return bufio.NewScanner(r), r, w
}

func ShutdownTest(t *testing.T, r *os.File, w *os.File) {
	err := r.Close()
	if err != nil {
		assert.Fail(t, "error closing reader was: %v ", err)
	}
	if err = w.Close(); err != nil {
		assert.Fail(t, "error closing writer was: %v ", err)
	}
}

func DM(set debug.Mask) {
	dm = set
}

func MaxSize(size int64) {
	maxSize = size
}

func FileSize() int64 {
	fhandler, _ := os.Stat(filePath)
	return fhandler.Size()
}
