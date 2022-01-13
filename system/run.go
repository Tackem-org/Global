package system

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	pbregclient "github.com/Tackem-org/Proto/pb/regclient"
	pbremoteweb "github.com/Tackem-org/Proto/pb/remoteweb"
	"google.golang.org/grpc"
)

func Run(data SetupData) {
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.Run(data SetupData)] {data=%+v}", data)
	Data = data
	healthcheckHealthy = true
	fmt.Printf("Starting Tackem %s System\n", Data.BaseData.ServiceName)
	MUp.StartDown()
	if !helpers.InDockerCheck() {
		return
	}
	fmt.Printf("Verbose %t\n", Data.VerboseLog)
	logging.Setup(Data.LogFile, Data.VerboseLog, Data.DebugLevel)
	logging.Info("Logger Started")

	logging.Info("Setup Registration Data")
	RegData().Setup(Data.BaseData)

	WG = &sync.WaitGroup{}

	logging.Infof("Setup %s System", data.BaseData.ServiceName)
	if Data.MainSystem != nil {
		Data.MainSystem()
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
		listen, err := net.Listen("tcp", fmt.Sprintf(":%d", RegData().GetPort()))
		if err != nil {
			logging.Errorf("gPRC could not listen on port %d", RegData().GetPort())
			logging.Fatal(err.Error())
		}
		if err := grpcServer.Serve(listen); err != nil {
			logging.Fatal(err.Error())
		}
	}()
	logging.Info("Starting gRPC server")

	if !RegData().Connect() {
		Shutdown(false)
	} else {
		MUp.Up()
		logging.Info("Registration Done")
		captureInterupts()
		WG.Wait()

	}
	if Data.Shutdown != nil {
		Data.Shutdown()
	}
	fmt.Println("Shutdown Complete Exiting Cleanly")
	os.Exit(0)
}

func captureInterupts() {
	logging.Debug(debug.FUNCTIONCALLS, "CALLED:[system.captureInterupts()]")
	termChan := make(chan os.Signal)
	ShutdownCommand = make(chan bool)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-termChan
		fmt.Print("\n")
		logging.Warning("SIGTERM received. Shutdown process initiated")
		Shutdown(true)
	}()

	go func() {
		<-ShutdownCommand
		logging.Warning("Shutdown Command received. Shutdown process initiated")
		Shutdown(true)
	}()
}
