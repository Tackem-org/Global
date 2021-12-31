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
	pbregclient "github.com/Tackem-org/Proto/pb/regclient"
	pbremoteweb "github.com/Tackem-org/Proto/pb/remoteweb"
	"google.golang.org/grpc"
)

func Run(data SetupData) {
	Data = data
	healthcheckHealthy = true
	fmt.Printf("Starting Tackem %s System\n", Data.BaseData.ServiceName)
	MUp.StartDown()
	if !helpers.InDockerCheck() {
		return
	}
	fmt.Printf("Verbose %t\n", Data.VerboseLog)
	logging.Setup(Data.LogFile, Data.VerboseLog)
	logging.Info("Logger Started")

	logging.Info("Setup Registration Data")
	RegData().Setup(Data.BaseData)

	WG = &sync.WaitGroup{}

	logging.Info(fmt.Sprintf("Setup %s System", data.BaseData.ServiceName))
	if Data.MainSystem != nil {
		Data.MainSystem()
	}

	logging.Info("Setup Web Service")
	if Data.WebSystems != nil {
		Data.WebSystems()
	}

	logging.Info("Setup GPRC Service")
	grpcServer = grpc.NewServer()

	pbregclient.RegisterRegClientServer(grpcServer, NewRegClientServer())
	if data.BaseData.WebAccess {
		pbremoteweb.RegisterRemoteWebServer(grpcServer, NewRemoteWebServer())
	}
	Data.GPRCSystems(grpcServer)

	WG.Add(1)
	go func() {
		port := fmt.Sprint(RegData().GetPort())
		listen, err := net.Listen("tcp", ":"+port)
		if err != nil {
			logging.Error("gPRC could not listen on port " + port)
			logging.Fatal(err)
		}
		if err := grpcServer.Serve(listen); err != nil {
			logging.Fatal(err)
		}
	}()
	logging.Info("Starting gRPC server")

	if !RegData().Connect() {
		Shutdown(false, false)
	} else {
		MUp.Up()
		logging.Info("Registration Done")
		captureInterupts()
		WG.Wait()

	}
	fmt.Println("Shutdown Complete Exiting Cleanly")
	os.Exit(0)
}

func captureInterupts() {
	termChan := make(chan os.Signal)
	ShutdownCommand = make(chan bool)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-termChan
		fmt.Print("\n")
		logging.Warning("SIGTERM received. Shutdown process initiated")
		Shutdown(true, false)
	}()

	go func() {
		<-ShutdownCommand
		logging.Warning("Shutdown Command received. Shutdown process initiated")
		Shutdown(true, false)
	}()
}