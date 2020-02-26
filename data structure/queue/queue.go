package main

import "fmt"

// queue size
var size = 5

// queue int slice
type queue []int

// isEmpty - check queue is empty or not
func (q *queue) isEmpty() bool {
	if len(*q) == 0 {
		fmt.Println("Queue is empty")
		return true
	}
	return false
}

// isFull - check queue is full or not
func (q *queue) isFull() bool {
	if len(*q) == size {
		fmt.Println("Queue is full")
		return true
	}
	return false
}

// enqueue - add node in list
func (q *queue) enqueue(elem int) {
	if q.isFull() {
		fmt.Println("queue Overflows!!!")
		return
	}
	*q = append(*q, elem)
}

// dequeue - delete node form list
func (q *queue) dequeue() {
	if q.isEmpty() {
		fmt.Println("queue Underflows")
		return
	}
	*q = (*q)[1:]
}

// display - display list nodes
func (q *queue) display() {
	for i := len(*q) - 1; i >= 0; i-- {
		fmt.Printf("%d ", (*q)[i])
	}
}
func main() {
	// queue type variable
	var q queue
	// push data in queue
	q.enqueue(10)
	q.enqueue(20)
	q.enqueue(30)
	q.enqueue(40)
	q.enqueue(50)
	fmt.Println("After enqueue: ")
	// print queue data
	q.display()
	// delete data from queue
	q.dequeue()
	fmt.Println("\nAfter dequeue: ")
	q.display()
}
