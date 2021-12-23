package system

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func captureInterupt() {
	termChan := make(chan os.Signal)
	signal.Notify(termChan, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		<-termChan
		fmt.Print("\nSIGTERM received. Shutdown process initiated\n")
		Shutdown(true)
	}()
}
