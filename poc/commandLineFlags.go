package main

import (
	"flag"
	"fmt"
)

func main() {
	var mystr string
	flag.StringVar(&mystr, "flagName", "Vikram", "This is your name")
	flag.Parse()
	fmt.Println("Hi " + mystr)
}
