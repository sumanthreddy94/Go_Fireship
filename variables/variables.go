package main

import "fmt"

//  declare single variable
var a int

// Declare a group of variables
var (
	b bool
	c float32
	d string
)

func main() {
 a = 42
 b,c = true, 3.14
 d = "string"
 fmt.Println(a,b,c,d)

//  Initialize AND assign variables
 e := 42
 f, g := true, 32.5

 
 fmt.Println(e,f,g)

}
