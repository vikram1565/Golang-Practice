package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("Starting")
	go func() {
		defer wg.Done()
		// time.Sleep(1 * time.Microsecond)
		for i := 'a'; i < 'a'+26; i++ {
			fmt.Printf("%c ", i)
		}
	}()
	go func() {
		defer wg.Done()
		for j := 1; j < 27; j++ {
			fmt.Printf("%d ", j)
		}
	}()
	fmt.Println("Waiting to finish")
	wg.Wait()
	fmt.Println("\nTerminating")
}
