package system

import (
	"github.com/Tackem-org/Global/logging"
)

func Shutdown(registered bool) {

	if registered {
		RegData().Disconnect()
		logging.Info("DeRegistration Done")
	}

	grpcServer.Stop()
	WG.Done()
	logging.Info("Shutdown gRPC Server")

	logging.Info("Closing Logger")
	logging.Shutdown()

}
