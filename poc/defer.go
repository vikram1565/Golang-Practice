package main

import (
	"fmt"
)

func main() {
	//demo()
	//test()
	// funny()
	//i := tunny()
	//fmt.Println("tunny: ", i)
	j := dummy()
	fmt.Println("dummy: ", j)

}

// defer with argument can evaluate at the time of call
func demo() (i int) {
	i = 0
	defer fmt.Println("demo: ", i)
	i++
	return i
}

// defer with function can change the value at the time of return
func test() (i int) {
	i = 0
	defer func() { fmt.Println("test: ", i) }()
	i++
	return i
}

// passing argument to defer function so calculated at the time of call and i is 1
func funny() (i int) {
	i = 0
	defer func(i int) { fmt.Println("funny: ", i) }(i)
	i++
	return i
}

// return  type so i = 1
func tunny() int {
	i := 0
	defer func() {
		i++
	}()
	i++
	return i
}

// return name of variable with type and variable is return then defer change the value of variable so i = 2
func dummy() (i int) {
	i = 0
	defer func() {
		i++
	}()
	i++
	return i
}
