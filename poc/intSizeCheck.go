package main

import (
	"fmt"
	"unsafe"
)

func main() {
	fmt.Println("Hello, playground")
	var a int
	var x int8
	var y int16
	var b int32
	var c int64
	fmt.Println("A - ", unsafe.Sizeof(a))
	fmt.Println("X - ", unsafe.Sizeof(x))
	fmt.Println("Y - ", unsafe.Sizeof(y))
	fmt.Println("B - ", unsafe.Sizeof(b))
	fmt.Println("C - ", unsafe.Sizeof(c))
	
}