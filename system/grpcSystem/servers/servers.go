package servers

import (
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
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
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Master.grpcSystem.servers.Setup")
	logging.Debugf(debug.FUNCTIONARGS, "[FUNCTIONARGS] WaitGroup webEnabled=%T", webEnabled)
	server = grpc.NewServer()
	pbrc.RegisterRegClientServer(server, &regClient.RegClientServer{})
	if webEnabled {
		pbrw.RegisterRemoteWebServer(server, &remoteWeb.RemoteWebServer{})
	}
	setupData.Data.GRPCSystems(server)

	wg.Add(1)
	go func() {
		bind := ""
		if val, ok := os.LookupEnv("BIND"); ok {
			bind = val
		}
		listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", bind, setupData.Port))
		if err != nil {
			logging.Errorf("GRPC could not listen on port %d", setupData.Port)
			logging.Fatal(err.Error())
		}
		if err := server.Serve(listen); err != nil {
			logging.Fatal(err.Error())
		}
	}()
}

func Shutdown(wg *sync.WaitGroup) {
	logging.Debug(debug.FUNCTIONCALLS, "[FUNCTIONCALL] Master.grpcSystem.servers.Shutdown")
	server.Stop()
	wg.Done()
	logging.Info("Shutdown gRPC Server")
}
