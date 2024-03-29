package system

import (
	"fmt"
	"sync"
	"time"

	"github.com/Tackem-org/Global/flags"
	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/structs"
	"github.com/Tackem-org/Global/system/channels"
	"github.com/Tackem-org/Global/system/grpcSystem/clients/registration"
	"github.com/Tackem-org/Global/system/grpcSystem/servers/regClient"
	"github.com/Tackem-org/Global/system/grpcSystem/servers/remoteWeb"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/setupData"
	"google.golang.org/grpc"

	pbrc "github.com/Tackem-org/Global/pb/regclient"
	pbr "github.com/Tackem-org/Global/pb/registration"
	pbrw "github.com/Tackem-org/Global/pb/remoteweb"
)

var (
	VersionString string
	Version       structs.Version
	WG            *sync.WaitGroup = &sync.WaitGroup{}
	Server        *grpc.Server
	WaitTime      time.Duration = time.Duration(5 * time.Second)

	Run              = run
	startup          = startupFunc
	mainLoop         = mainLoopFunc
	shutdown         = shutdownFunc
	connect          = connectFunc
	connectConnector = connectConnectorFunc
	serverServe      = serverServeFunc
)

func run(version string, d *setupData.SetupData) {
	Version, _ = structs.StringToVersion(version)
	VersionString = version
	flags.Parse()
	if flags.Version() {
		fmt.Println(VersionString)
		return
	}
	logging.Setup(flags.LogFolder()+d.Filename("log"), d.VerboseLog)
	defer logging.Shutdown()
	setupData.Data = d
	logging.Info("Starting Tackem %s System", d.Name())
	if setupData.Data.MainSetup != nil {
		setupData.Data.MainSetup()
	}

	if startup() {
		logging.Info("Started Tackem %s System", d.Name())
		mainLoop()

		logging.Info("Stopping Tackem %s System", d.Name())
		shutdown(true)
	}

	if setupData.Data.MainShutdown != nil {
		setupData.Data.MainShutdown()
	}
	WG.Wait()
	logging.Info("Stopped Tackem %s System", d.Name())
}

func startupFunc() bool {
	channels.Setup()
	if !masterData.Setup(flags.ConfigFolder() + setupData.Data.Filename("json")) {
		logging.Fatal("NO REGISTRATION KEY TO USE UNABLE TO START")
	}

	logging.Info("Setup GRPC Service")
	Server = grpc.NewServer()
	pbrc.RegisterRegClientServer(Server, &regClient.RegClientServer{})
	if len(setupData.Data.AdminPaths)+len(setupData.Data.Paths)+len(setupData.Data.Sockets) > 0 {
		pbrw.RegisterRemoteWebServer(Server, &remoteWeb.RemoteWebServer{})
	}

	if setupData.Data.GRPCSystems != nil {
		setupData.Data.GRPCSystems(Server)
	}

	WG.Add(1)
	go serverServe()

	logging.Info("Registration Started")
	if !connect(setupData.Data.RegisterProto()) {
		return false
	}

	masterData.UP.Up()
	logging.Info("Registration Done")
	if setupData.Data.StartActive {
		logging.Info("System Active")
		setupData.Active = true
	}
	return true
}

func serverServeFunc() {
	Server.Serve(setupData.FreeTCPPort())
}

func mainLoopFunc() {
	if setupData.Data.MainSystem == nil {
		select {
		case <-channels.Root.TermChan:
			fmt.Print("\n")
			logging.Info("SIGTERM received. Shutdown process initiated")
		case <-channels.Root.Shutdown:
			logging.Info("Shutdown Command received. Shutdown process initiated")
		}

	} else if !setupData.Data.SingleRun {
		channels.Setup()
		loopBool := true
		for loopBool {
			select {
			case x := <-channels.Root.TermChan:
				fmt.Print("\n")
				logging.Info("%s received. Shutdown process initiated", x.String())
				loopBool = false
			case <-channels.Root.Shutdown:
				logging.Info("Shutdown Command received. Shutdown process initiated")
				loopBool = false
			default:
				setupData.Data.MainSystem()
			}
		}
	} else {
		setupData.Data.MainSystem()
	}

}

func shutdownFunc(registered bool) {
	Server.Stop()
	WG.Done()
	logging.Info("Shutdown gRPC Server")

	if registered && masterData.UP.Check() {
		WG.Add(1)
		logging.Info("Disconnect: %t", masterData.UP.Check())
		disconnectResponse, err := registration.Disconnect(&pbr.DisconnectRequest{})
		if err != nil || !disconnectResponse.Success {
			logging.Warning("failed to disconnect service from master: %s", disconnectResponse.ErrorMessage)
		}
		logging.Info("Disconnect Done")
		WG.Done()
	}
	// WG.Wait()
}

func connectFunc(request *pbr.RegisterRequest) bool {
	connectTimer := time.NewTimer(WaitTime)
	for !connectConnector(request) {
		logging.Info("Master System Is Down Waiting for %d seconds before retrying", WaitTime)
		select {
		case <-channels.Root.TermChan:
			fmt.Print("\n")
			logging.Info("SIGTERM received. Shutdown process initiated")
			return false
		case <-channels.Root.Shutdown:
			logging.Info("Shutdown Command received. Shutdown process initiated")
			return false
		case <-connectTimer.C:
			logging.Info("retrying connection")
		}
	}

	return true
}

func connectConnectorFunc(request *pbr.RegisterRequest) bool {
	response, err := registration.Register(request)
	if err != nil {
		logging.Error(err.Error())
		return false
	}
	if response.Success {
		setupData.BaseID = response.BaseId
		setupData.ServiceID = response.ServiceId

		masterData.ConnectionInfo = masterData.ConnectionInfostruct{
			Key: response.Key,
			IP:  helpers.GrabIPFromURL(masterData.Info.URL),
		}
		return true
	}
	logging.Error(response.ErrorMessage)
	return false
}
