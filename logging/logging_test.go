package logging_test

import (
	"os"
	"testing"

	"github.com/Tackem-org/Global/logging"
	"github.com/stretchr/testify/assert"
)

func TestLoggingInterface(t *testing.T) {
	filename := "temp.log"
	assert.NotPanics(t, func() { logging.Setup(filename, false) })
	assert.NotNil(t, logging.I)
	assert.NotNil(t, logging.CustomLogger("Test"))
	assert.NotPanics(t, func() { logging.Custom("Test", "Test") })
	assert.NotPanics(t, func() { logging.Info("Test") })
	assert.NotPanics(t, func() { logging.Warning("Test") })
	assert.NotPanics(t, func() { logging.Error("Test") })
	assert.NotPanics(t, func() { logging.Todo("Test") })
	assert.Panics(t, func() { logging.Fatal("Test") })
	assert.NotPanics(t, func() { logging.Shutdown() })
	assert.True(t, exists(filename))
	os.Remove(filename)
	assert.False(t, exists(filename))
}
