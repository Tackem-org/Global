package system

import (
	"github.com/Tackem-org/Global/logging"
)

func Shutdown(registered bool) {

	if registered && MUp.Check() {
		logging.Infof("DeRegistration: %t", MUp.Check())
		RegData().Disconnect()
		logging.Info("DeRegistration Done")
	}

	grpcServer.Stop()
	WG.Done()
	logging.Info("Shutdown gRPC Server")

	logging.Info("Closing Logger")
	logging.Shutdown()

}
