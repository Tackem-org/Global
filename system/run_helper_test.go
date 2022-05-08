package system

import pbr "github.com/Tackem-org/Global/pb/registration"

var (
	StartupFunc      = startupFunc
	MainLoop         = mainLoopFunc
	ServerServe      = serverServeFunc
	Shutdown         = shutdownFunc
	Connect          = connectFunc
	ConnectConnector = connectConnectorFunc
)

func ResetFunctions() {
	Run = run
	startup = startupFunc
	mainLoop = mainLoopFunc
	shutdown = shutdownFunc
	connect = connectFunc
	connectConnector = connectConnectorFunc
	serverServe = serverServeFunc
}

func SetupForRun(startupF func() bool, mainLoopF func(), shutdownF func(registered bool)) {
	startup = startupF
	mainLoop = mainLoopF
	shutdown = shutdownF
	Version = "v1.0.0"
}

func SetupForStartup(connectF func(request *pbr.RegisterRequest) bool, serverServeF func()) {
	connect = connectF
	serverServe = serverServeF
}

func SetupForConnect(connectConnectorF func(request *pbr.RegisterRequest) bool) {
	connectConnector = connectConnectorF
}
