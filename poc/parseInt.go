package main

import (
	"fmt"
	"strconv"
)

func main() {
	// i, err := strconv.ParseInt("123456", 10, 16)
	i, err := strconv.ParseInt("12345", 10, 16) // int16 range is 32767
	if err != nil {
		fmt.Println("Err - ", err)
	} else {
		fmt.Println("I - ", i)
	}
}
