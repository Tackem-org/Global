package logging_test

import (
	"errors"
	"os"
	"testing"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	"github.com/stretchr/testify/assert"
)

//https://stackoverflow.com/questions/44119951/how-to-check-a-log-output-in-go-test

// func TestsetupBackend(t *testing.T) {

// }

// func TestcheckLogSize(t *testing.T) {

// }

// func TestmoveBackupLogFiles(t *testing.T) {

// }

// func TestfileExists(t *testing.T) {

// }

func TestCustomLogger(t *testing.T) {
	tests := []string{
		"TEST",
		"Custom",
		"GORM",
	}

	for _, v := range tests {
		logger := logging.CustomLogger(v)
		assert.NotNil(t, logger)
		assert.Equal(t, v+": ", logger.Prefix())
	}
}

func TestCustom(t *testing.T) {
	s, r, w := logging.SetupTest(t)
	defer logging.ShutdownTest(t, r, w)
	logging.Custom("TESTING", "TEST")
	s.Scan()
	got := s.Text()
	assert.Equal(t, "TESTING: TEST", got)
}

func TestInfo(t *testing.T) {
	s, r, w := logging.SetupTest(t)
	defer logging.ShutdownTest(t, r, w)
	logging.Info("TEST")
	s.Scan()
	got := s.Text()
	assert.Equal(t, "INFO: TEST", got)
}

func TestDebug(t *testing.T) {
	s, r, w := logging.SetupTest(t)
	defer logging.ShutdownTest(t, r, w)
	logging.DM(debug.FUNCTIONCALLS)
	logging.Debug(debug.FUNCTIONCALLS, "TEST")
	s.Scan()
	got := s.Text()
	assert.Equal(t, "DEBUG: TEST", got)
	logging.DM(debug.FUNCTIONCALLS | debug.FUNCTIONARGS)
	logging.Debug(debug.FUNCTIONCALLS, "TEST")
	s.Scan()
	got = s.Text()
	assert.Equal(t, "DEBUG: TEST", got)
}

func TestWarning(t *testing.T) {
	s, r, w := logging.SetupTest(t)
	defer logging.ShutdownTest(t, r, w)
	logging.Warning("TEST")
	s.Scan()
	got := s.Text()
	assert.Equal(t, "WARNING: TEST", got)
}

func TestError(t *testing.T) {
	s, r, w := logging.SetupTest(t)
	defer logging.ShutdownTest(t, r, w)
	logging.Error("TEST")
	s.Scan()
	got := s.Text()
	assert.Equal(t, "ERROR: TEST", got)
}

func TestTodo(t *testing.T) {
	s, r, w := logging.SetupTest(t)
	defer logging.ShutdownTest(t, r, w)
	logging.Todo("TEST")
	s.Scan()
	got := s.Text()
	assert.Equal(t, "TODO: TEST", got)
}

func TestFatal(t *testing.T) {
	s, r, w := logging.SetupTest(t)
	defer logging.ShutdownTest(t, r, w)
	logging.Fatal("TEST")
	s.Scan()
	got := s.Text()
	assert.Equal(t, "FATAL: TEST", got)
}

func exists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func TestCheckLogSize(t *testing.T) {
	file := "temptest.log"
	os.Remove(file)
	logging.Setup(file, false, debug.NONE)
	logging.MaxSize(100)
	logging.Info("logging Time1")
	logging.Info("logging Time2")
	logging.Info("logging Time3")
	logging.Info("logging Time4")
	logging.Info("logging Time5")
	logging.Info("logging Time6")
	logging.Info("logging Time7")
	assert.True(t, exists(file+".0.bak"))
	os.Remove(file + "*")
}
