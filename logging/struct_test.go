package logging_test

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/Tackem-org/Global/logging"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func TestLoggingBackend(t *testing.T) {
	filename := "temptest.log"
	l := logging.DefaultLogging()
	os.Remove(filename)
	l.Setup(filename, false)

	for i := uint8(0); i <= l.FileCountLimit(); i++ {
		for m := 0; m < 1376; m++ {
			l.Info("%d", m)
		}
		if int8(i)-1 > 0 {
			assert.True(t, exists(fmt.Sprintf("%s.%d.bak", filename, i-1)))
		}
		if int8(i) > 0 {
			assert.False(t, exists(fmt.Sprintf("%s.%d.bak", filename, i)))
		}
	}

	for i := 0; i < 1377*3; i++ {
		l.Info("%d", i)
	}
	assert.False(t, exists(filename+".6.bak"))

	os.Remove(filename)
	for i := uint8(0); i <= l.FileCountLimit(); i++ {
		os.Remove(fmt.Sprintf("%s.%d.bak", filename, i))
	}
	l.Shutdown()
	assert.True(t, l.FileClosed())
}

func TestCustomLogger(t *testing.T) {
	tests := []string{
		"TEST",
		"Custom",
		"GORM",
	}

	l := logging.DefaultLogging()
	for _, v := range tests {
		logger := l.CustomLogger(v)
		assert.NotNil(t, logger)
		assert.Equal(t, v+": ", logger.Prefix())
	}
}

func TestLoggingCallsSuite(t *testing.T) {
	suite.Run(t, new(LoggingCallsSuite))
}

type LoggingCallsSuite struct {
	suite.Suite
	l *logging.Logging
	r *os.File
	w *os.File
	b *bufio.Scanner
}

func (s *LoggingCallsSuite) SetupTest() {
	s.l = logging.DefaultLogging()
	s.l.LogSettings(0)
	r, w, err := os.Pipe()
	if err != nil {
		assert.Fail(s.T(), "couldn't get os Pipe: %v", err)
	}
	s.l.Writer(w)
	s.w = w
	s.r = r
	s.b = bufio.NewScanner(r)
	s.l.SetupLoggers()
}

func (s *LoggingCallsSuite) TearDownTest() {
	err := s.r.Close()
	if err != nil {
		assert.Fail(s.T(), "error closing reader was: %v ", err)
	}
	if err = s.w.Close(); err != nil {
		assert.Fail(s.T(), "error closing writer was: %v ", err)
	}
}

func (s *LoggingCallsSuite) TestInfo() {
	s.l.Info("TEST")
	s.b.Scan()
	got := s.b.Text()
	assert.Equal(s.T(), "INFO: TEST", got)
}

func (s *LoggingCallsSuite) TestCustom() {
	s.l.Custom("TESTING", "TEST")
	s.b.Scan()
	got := s.b.Text()
	assert.Equal(s.T(), "TESTING: TEST", got)
}

func (s *LoggingCallsSuite) TestWarning() {
	s.l.Warning("TEST")
	s.b.Scan()
	got := s.b.Text()
	assert.Equal(s.T(), "WARNING: TEST", got)
}

func (s *LoggingCallsSuite) TestError() {
	s.l.Error("TEST")
	s.b.Scan()
	got := s.b.Text()
	assert.Equal(s.T(), "ERROR: TEST", got)
}

func (s *LoggingCallsSuite) TestTodo() {
	s.l.Todo("TEST")
	s.b.Scan()
	got := s.b.Text()
	assert.Equal(s.T(), "TODO: TEST", got)
}

func (s *LoggingCallsSuite) TestFatal() {
	s.l.Fatal("TEST")
	s.b.Scan()
	got := s.b.Text()
	assert.Equal(s.T(), "FATAL: TEST", got)
}

func TestPanicIfFileCannotOpen(t *testing.T) {
	unreachableFile := "/temptest.log"
	l := logging.DefaultLogging()
	assert.Panics(t, func() { l.Setup(unreachableFile, false) })
}

func TestLogVerboseMode(t *testing.T) {
	filename := "temptest.log"
	l := logging.DefaultLogging()
	assert.NotPanics(t, func() { l.Setup(filename, true) })
	os.Remove(filename)
	l.Shutdown()
	assert.True(t, l.FileClosed())
}

// go test -cover -v -coverprofile=cover.out
// go tool cover -html=cover.out -o coverage.html
