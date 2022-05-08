package system_test

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"
	"time"

	"github.com/Tackem-org/Global/logging"
	pbr "github.com/Tackem-org/Global/pb/registration"
	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/Global/system"
	"github.com/Tackem-org/Global/system/channels"
	"github.com/Tackem-org/Global/system/grpcSystem/clients/registration"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/setupData"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
)

type MockLogging struct {
	InfoCount    int
	ErrorCount   int
	WarningCount int
}

func (l *MockLogging) Reset() {
	l.InfoCount = 0
	l.ErrorCount = 0
	l.WarningCount = 0
}
func (l *MockLogging) Setup(logFile string, verbose bool)                          {}
func (l *MockLogging) Shutdown()                                                   {}
func (l *MockLogging) CustomLogger(prefix string) *log.Logger                      { return log.New(nil, prefix+": ", 0) }
func (l *MockLogging) Custom(prefix string, message string, values ...interface{}) {}
func (l *MockLogging) Info(message string, values ...interface{})                  { l.InfoCount++ }
func (l *MockLogging) Warning(message string, values ...interface{})               { l.WarningCount++ }
func (l *MockLogging) Error(message string, values ...interface{})                 { l.ErrorCount++ }
func (l *MockLogging) Todo(message string, values ...interface{})                  {}
func (l *MockLogging) Fatal(message string, values ...interface{}) error {
	return fmt.Errorf(message, values...)
}

type MockRegistrationClient struct {
	Success bool
	Error   bool
}

func (mrc *MockRegistrationClient) Activate(request *pbr.ActivateRequest) (*pbr.ActivateResponse, error) {
	return nil, nil
}
func (mrc *MockRegistrationClient) Deactivate(request *pbr.DeactivateRequest) (*pbr.DeactivateResponse, error) {
	return nil, nil
}
func (mrc *MockRegistrationClient) Deregister(request *pbr.DeregisterRequest) (*pbr.DeregisterResponse, error) {
	return nil, nil
}
func (mrc *MockRegistrationClient) Disconnect(request *pbr.DisconnectRequest) (*pbr.DisconnectResponse, error) {
	return &pbr.DisconnectResponse{Success: mrc.Success}, nil
}
func (mrc *MockRegistrationClient) Register(request *pbr.RegisterRequest) (*pbr.RegisterResponse, error) {
	if mrc.Error {
		return nil, errors.New("fail")
	}
	return &pbr.RegisterResponse{
		Success:      mrc.Success,
		ErrorMessage: "",
		BaseId:       "",
		ServiceId:    0,
		Key:          "",
	}, nil
}

func TestRun(t *testing.T) {
	system.ResetFunctions()
	logging.I = &MockLogging{}
	system.SetupForRun(
		func() bool { return true },
		func() {},
		func(registered bool) {},
	)
	assert.NotPanics(t,
		func() {
			system.Run(&setupData.SetupData{
				LogFile:      "",
				VerboseLog:   false,
				MainSetup:    func() {},
				MainShutdown: func() {},
			}, "data.json", "log.log")
		})

	pflag.Set("version", "true")

	assert.NotPanics(t,
		func() {
			system.Run(&setupData.SetupData{
				LogFile:      "",
				VerboseLog:   false,
				MainSetup:    func() {},
				MainShutdown: func() {},
			}, "data.json", "log.log")
		})

	system.ResetFunctions()
}

func TestStartup(t *testing.T) {
	setupData.Data = &setupData.SetupData{
		StartActive: true,
		MasterConf:  "/missing",
		GRPCSystems: func(server *grpc.Server) {},
		Paths: []*setupData.PathItem{
			{
				Path:        "test",
				Permission:  "",
				PostAllowed: false,
				GetDisabled: false,
				Call: func(in *structs.WebRequest) (*structs.WebReturn, error) {
					return nil, nil
				},
			},
		},
	}
	logging.I = &MockLogging{}
	system.SetupForStartup(
		func(request *pbr.RegisterRequest) bool { return false },
		func() {},
	)
	masterData.Setup = func(masterConf string) bool { return false }
	assert.Panics(t, func() { system.StartupFunc() })
	masterData.Setup = func(masterConf string) bool { return true }

	os.Setenv("REGKEY", "Key")
	os.Setenv("URL", "localhost")
	os.Setenv("PORT", "0")
	masterData.UP.Down()
	assert.False(t, system.StartupFunc())

	system.SetupForStartup(func(request *pbr.RegisterRequest) bool { return true }, func() {})
	assert.True(t, system.StartupFunc())
	system.ResetFunctions()

}

