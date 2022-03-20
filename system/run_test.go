package system_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/system"
	"github.com/Tackem-org/Global/system/setupData"
	"github.com/stretchr/testify/assert"
)

type MockLogging struct{}

func (ml *MockLogging) Setup(logFile string, verbose bool) {}
func (ml *MockLogging) Shutdown()                          {}
func (ml *MockLogging) CustomLogger(prefix string) *log.Logger {
	return nil
}
func (ml *MockLogging) Custom(prefix string, message string, values ...interface{}) {}
func (ml *MockLogging) Info(message string, values ...interface{})                  {}
func (ml *MockLogging) Warning(message string, values ...interface{})               {}
func (ml *MockLogging) Error(message string, values ...interface{})                 {}
func (ml *MockLogging) Todo(message string, values ...interface{})                  {}
func (ml *MockLogging) Fatal(message string, values ...interface{}) error {
	return fmt.Errorf(message, values...)
}

func mockTrueStartup() bool        { return true }
func mockFalseStartup() bool       { return false }
func mockMainLoop()                {}
func mockShutdown(registered bool) {}

func TestRun(t *testing.T) {
	logging.I = &MockLogging{}
	system.SetCalls(mockTrueStartup, mockMainLoop, mockShutdown)
	assert.NotPanics(t,
		func() {
			system.Run(&setupData.SetupData{
				LogFile:      "",
				VerboseLog:   false,
				MainSetup:    func() {},
				MainShutdown: func() {},
			})
		})
}

func Test_startup(t *testing.T) {

}

func Test_mainLoop(t *testing.T) {

}

func Test_shutdown(t *testing.T) {

}

func Test_connect(t *testing.T) {

}
