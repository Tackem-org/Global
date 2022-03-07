package system

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/Tackem-org/Global/channels"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/system/grpcSystem/clients/registration"
	"github.com/Tackem-org/Global/system/grpcSystem/servers"
	"github.com/Tackem-org/Global/system/masterData"
	"github.com/Tackem-org/Global/system/setupData"

	pb "github.com/Tackem-org/Proto/pb/registration"
)

var (
	WG *sync.WaitGroup = &sync.WaitGroup{}
)

func Run(d *setupData.SetupData) {
	logging.Setup(d.LogFile, d.VerboseLog, d.DebugLevel)
	defer logging.Shutdown()
	setupData.Data = d
	logging.Infof("Starting Tackem %s System", d.Name())
	if setupData.Data.MainSetup != nil {
		setupData.Data.MainSetup()
	}
	startup()
	logging.Infof("Started Tackem %s System", d.Name())

	mainLoop()

	logging.Infof("Stopping Tackem %s System", d.Name())
	Shutdown(true)
	if setupData.Data.MainShutdown != nil {
		setupData.Data.MainShutdown()
	}
	logging.Infof("Stopped Tackem %s System", d.Name())

	os.Exit(0)
}
func startup() {
	channels.Setup()
	masterData.Setup()

	logging.Info("Setup GRPC Service")
	servers.Setup(WG, len(setupData.Data.AdminPaths)+len(setupData.Data.Paths)+len(setupData.Data.Sockets) > 0)

	waitTime := time.Duration(5)
	for !connect(setupData.Data.RegisterProto()) {
		logging.Infof("Master System Is Down Waiting for %d seconds before retrying", waitTime)
		time.Sleep(waitTime * time.Second)
	}

	masterData.UP.Up()
	logging.Info("Registration Done")
	if setupData.Data.StartActive {
		logging.Info("System Active")
		setupData.Active = true
	}
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
	servers.Shutdown(WG)
	if registered && masterData.UP.Check() {
		WG.Add(1)
		logging.Infof("Disconnect: %t", masterData.UP.Check())
		disconnectResponse, err := registration.Disconnect(&pb.DisconnectRequest{
			BaseId: setupData.BaseID,
		})
		if err != nil || !disconnectResponse.Success {
			logging.Warningf("failed to disconnect service from master: %s", disconnectResponse.ErrorMessage)
		}
		logging.Info("Disconnect Done")
		WG.Done()
	}
	WG.Wait()
}

func connect(request *pb.RegisterRequest) bool {
	response, err := registration.Register(request)
	if err != nil {
		logging.Fatal(err.Error())
	}
	if response.Success {
		setupData.BaseID = response.BaseId
		setupData.ServiceID = response.ServiceId
		setupData.Key = response.Key
		return true
	}
	logging.Error(response.ErrorMessage)
	return false
}
