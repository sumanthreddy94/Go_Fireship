package main

import "fmt"

func main() {
	var address *int // Declare int pointer

	number := 42 // int mem allocation

	address = &number // address stores memory address of 'number' =>  0x1400000e090

	value := *address // dereferencing the value => getting value from address => 42

	fmt.Printf("address: %v\n", address) // address:  0x1400000e090
	fmt.Printf("value: %v\n", value) 

}
