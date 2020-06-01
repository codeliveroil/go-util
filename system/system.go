package system

import (
	"os"
	"os/signal"
)

// HandleInterrupt traps an interrupt to the
// program (usually Ctrl+C) and calls the
// callback function to perform cleanup.
func HandleInterrupt(callback func()) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	go func() {
		<-interrupt
		callback()
	}()
}
