package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func handleSignal(signal os.Signal) {
	fmt.Println("handleSignal() caught:", signal)
}

func main() {
	sigs := make(chan os.Signal, 1)                    // define a channel named sigs that helps you pass data around
	signal.Notify(sigs, os.Interrupt, syscall.SIGINFO) // to state the signals that interest you
	go func() {                                        // anonymous function that runs as a goroutine
		for {
			sig := <-sigs
			switch sig {
			case os.Interrupt:
				fmt.Println("Caught:", sig)
			case syscall.SIGINFO:
				handleSignal(sig)
			}
		}
	}()

	for {
		fmt.Print(".")
		time.Sleep(20 * time.Second)
	}
}
