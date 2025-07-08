package main

import "fmt"
const ROUTINE_SIZE int = 50
func main() {
	c := make(chan int)

	for i := range ROUTINE_SIZE {
		go cookingGoChef(i, c) // start the go routine (Thread)
	}

	i := 0
	for i < ROUTINE_SIZE {
		goChefId := <- c // receive from 
		fmt.Println("goChecf", goChefId, "finished the dish")
		i++
	}

}

// Initialize the Go Routine (thread)
func cookingGoChef(id int, c chan int) {
	fmt.Println("Go Chef", id, "started cooking")
	c <- id // send value to channel from this routine (thread)
}