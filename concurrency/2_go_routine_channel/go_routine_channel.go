package main

import "fmt"

// Main go Routine also communicate with its go routines using CHANNELS
// We can implement Join using this. Nice..!
// CHANNELS are means to communicate between go routines
// We can think of Channel like FIFO queue
func main() {
	myChannel := make(chan string)


	go func(){  // Anonymous routine 
		myChannel <- "data"
	}()         // Forked Off

	// This will be blocking code
	msg := <- myChannel // Rejoining the main routine using channel, Main routine waits to receive data from channel

	fmt.Println(msg)

}
