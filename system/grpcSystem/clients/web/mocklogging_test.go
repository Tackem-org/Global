package web_test

import (
	"fmt"
	"log"
)

type MockLogging struct{}

func (l *MockLogging) Setup(logFile string, verbose bool)                          {}
func (l *MockLogging) Shutdown()                                                   {}
func (l *MockLogging) CustomLogger(prefix string) *log.Logger                      { return log.New(nil, prefix+": ", 0) }
func (l *MockLogging) Custom(prefix string, message string, values ...interface{}) {}
func (l *MockLogging) Info(message string, values ...interface{})                  {}
func (l *MockLogging) Warning(message string, values ...interface{})               {}
func (l *MockLogging) Error(message string, values ...interface{})                 {}
func (l *MockLogging) Todo(message string, values ...interface{})                  {}
func (l *MockLogging) Fatal(message string, values ...interface{}) error {
	return fmt.Errorf(message, values...)
}
