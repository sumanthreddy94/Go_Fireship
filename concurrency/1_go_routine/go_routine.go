package main

import (
	"fmt"
	"time"
)

func someFunc(num int) {
	fmt.Println(num)
}

func someAsyncFunc(num int) {
	fmt.Println(num)
}

func main() {
	someFunc(2)

	// Go Follows Fork Join Model
	// Go forks this someAsyncFunc, but we have to implement Join.
	go someAsyncFunc(33)
	go someAsyncFunc(44)
	go someAsyncFunc(55)

	// Adding this is a CRUDE way to assume go routines will complete in 2 sec while main routine sleeps/waits
	time.Sleep(time.Second * 2)

	fmt.Println("hi")
}

// Go will terminate program as soon as it reaches last line, 
// Go implicitly does not wait for its go routines to complete
// Explicitly u may 
// U may not see output of async funcs