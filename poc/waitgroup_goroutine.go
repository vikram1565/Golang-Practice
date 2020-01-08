package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	str := []string{"Vikram", "Akshay", "Kunal"}
	for _, s := range str {
		wg.Add(1)
		go printName(s, &wg)
	}
	wg.Wait()
}

func printName(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("my name is", s)
}