func TestServerServeFunc(t *testing.T) {
	go system.ServerServe()
	assert.NotPanics(t, func() { system.Server.Stop() })
}
func TestMainLoop(t *testing.T) {
	mainsysFunc := func() { logging.Info("mainsystemRan") }
	system.ResetFunctions()
	setupData.Data = &setupData.SetupData{MainSystem: nil}
	channels.Setup()
	l := &MockLogging{}
	logging.I = l
	go system.MainLoop()
	channels.Root.TermChan <- os.Interrupt
	time.Sleep(time.Millisecond)
	assert.Equal(t, 1, l.InfoCount)
	l.InfoCount = 0
	go system.MainLoop()
	channels.Root.Shutdown <- true
	time.Sleep(time.Millisecond)
	assert.Equal(t, 1, l.InfoCount)
	l.InfoCount = 0
	setupData.Data = &setupData.SetupData{MainSystem: mainsysFunc, SingleRun: false}
	go system.MainLoop()
	time.Sleep(time.Millisecond)
	channels.Root.TermChan <- os.Interrupt
	time.Sleep(time.Millisecond * 2)
	assert.NotZero(t, l.InfoCount)
	l.InfoCount = 0
	go system.MainLoop()
	channels.Root.Shutdown <- true
	time.Sleep(time.Millisecond)
	assert.NotZero(t, l.InfoCount)
	l.InfoCount = 0
	setupData.Data = &setupData.SetupData{MainSystem: mainsysFunc, SingleRun: true}
	system.MainLoop()
	assert.Equal(t, 1, l.InfoCount)
	system.ResetFunctions()
}

func TestShutdown(t *testing.T) {
	system.ResetFunctions()
	l := &MockLogging{}
	logging.I = l
	mrc := &MockRegistrationClient{Success: false}
	registration.I = mrc
	system.Server = grpc.NewServer()

	system.WG.Add(1)
	system.Shutdown(true)
	assert.Equal(t, 3, l.InfoCount)
	assert.Equal(t, 1, l.WarningCount)
}

func TestConnect(t *testing.T) {
	system.ResetFunctions()
	channels.Setup()
	l := &MockLogging{}
	logging.I = l
	loopCount := 0
	system.SetupForConnect(func(request *pbr.RegisterRequest) bool { return true })
	assert.True(t, system.Connect(nil))

	system.SetupForConnect(func(request *pbr.RegisterRequest) bool { return false })
	channels.Root.TermChan <- os.Interrupt
	assert.False(t, system.Connect(nil))

	channels.Root.Shutdown <- true
	assert.False(t, system.Connect(nil))

	system.SetupForConnect(func(request *pbr.RegisterRequest) bool {
		switch loopCount {
		case 1:
			return true
		default:
			loopCount++
			return false
		}
	})
	system.WaitTime = time.Microsecond
	assert.True(t, system.Connect(nil))
	system.ResetFunctions()
}

func Test_connectConnectorFunc(t *testing.T) {
	logging.I = &MockLogging{}
	masterData.Info = masterData.Infostruct{
		URL: "127.0.0.1",
	}
	system.ResetFunctions()
	mrc := &MockRegistrationClient{Success: false, Error: true}
	registration.I = mrc

	assert.False(t, system.ConnectConnector(&pbr.RegisterRequest{}))
	mrc.Error = false
	assert.False(t, system.ConnectConnector(&pbr.RegisterRequest{}))
	mrc.Success = true
	assert.True(t, system.ConnectConnector(&pbr.RegisterRequest{}))
}
