# **Functions and methods**

## Important links

1. [First class functions/GolangBot](https://golangbot.com/first-class-functions/)
2. [Go tour](https://tour.golang.org/methods/1)

## *What are first class functions*

Conditions for first class functions

1. Can be assigned to variable
2. Can be passed as argument to other functions
3. Can be returned from functions.

Go supports first class functions.

## *What are higher order functions*

Conditions for higher order functions

1. Accepts a function as an argument.

    ```go
    func simple(a func(a, b int) int) {  
        fmt.Println(a(60, 7))
    }
    ```

2. Returns a function as result.

    ```go
    func simple() func(a, b int) int {  
        f := func(a, b int) int {
            return a + b
        }
        return f
    }
    ```

## *What are function closures*

A function which accesses a variable outside the function other than arguments.

```go
package main

import (  
    "fmt"
)

func appendStr() func(string) string {  
    t := "Hello"
    c := func(b string) string {
        t = t + " " + b
        return t
    }
    return c
}

func main() {  
    a := appendStr()
    b := appendStr()
    fmt.Println(a("World"))
    fmt.Println(b("Everyone"))

    fmt.Println(a("Gopher"))
    fmt.Println(b("!"))
}

// Hello World  
// Hello Everyone  
// Hello World Gopher  
// Hello Everyone !
```

## *What are anonymous functions*

A function without any name is called as an anonymous function.

```go
func (i int) (int) {
    return i * i
}(5)
```

## *What is the difference between function and method. Tell some use cases to use them effectively*

Methods are also the functions with a reciver argument. The receiver can be a value or pointer.
Use of methods is better in case of interface implementation or when we want to manipulate the struct literals and in either case, functions are better.

## *Can we have two different functions with same name in a package*

Yes, we can have two different functions with same name, rather, two main functions with different build tags.

## *Explain named/naked return values*

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}
```

Here x, y are named return values, can be used as normal variables in function. Suitable for short functions.

## *What is pointer indirection in methods*

In go, a methods with a pointer reciver can be called on a value.

```go
package main

type Num int

func (n *Num) call() {
    fmt.Println(n)
}

func main() {
    n := Num{} // not a pointer

    n.call() // (*n).call(), implicit indirection
}

```
