package main

import (
	"fmt"
	"time"
)

func main() {
	// total jobs channel
	numJobs := make(chan int, 9)
	// job output channel
	outputs := make(chan int, 9)
	// create 3 worker pool
	for w := 1; w <= 3; w++ {
		go workerpool(w, numJobs, outputs)
	}
	// send data to channel
	for j := 1; j <= 9; j++ {
		numJobs <- j
	}
	close(numJobs) // close job channel
	// receive data from output channel
	for j := 1; j <= 9; j++ {
		<-outputs
	}
}

// numJobs - receive channel and outputs - send channel
func workerpool(id int, numJobs <-chan int, outputs chan<- int) {

	for num := range numJobs {
		fmt.Println("Worker ", id, "Started Job ", num)
		time.Sleep(time.Second)
		fmt.Println("Worker ", id, "finished Job ", num)
		outputs <- num
	}
}
