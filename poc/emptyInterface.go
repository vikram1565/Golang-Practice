package main

import "fmt"

type mystr string

type EmptyInterface struct {
	Name string
	Age  int
}

func print(ei interface{}) {
	fmt.Println(ei)
}

func main() {
	fmt.Println("Empty Interface")
	m := mystr("This is empty interface")
	print(m)

	e := EmptyInterface{"Vikram", 25}
	print(e)
}
