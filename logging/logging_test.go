package logging_test

import (
	"os"
	"testing"

	"github.com/Tackem-org/Global/logging"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestLoggingInterface(t *testing.T) {
	suite.Run(t, new(LoggingInterfaceNilSuite))
	suite.Run(t, new(LoggingInterfaceSuite))
}

type LoggingInterfaceNilSuite struct {
	suite.Suite
}

func (s *LoggingInterfaceNilSuite) SetupTest() {
	logging.LI = nil
}

func (s *LoggingInterfaceNilSuite) TearDownTest() {
	assert.Nil(s.T(), logging.LI)
}

func (s *LoggingInterfaceNilSuite) TestInterfaceShutdown() {
	assert.NotPanics(s.T(), func() { logging.Shutdown() })
}
func (s *LoggingInterfaceNilSuite) TestInterfaceCustomLogger() {
	assert.NotPanics(s.T(), func() { assert.Nil(s.T(), logging.CustomLogger("Test")) })
}

func (s *LoggingInterfaceNilSuite) TestInterfaceCustom() {
	assert.NotPanics(s.T(), func() { logging.Custom("Test", "Test") })

}

func (s *LoggingInterfaceNilSuite) TestInterfaceInfo() {
	assert.NotPanics(s.T(), func() { logging.Info("Test") })
}

func (s *LoggingInterfaceNilSuite) TestInterfaceWarning() {
	assert.NotPanics(s.T(), func() { logging.Warning("Test") })
}

func (s *LoggingInterfaceNilSuite) TestInterfaceError() {
	assert.NotPanics(s.T(), func() { logging.Error("Test") })
}

func (s *LoggingInterfaceNilSuite) TestInterfaceTodo() {
	assert.NotPanics(s.T(), func() { logging.Todo("Test") })
}

func (s *LoggingInterfaceNilSuite) TestInterfaceFatal() {
	assert.NotPanics(s.T(), func() { logging.Fatal("Test") })
}

type LoggingInterfaceSuite struct {
	suite.Suite
	filename string
}

func (s *LoggingInterfaceSuite) SetupSuite() {
	s.filename = "temp.log"
	assert.NotPanics(s.T(), func() { logging.Setup(s.filename, false) })
	assert.NotNil(s.T(), logging.LI)
}

func (s *LoggingInterfaceSuite) TearDownSuite() {
	assert.NotPanics(s.T(), func() { logging.Shutdown() })
	assert.True(s.T(), exists(s.filename))
	os.Remove(s.filename)
	assert.False(s.T(), exists(s.filename))
}

func (s *LoggingInterfaceSuite) TestInterfaceCustomLogger() {
	assert.NotPanics(s.T(), func() { assert.NotNil(s.T(), logging.CustomLogger("Test")) })
}

func (s *LoggingInterfaceSuite) TestInterfaceCustom() {
	assert.NotPanics(s.T(), func() { logging.Custom("Test", "Test") })

}

func (s *LoggingInterfaceSuite) TestInterfaceInfo() {
	assert.NotPanics(s.T(), func() { logging.Info("Test") })
}

func (s *LoggingInterfaceSuite) TestInterfaceWarning() {
	assert.NotPanics(s.T(), func() { logging.Warning("Test") })
}

func (s *LoggingInterfaceSuite) TestInterfaceError() {
	assert.NotPanics(s.T(), func() { logging.Error("Test") })
}

func (s *LoggingInterfaceSuite) TestInterfaceTodo() {
	assert.NotPanics(s.T(), func() { logging.Todo("Test") })
}

func (s *LoggingInterfaceSuite) TestInterfaceFatal() {
	assert.Panics(s.T(), func() { logging.Fatal("Test") })
}
