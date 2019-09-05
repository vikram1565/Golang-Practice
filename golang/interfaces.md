# Interfaces
An interface is a collection of method signatures that an Object can implement. Hence interface defines the behavior of the object.

The primary job of an interface is to provide only method signatures consisting of the method name, input arguments and return types. 

In Go, you do not explicitly mention if a type implements an interface. If a type implements a method of signature that is defined in an interface, then that type is said to implement that interface.

interface declaration is as follows -
type Shape interface {
    Area() float64
}

The interface has two types. A static type of the interface is interface itself. 
A variable of an interface type can hold a value of the Type that implements the interface. The value of that Type becomes dynamic value of the interface and that type becomes the dynamic type of the interface.

## *Example 1*

```go
    package main

    import "fmt"

    type Shape interface {
	    Area() float64
    }

    func main() {
	    var s Shape
	    fmt.Println("value of s is", s)
	    fmt.Printf("type of s is %T\n", s)
    }

```

the dynamic type of interface also called as concrete type as when we access type of interface, it returns type of it’s underlying dynamic value and it’s static type remains hidden.

## *Example 2*

```go
    package main

    import "fmt"

    type Shape interface {
	    Area() float64
    }
    type Rectangle struct {
        Hight float64
        Width float64
    }
    func (r Rectangle) Area() float64 {
        return r.Hight * r.Width
    }
    func main() {
        var s Shape
        s := Rectangle{5,6}
	    fmt.Println("value of s is", s)
	    fmt.Printf("type of s is %T\n", s)
        fmt.Println("area : ", s.Area())
}

```

you need to implement all methods declared by the interface.

When an interface has zero methods, it is called an empty interface. This is represented by interface{}. Since empty interface has zero methods, all types implement this interface.

A nil interface has a nil type.

