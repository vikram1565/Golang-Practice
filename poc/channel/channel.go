package main

import "fmt"

// Please uncomment case wise code and then check code
func main() {
	// CASE - I : we can't send and receive on nil channel
	var ch chan int // nil channel
	go func() {
		fmt.Print(<-ch)
	}()
	ch <- 1
	fmt.Println("hello")

	// CASE - II : we can send and receive on open channel
	// ch := make(chan int) // open channel , unbuffered channel
	// go func() {
	// 	fmt.Println(<-ch)
	// }()
	// ch <- 1
	// fmt.Println("hello")

	// CASE - III : We can't send on close channel but we can receive on close channel
	// ch := make(chan int, 3) // buffered channel
	// go func() {
	// 	fmt.Println(<-ch)
	// }()
	// ch <- 1
	// ch <- 2
	// fmt.Println("hello")
	// close(ch)
	// ch <- 3 // panic: send on closed channel
}
