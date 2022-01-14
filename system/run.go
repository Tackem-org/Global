package system

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/logging/debug"
	pbregclient "github.com/Tackem-org/Proto/pb/regclient"
	pbremoteweb "github.com/Tackem-org/Proto/pb/remoteweb"
	"google.golang.org/grpc"
)

func Run(data SetupData) {
	logging.Setup(data.LogFile, data.VerboseLog, data.DebugLevel)
	logging.Debugf(debug.FUNCTIONCALLS, "CALLED:[system.Run(data SetupData)] {data=%+v}", data)
	logging.Infof("Starting Tackem %s System", data.BaseData.ServiceName)
	healthcheckHealthy = true
	Data = data
	MUp.StartDown()

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
