package system

import (
	"fmt"

	"github.com/Tackem-org/Global/logging"
)

func Shutdown(registered bool, fullShutdown bool) {

	if registered && MUp.Check() {
		logging.Info("DeRegistration:" + fmt.Sprintf("%t", MUp.Check()))
		RegData().Disconnect(fullShutdown)
		logging.Info("DeRegistration Done")
	}

	grpcServer.Stop()
	WG.Done()
	logging.Info("Shutdown gRPC Server")

	logging.Info("Closing Logger")
	logging.Shutdown()

}
