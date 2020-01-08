package main

import (
	"fmt"
	"log"
)

func main() {
	var i int
	fmt.Print("enter number : ")
	_, err := fmt.Scan(&i)
	if err != nil {
		log.Fatal("Fail to get input")
	}

	str := make([]string, i)

	for j := 0; j < i; j++ {
		fmt.Println("enter name : ")
		fmt.Scan(&str[j])
	}
	fmt.Println("String Array Using FMT : ", str)
}
