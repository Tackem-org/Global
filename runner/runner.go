package runner

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/Tackem-org/Global/checkerServer"
	"github.com/Tackem-org/Global/helpers"
	"github.com/Tackem-org/Global/logging"
	"github.com/Tackem-org/Global/regClientServer"
	"github.com/Tackem-org/Global/remoteWebSystem"
	"github.com/Tackem-org/Global/system"
	pbchecker "github.com/Tackem-org/Proto/pb/checker"
	pbregclient "github.com/Tackem-org/Proto/pb/regclient"
	pbremoteweb "github.com/Tackem-org/Proto/pb/remoteweb"
)

func Run(data system.SetupData) {
	system.Data = data
	fmt.Printf("Starting Tackem %s System\n", system.Data.BaseData.ServiceName)

	if !helpers.InDockerCheck() {
		return
	}
	fmt.Printf("Verbose %t\n", system.Data.VerboseLog)
	logging.Setup(system.Data.LogFile, system.Data.VerboseLog)
	logging.Info("Logger Started")

	logging.Info("Setup Registration Data")
	system.RegData().Setup(system.Data.BaseData)

	system.WG = &sync.WaitGroup{}

	logging.Info(fmt.Sprintf("Setup %s System", data.BaseData.ServiceName))
	if system.Data.MainSystem != nil {
		system.Data.MainSystem()
	}

	logging.Info("Setup Web Service")
	if system.Data.WebSystems != nil {
		system.Data.WebSystems()
	}

	logging.Info("Setup GPRC Service")
	pbremoteweb.RegisterRemoteWebServer(system.GRPCServer(), remoteWebSystem.NewServer())
	pbchecker.RegisterCheckerServer(system.GRPCServer(), checkerServer.NewServer())
	pbregclient.RegisterRegClientServer(system.GRPCServer(), regClientServer.NewServer())
	system.Data.GPRCSystems(system.GRPCServer())

	system.WG.Add(1)
	go func() {
		port := fmt.Sprint(system.RegData().GetPort())
		listen, err := net.Listen("tcp", ":"+port)
		if err != nil {
			logging.Error("gPRC could not listen on port " + port)
			logging.Fatal(err)
		}
		if err := system.GRPCServer().Serve(listen); err != nil {
			logging.Fatal(err)
		}
	}()
	logging.Info("Starting gRPC server")

	if !system.RegData().Connect() {
		system.Shutdown(false)
	} else {
		system.MasterUp = true
		logging.Info("Registration Done")
		captureInterupt()
		system.WG.Wait()

	}
	fmt.Println("Shutdown Complete Exiting Cleanly")
	os.Exit(0)
}

func captureInterupt() {
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-termChan
		fmt.Print("\nSIGTERM received. Shutdown process initiated\n")
		system.Shutdown(true)
	}()
}
