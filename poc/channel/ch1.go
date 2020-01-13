package main

import (
	"fmt"
)

func main() {

	ch := make(chan string, 1)
	complete := make(chan struct{})
	go getNames(complete)

	for {
		select {
		case p := <-ch:
			fmt.Println("Blocked User : ", p)
			continue

		case <-complete:
			return
		}
	}
}

func getNames(complete chan struct{}) {

	defer func() {
		close(complete)
	}()

	s := make([]string, 5)
	fmt.Println("Enter 5 user names")
	for i := 0; i < 5; i++ {
		fmt.Print("Enter Name : ")
		fmt.Scan(&s[i])
		if s[i] == "vikram" {
			fmt.Println("Blocked User : ", s[i])
			return
		}
	}
	fmt.Println("You entered valid users list : ", s)
}
