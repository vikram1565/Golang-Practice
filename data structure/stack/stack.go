package main

import "fmt"

// stack size
var size = 5

// stack int slice
type stack []int

// isEmpty - check stack is empty or not
func (s *stack) isEmpty() bool {
	if len(*s) == 0 {
		fmt.Println("Stack is empty")
		return true
	}
	return false
}

// isFull - check stack is full or not
func (s *stack) isFull() bool {
	if len(*s) == size {
		fmt.Println("Stack is full")
		return true
	}
	return false
}

// push - add element in stack
func (s *stack) push(elem int) {
	if s.isFull() {
		fmt.Println("Stack Overflows!!!")
		return
	}
	*s = append(*s, elem)
}

// pop - remove the stack element
func (s *stack) pop() {
	if s.isEmpty() {
		fmt.Println("Stack Underflows")
		return
	}
	*s = (*s)[:len(*s)-1]
}

// display - print all stack elements
func (s *stack) display() {
	for i := len(*s) - 1; i >= 0; i-- {
		fmt.Printf("%d ", (*s)[i])
	}
}
func main() {
	// stack variable
	var s stack
	// push elements in stack
	s.push(10)
	s.push(20)
	s.push(30)
	s.push(40)
	s.push(50)
	fmt.Println("After Stack Push: ")
	// print elements
	s.display()
	// pop stack
	s.pop()
	fmt.Println("\nAfter Stack Pop: ")
	s.display()
}
