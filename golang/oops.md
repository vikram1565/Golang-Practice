# **OOPs**

## *Is Go an Object Oriented Language ?*

No, Go is not object oriented.

But by using structs and interfaces we can achieve oops concepts in go.

Go has types and methods and allows an object-oriented style of programming, there is no type hierarchy.

The concept of “interface” in Go provides a different approach that we believe is easy to use and in some ways more general.

For more details please [read this article](https://medium.com/gophersland/gopher-vs-object-oriented-golang-4fa62b88c701)

There is no inheritance in go but embedding one struct into another we can achieve this.

## *Inheritance*

```go
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

```

## *Abstraction*

```go
   package main

import "fmt"

// Define a new data type "Triangle"
type Triangle struct {
    base, height float32
}

// Define a new data type "Square"
type Square struct {
    length float32
}

// Define a new data type "Rectangle"
type Rectangle struct {
    length, width float32
}

// Define a new data type "Circle"
type Circle struct {
    radius float32
}

// A method for type "Triangle"
func (t Triangle) Area() float32 {
    return 0.5 * t.base * t.height
}

// A method for type "Square"
func (l Square) Area() float32 {
    return l.length * l.length
}

// A method for type "Rectangle"
func (r Rectangle) Area() float32 {
    return r.length * r.width
}

// A method for type "Circle"
func (c Circle) Area() float32 {
    return 3.14 * (c.radius * c.radius)
}

// Define an interface as achieve abstraction
type Area interface {
    Area() float32
}

func main() {
    // Declare and assign values to varaibles
    t := Triangle{base: 15, height: 25}
    s := Square{length: 5}
    r := Rectangle{length: 5, width: 10}
    c := Circle{radius: 5}

    // Define a variable of type interface
    var a Area

    // Assign to the interface a variable of type "Triangle"
    a = t
    fmt.Println("Area of Triangle", a.Area())

    // Assign to the interface a variable of type "Square"
    a = s
    fmt.Println("Area of Square", a.Area())

    // Assign to the interface a variable of type "Rectangle"
    a = r
    fmt.Println("Area of Rectangle", a.Area())

    // Assign to the interface a variable of type "Circle"
    a = c
    fmt.Println("Area of Circle", a.Area())
}

```

## *Polymorphism*

```go
    package main

import (
    "fmt"
    "strings"
)

// Any type which defines all the methods of
// an interface is said to implicitly implement that interface.

type MyInterface interface {
    addition()
    substraction()
}

// Define a new data type additionStruct
type additionStruct struct {
    a, b int
}

// Define a new data type substractionStruct
type substractionStruct struct {
    m, n string
}

// Implemented method for additionStruct
func (add additionStruct) addition() {
    fmt.Println("First Addition is : ", add.a+add.b)
}

// Implemented method for additionStruct
func (add additionStruct) substraction() {
    fmt.Println("First Substraction is : ", add.a-add.b)
}

// Implemented method for substractionStruct
func (sub substractionStruct) addition() {
    fmt.Println("Second Addition is : ", sub.m+sub.n)
}

// Implemented method for substractionStruct
func (sub substractionStruct) substraction() {
    fmt.Println("Second Substraction is : ", strings.ToUpper(sub.m)+strings.ToUpper(sub.n))
}

func main() {
    add := additionStruct{30, 20}
    sub := substractionStruct{"hello", "world"}
    // Calling interface methods with different signature
    MyInterface.addition(add)
    MyInterface.substraction(add)
    MyInterface.addition(sub)
    MyInterface.substraction(sub)
}

```

## *How can we achieve encapsulation in Go ?*

Answer to question
