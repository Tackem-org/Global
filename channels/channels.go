package channels

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
)

var (
	Root      RootCommands
	setupOnce sync.Once
)

type RootCommands struct {
	TermChan chan os.Signal
	Shutdown chan bool
}

func Setup() {
	setupOnce.Do(func() {
		Root.Shutdown = make(chan bool, 1)
		Root.TermChan = make(chan os.Signal, 1)
		signal.Notify(Root.TermChan, syscall.SIGTERM, syscall.SIGINT)
	})
}
