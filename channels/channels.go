package channels

import (
	"os"
	"os/signal"
	"syscall"
)

var (
	Root RootCommands
)

type RootCommands struct {
	TermChan chan os.Signal
	Shutdown chan bool
}

func Setup() {
	Root.Shutdown = make(chan bool, 1)
	Root.TermChan = make(chan os.Signal, 1)
	signal.Notify(Root.TermChan, syscall.SIGTERM, syscall.SIGINT)
}
