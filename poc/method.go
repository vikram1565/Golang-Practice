package main

import (
	"fmt"
)

type MyName struct {
	Name string
}

func (m MyName) PrintName() string {
	return m.Name
}

func (m *MyName) SetName(name string) {
	m.Name = name
}
func main() {
	var name MyName
	(*MyName).SetName(&name, "Vikram")
	fmt.Println(MyName.PrintName(name))
}
