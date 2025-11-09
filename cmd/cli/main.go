package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/MutsiMutsi/go-novon/core"
)

func main() {
	streamer := core.NewStreamer()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := streamer.Start(); err != nil {
			fmt.Println("Error starting streamer:", err)
		}
	}()

	fmt.Println("Streamer running. Press Ctrl+C to stop.")

	<-stop
	fmt.Println("\nStopping streamer...")
	streamer.Stop() // you need a method to cancel context and cleanup
}
