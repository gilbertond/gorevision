package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	handlerSignal()
}

// example when you receive a signal, gracefully shutdown server or commandline tool to stop processing input
func handlerSignal() {
	expectedSignals := make(chan os.Signal, 1)
	doneSignals := make(chan bool, 1)

	// register channel to receive 2 signals
	signal.Notify(expectedSignals, syscall.SIGTERM, syscall.SIGINT)

	// this routine is blocking, i.e. when it gets one signal it prints it and notifies the program that it can finish
	go func() {
		sig := <-expectedSignals
		fmt.Println()
		fmt.Println(sig.String())
		doneSignals <- true
	}()

	fmt.Println("awaiting signal...")

	<-doneSignals

	fmt.Println("exiting...")
}


