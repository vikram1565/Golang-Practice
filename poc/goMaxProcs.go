// Learn below concepts -
// 1. Anonymous goroutine
// 2. Use of GOMAXPROCS
// 3. Availability of Cores to run main goroutine

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	j := 0
	n := runtime.GOMAXPROCS(0) // runtime.GOMAXPROCS(0) - 1 // this syntax is used to run main goroutine
	// fmt.Println(runtime.GOMAXPROCS(0)) // print available logical cores/ processors
	for i := 0; i < n; i++ {
		go func() {
			for { // infinite loop
				j++
			}
			// j++ // if anonymous function exits then anonymous goroutines also exit.
		}()
	}
	time.Sleep(time.Second)
	fmt.Println(j)
}
