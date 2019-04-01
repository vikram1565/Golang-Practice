package main

import (
	"fmt"
	"strings"
)

// Any type which defines all the methods of
// an interface is said to implicitly implement that interface.

type MyInterface interface {
	addition()
	substraction()
}

// Define a new data type additionStruct
type additionStruct struct {
	a, b int
}

// Define a new data type substractionStruct
type substractionStruct struct {
	m, n string
}

// Implemented method for additionStruct
func (add additionStruct) addition() {
	fmt.Println("First Addition is : ", add.a+add.b)
}

// Implemented method for additionStruct
func (add additionStruct) substraction() {
	fmt.Println("First Substraction is : ", add.a-add.b)
}

// Implemented method for substractionStruct
func (sub substractionStruct) addition() {
	fmt.Println("Second Addition is : ", sub.m+sub.n)
}

// Implemented method for substractionStruct
func (sub substractionStruct) substraction() {
	fmt.Println("Second Substraction is : ", strings.ToUpper(sub.m)+strings.ToUpper(sub.n))
}

func main() {
	add := additionStruct{30, 20}
	sub := substractionStruct{"hello", "world"}
	// Calling interface methods with different signature
	MyInterface.addition(add)
	MyInterface.substraction(add)
	MyInterface.addition(sub)
	MyInterface.substraction(sub)
}
