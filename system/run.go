package system

import (
	"fmt"
	"net"
	"os"
	"sync"

	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	pbchecker "github.com/Tackem-org/Proto/pb/checker"
	pbregclient "github.com/Tackem-org/Proto/pb/regclient"
	pbremoteweb "github.com/Tackem-org/Proto/pb/remoteweb"
	"google.golang.org/grpc"
)

func Run(data SetupData) {
	Data = data
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

	pbremoteweb.RegisterRemoteWebServer(grpcServer, NewRemoteWebServer())
	pbchecker.RegisterCheckerServer(grpcServer, NewCheckerServer())
	pbregclient.RegisterRegClientServer(grpcServer, NewRegClientServer())
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
		captureInterupt()
		WG.Wait()

	}
	fmt.Println("Shutdown Complete Exiting Cleanly")
	os.Exit(0)
}
