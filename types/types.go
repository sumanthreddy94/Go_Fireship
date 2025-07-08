package main

import "fmt"

func main() {
	const a int32 = 32
	var c complex128 = 1 + 4i

	b := 32.5 
	x := "some string" 
	fmt.Printf("Programmer declared data types: \n %T %T \n", a, c)
	fmt.Printf("Program Infered data types: \n %T %T \n", b, x)
}
