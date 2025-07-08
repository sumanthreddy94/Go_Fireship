package main

import (
	"fmt"
)

//  SELECT makes go routine WAIT for one of the channels
func main() {
	myChannel := make(chan string)
	myAnotherChannel := make(chan string)

	go func(){ 
		myChannel <- "data"
	}()  
	
	go func(){ 
		myAnotherChannel <- "APPLE"
	}()

	// msgFromMyChannel := <- myChannel  // WAITS for channel
	// fmt.Println(msgFromMyChannel)

	// msgFromAnotherChannel := <- myAnotherChannel  // And then WAIT for another channel
	// fmt.Println(msgFromAnotherChannel)


	
	// SELECT waits for message from one of the channels
	// If BOTH messages come at same time, one of the msgs is selected at RANDOM
	select {
	case msgFromMyChannel := <- myChannel:
		fmt.Println(msgFromMyChannel)
	case msgFromAnotherChannel := <- myAnotherChannel:
		fmt.Println(msgFromAnotherChannel)
	}
	


}