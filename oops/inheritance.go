package main

import (
	"fmt"
	"strconv"
)

// Define a new data type "author"
type author struct {
	name string
	age  int
}

// Composition - embedding one struct into another
type book struct {
	bname string
	pages int
	price float64
	author
}

// A method for type "author"
func (a author) authorInfo() string {
	return "author is " + a.name + " and age is " + strconv.Itoa(a.age)
}

// A method for type "book"
func (b book) bookInfo() {
	fmt.Println("Book: ", b.bname)
	fmt.Println("Book", b.authorInfo()) // access the embedded fields
	fmt.Println("Pages:   ", b.pages)
	fmt.Println("Price: ", b.price)
}

func main() {
	// Declare and assign values to varaibles
	a := author{
		"Vikram Ingawale", 25,
	}

	b := book{
		"Go Programming Language", 500, 320.00, a,
	}

	b.bookInfo()
}
