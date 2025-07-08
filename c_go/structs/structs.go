package main

import "fmt"

type myStack struct {
	index int
	data [5]int
}

func (s *myStack) isEmpty() bool {
	return s.index == 0
}

func (s *myStack) push(k int) {

	if s.index >= len(s.data) {
		fmt.Println("Stack overflow!")
		return
	}
	s.data[s.index] = k
	s.index++
}
// Go doesn’t pass values by reference
// Without * we pass copy of stack
func (s *myStack) pop() int {
	s.index--
	return s.data[s.index]
}

func main() {
	s := new(myStack)
	s.push(3)
	s.push(7)
	s.push(5)
	s.push(2)
	s.push(6)

	for !s.isEmpty() {
		fmt.Printf("popped: %d\n", s.pop())
	}
}