package system

import (
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
)

func Shutdown(registered bool) {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.Shutdown(registered bool)] {registered=%t}", registered)
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
