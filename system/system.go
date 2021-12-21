package system

import (
	"sync"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/registerService"
	"google.golang.org/grpc"
)

var (
	Data       SetupData
	regData    *registerService.Register
	grpcServer *grpc.Server
	WG         *sync.WaitGroup
	MasterUp   bool
)

type SetupData struct {
	BaseData    registerService.BaseData
	LogFile     string
	VerboseLog  bool
	GPRCSystems func(server *grpc.Server)
	WebSystems  func()
	MainSystem  func()
}

func Shutdown(registered bool) {

	if registered {
		regData.Disconnect()
		logging.Info("DeRegistration Done")
	}

	grpcServer.Stop()
	WG.Done()
	logging.Info("Shutdown gRPC Server")

	logging.Info("Closing Logger")
	logging.Shutdown()

}

func RegData() *registerService.Register {
	if regData == nil {
		regData = registerService.NewRegister()
	}
	return regData
}

func GRPCServer() *grpc.Server {
	if grpcServer == nil {
		grpcServer = grpc.NewServer()
	}
	return grpcServer
}
