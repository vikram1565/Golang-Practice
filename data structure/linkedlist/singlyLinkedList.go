package main

import "fmt"

// list struct
type node struct {
	number int
	next   *node
}

// node struct
type singlylist struct {
	head *node
}

// addNode - add number in singly list
func (list *singlylist) addNode(n int) {
	l := &node{
		number: n,
	}
	// empty list
	if list.head == nil {
		list.head = l
		return
	}
	// add at the end of list
	current := list.head
	if current.next != nil {
		current = current.next
	}
	current.next = l
	return
}

// displayList - display given list
func (list *singlylist) displayList() {
	current := list.head
	// empty list
	if current == nil {
		fmt.Println("List is empty")
		return
	}
	// single node
	fmt.Printf("%d ", current.number)
	// 1+ nodes
	for current.next != nil {
		current = current.next
		fmt.Printf("%d ", current.number)
	}
	return
}

// insertAfter - insert node in list after given node
func (list *singlylist) insertAfter(after, n int) {

	// check list is empty or not
	if list == nil {
		fmt.Println("List is empty")
		return
	}
	l := &node{
		number: n,
	}
	current := list.head
	// find and insert node
	for current != nil {
		if current.number == after {
			l.next = current.next
			current.next = l
			return
		}
		current = current.next
	}
	return
}

func (list *singlylist) deleteNode(n int) {
	// check list is empty or not
	if list == nil {
		fmt.Println("List is empty")
		return
	}
	current := list.head
	var previous *node
	for current != nil {
		if current.number == n {
			previous = current.next
			current.number = previous.number
			current.next = previous.next
			return
		}
		current = current.next
	}
}

// main - main function
func main() {
	// linked list variable
	sl := &singlylist{}
	// add node in list
	sl.addNode(10)
	sl.addNode(20)
	fmt.Println("List is : ")
	// print list nodes
	sl.displayList()
	// insert after
	sl.insertAfter(10, 30)
	fmt.Println("\nAfter Insert At : ")
	sl.displayList()
	// delete node
	sl.deleteNode(30)
	fmt.Println("\nAfter Delete : ")
	sl.displayList()
}
