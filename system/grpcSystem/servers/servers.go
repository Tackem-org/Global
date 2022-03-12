package servers

import (
	"sync"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/system/grpcSystem/servers/regClient"
	"github.com/Tackem-org/Global/system/grpcSystem/servers/remoteWeb"
	"github.com/Tackem-org/Global/system/setupData"
	pbrc "github.com/Tackem-org/Proto/pb/regclient"
	pbrw "github.com/Tackem-org/Proto/pb/remoteweb"

	"google.golang.org/grpc"
)

var (
	server *grpc.Server
)

func Setup(wg *sync.WaitGroup, webEnabled bool) {
	server = grpc.NewServer()
	pbrc.RegisterRegClientServer(server, &regClient.RegClientServer{})
	if webEnabled {
		pbrw.RegisterRemoteWebServer(server, &remoteWeb.RemoteWebServer{})
	}
	setupData.Data.GRPCSystems(server)

	wg.Add(1)
	go func() {
		listen, err := setupData.FreeTCPPort()
		if err != nil {
			logging.Fatal("GRPC Error: %s", err.Error())
		}
		if listen == nil {
			logging.Fatal("GRPC Error NO TCP LISTENER")
		}
		if err := server.Serve(listen); err != nil {
			logging.Fatal("GRPC Error CANNOT SERVER ON LISTENER: %s", err.Error())
		}
	}()
}

func Shutdown(wg *sync.WaitGroup) {
	server.Stop()
	wg.Done()
	logging.Info("Shutdown gRPC Server")
}
