package system

import (
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/Tackem-org/Global/channels"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	pbregclient "github.com/Tackem-org/Proto/pb/regclient"
	pbremoteweb "github.com/Tackem-org/Proto/pb/remoteweb"
	"google.golang.org/grpc"
)

func Run(data SetupData) {
	logging.Setup(data.LogFile, data.VerboseLog, data.DebugLevel)
	defer logging.Shutdown()
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.Run(data SetupData)] {data=%+v}", data)
	logging.Infof("Starting Tackem %s System", data.BaseData.ServiceName)
	healthcheckHealthy = true
	Data = data
	channels.Setup()
	MUp.StartDown()

	logging.Info("Setup Registration Data")
	RegData().Setup(Data.BaseData)

	WG = &sync.WaitGroup{}

	logging.Infof("Setup %s System", data.BaseData.ServiceName)
	if Data.MainSetup != nil {
		Data.MainSetup()
	}

	logging.Info("Setup Web Service")
	if Data.WebSystems != nil {
		Data.WebSystems()
	}

	logging.Info("Setup Web Sockets")
	if Data.WebSockets != nil {
		Data.WebSockets()
	}

	logging.Info("Setup GPRC Service")
	grpcServer = grpc.NewServer()

	pbregclient.RegisterRegClientServer(grpcServer, &RegClientServer{})
	if data.BaseData.WebAccess {
		pbremoteweb.RegisterRemoteWebServer(grpcServer, NewRemoteWebServer())
	}
	Data.GPRCSystems(grpcServer)

	WG.Add(1)
	go func() {
		bind := ""
		if val, ok := os.LookupEnv("BIND"); ok {
			bind = val
		}
		listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", bind, RegData().GetPort()))
		if err != nil {
			logging.Errorf("gPRC could not listen on port %d", RegData().GetPort())
			logging.Fatal(err.Error())
		}
		if err := grpcServer.Serve(listen); err != nil {
			logging.Fatal(err.Error())
		}
	}()
	logging.Info("Starting gRPC server")
	rd := RegData()
	if !rd.Connect() {
		Shutdown(false)
	} else {
		MUp.Up()
		logging.Info("Registration Done")
		rd.Activate()
		logging.Info("System Active")
		mainLoop()
		WG.Wait()
		rd.Deactivate()
	}

	if Data.Shutdown != nil {
		Data.Shutdown()
	}
	logging.Info("Shutdown Complete Exiting Cleanly")
	os.Exit(0)
}

func mainLoop() {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.captureInterupts()]")
	if Data.MainSystem == nil {
		select {
		case <-channels.Root.TermChan:
			fmt.Print("\n")
			logging.Info("SIGTERM received. Shutdown process initiated")
		case <-channels.Root.Shutdown:
			logging.Info("Shutdown Command received. Shutdown process initiated")
		}

	} else if !Data.BaseData.SingleRun {
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
				Data.MainSystem()
			}
		}
	} else {
		Data.MainSystem()
	}
	Shutdown(true)
}

func Shutdown(registered bool) {
	grpcServer.Stop()
	WG.Done()
	logging.Info("Shutdown gRPC Server")
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.Shutdown(registered bool)] {registered=%t}", registered)
	if registered && MUp.Check() {
		logging.Infof("DeRegistration: %t", MUp.Check())
		RegData().Disconnect()
		logging.Info("DeRegistration Done")
	}
}
