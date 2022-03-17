package system

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/Tackem-org/Global/channels"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/system/grpcSystem/clients/registration"
	"github.com/Tackem-org/Global/system/grpcSystem/servers/regClient"
	"github.com/Tackem-org/Global/system/grpcSystem/servers/remoteWeb"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/setupData"
	"google.golang.org/grpc"

	pbrc "github.com/Tackem-org/Proto/pb/regclient"
	pbr "github.com/Tackem-org/Proto/pb/registration"
	pbrw "github.com/Tackem-org/Proto/pb/remoteweb"
)

var (
	wg     *sync.WaitGroup = &sync.WaitGroup{}
	server *grpc.Server
)

func Run(d *setupData.SetupData) {
	logging.Setup(d.LogFile, d.VerboseLog)
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
		Shutdown(true)
	}

	if setupData.Data.MainShutdown != nil {
		setupData.Data.MainShutdown()
	}
	logging.Info("Stopped Tackem %s System", d.Name())

	os.Exit(0)
}

func startup() bool {
	channels.Setup()
	if !masterData.Setup(setupData.Data.MasterConf) {
		logging.Fatal("NO REGISTRATION KEY TO USE UNABLE TO START")
		return false
	}

	logging.Info("Setup GRPC Service")
	server = grpc.NewServer()
	pbrc.RegisterRegClientServer(server, &regClient.RegClientServer{})
	if len(setupData.Data.AdminPaths)+len(setupData.Data.Paths)+len(setupData.Data.Sockets) > 0 {
		pbrw.RegisterRemoteWebServer(server, &remoteWeb.RemoteWebServer{})
	}
	setupData.Data.GRPCSystems(server)

	wg.Add(1)
	go func() {
		if err := server.Serve(setupData.FreeTCPPort()); err != nil {
			logging.Fatal("GRPC Error CANNOT SERVER ON LISTENER: %s", err.Error())
		}
	}()

	waitTime := time.Duration(5)
	for !connect(setupData.Data.RegisterProto()) {
		select {
		case <-channels.Root.TermChan:
			fmt.Print("\n")
			logging.Info("SIGTERM received. Shutdown process initiated")
			return false
		case <-channels.Root.Shutdown:
			logging.Info("Shutdown Command received. Shutdown process initiated")
			return false
		default:
			logging.Info("Master System Is Down Waiting for %d seconds before retrying", waitTime)
			time.Sleep(waitTime * time.Second)
		}
	}

	masterData.UP.Up()
	logging.Info("Registration Done")
	if setupData.Data.StartActive {
		logging.Info("System Active")
		setupData.Active = true
	}
	return true
}

func mainLoop() {
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
			case <-channels.Root.TermChan:
				fmt.Print("\n")
				logging.Info("SIGTERM received. Shutdown process initiated")
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

func Shutdown(registered bool) {
	server.Stop()
	wg.Done()
	logging.Info("Shutdown gRPC Server")

	if registered && masterData.UP.Check() {
		wg.Add(1)
		logging.Info("Disconnect: %t", masterData.UP.Check())
		disconnectResponse, err := registration.Disconnect(&pbr.DisconnectRequest{})
		if err != nil || !disconnectResponse.Success {
			logging.Warning("failed to disconnect service from master: %s", disconnectResponse.ErrorMessage)
		}
		logging.Info("Disconnect Done")
		wg.Done()
	}
	wg.Wait()
}

func connect(request *pbr.RegisterRequest) bool {
	response, err := registration.Register(request)
	if err != nil {
		logging.Error(err.Error())
	}
	if response.Success {
		setupData.BaseID = response.BaseId
		setupData.ServiceID = response.ServiceId

		masterData.ConnectionInfo = masterData.ConnectionInfostruct{
			Key: response.Key,
			IP:  masterData.GrabIPFromURL(masterData.Info.URL),
		}
		return true
	}
	logging.Error(response.ErrorMessage)
	return false
}
