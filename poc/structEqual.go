package main

import (
	"fmt"
	"reflect"
)

type Abc struct {
	name string
	age  int
}
type Bbc struct {
	name string
	age  int
}

func main() {

	a := Abc{}
	b := Abc{}
	// b := Bbc{}
	fmt.Println(reflect.DeepEqual(a, b))
}
