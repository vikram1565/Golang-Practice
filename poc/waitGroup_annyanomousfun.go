package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	keys := []string{"vikram", "akshay", "kunal"}

	for _, k := range keys {
		wg.Add(1)
		go func(s string) {
			fmt.Println(s)
			defer wg.Done()
		}(k)
	}
	wg.Wait()
}
