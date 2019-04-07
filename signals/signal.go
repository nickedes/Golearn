package main

import (
	"fmt"
	"time"
)

func sendSignal(ch chan struct{}) {
	time.Sleep(10 * time.Second) // sleeps for 10 seconds
	close(ch)                    // channel close will cause signal
}

func main() {
	ch := make(chan struct{})

	go sendSignal(ch)

	for {
		select {
		case <-ch:
			fmt.Println("Signal Received!")
			return
		default:
			fmt.Println("Waiting for Signal!")
		}
	}
}
