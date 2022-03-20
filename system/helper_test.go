package system

func SetCalls(startup func() bool, mainLoop func(), shutdown func(bool)) {
	callStartup = startup
	callMainLoop = mainLoop
	callShutdown = shutdown
}
