# **Arrays and Slices**

## *Arrays*

- Array is collection of elements/items with same type.

- Array can be declared by using below syntax 
    [n]T
    - n is the total number of elements
    - T is the type of each elements

- Array simple example
    var a [3]int
 - this is array of integer with length 3
 - all elements in an array are automatically assigned to 0 value
 - so this array is [0 0 0]
  
- Array index is starts from 0 and ends with length-1

##*What is the output of below code*
```go
package main

import (  
    "fmt"
)

func main() {  
    arr := [3]int{10} 
    fmt.Println(arr)
}
```
- here , array declare with 3 length but provided one value.
so the array remaining elements is set to 0
and the final result is [10 0 0]

- You can even ignore the length of the array in the declaration and replace it with ... 
the compiler find the length.
arr := [...]int{10,20,30,40,50} 

- array can't be resized . [3]int and [5]int are distinct types.

##* What is output of below code*
```go
package main
import "fmt"
func main() {  
    a := [3]int{10,20,30}
    var b [5]int
    b = a
    fmt.Println(b)
}
```
- compiler error : cannot use a (type [3]int) as type [5]int in assignment.

- Array in go are value types. when you assign array to variable then copy of original array is assigned to
  the new variable. If changes are made to the new variable, it will not be reflected in the original array.

##* What is output of below code*
```go
package main
import "fmt"
func main() {  
    a := [...]int{10,20,30}
    b := a
    b[0] = "40"
    fmt.Println(a)
    fmt.Println(b) 
} 
```
- Similarly when arrays are passed to functions as parameters, they are passed by value and the original array in unchanged.
- length of array can be find by using inbuilt function len()
  
* Array is fix sized and we can't resized array length. to overcome this drawback of array,  Slices comes in picture.



## *Slices*

- A Slice is a segment of an array. Slices build on arrays and provide more power, flexibility, and convenience compared to arrays.
- Just like arrays, Slices are indexable and have a length. But unlike arrays, they can be resized.
- Internally, A Slice is just a reference to an underlying array.
- A slice of type T is declared using []T.
- var arr []int , here we can declare a slice of int type. here we do not specify any size in the brackets [].

## *Slice Example*
```go
package main
import "fmt"
func main() {
	// Creating a slice using a slice literal
	var s = []int{10,20,30,40,50}
	fmt.Println("s = ", s)
}
```
- a slice is a segment of an array, we can create a slice from an array.
- to create a slice from an array a, we specify two indices low (lower bound) and high (upper bound) separated by a      colon - a[low:high]
- The resulting slice includes all the elements starting from index low to high, but excluding the element at index      high.

## *What is the output of below code*
```go
package main
import "fmt"

func main() {
	var a = [5]int{10,20,30,40,50}
	var s []string = a[1:4]
	fmt.Println("Array a = ", a)
	fmt.Println("Slice s = ", s)
}
Answer - ?
```

- The low and high indices in the slice expression are optional. The default value for low is 0, and high is the length of the slice.
  
## *What is the output of below code*
```go
package main
import "fmt"

func main() {
	var a = [5]int{10,20,30,40,50}
	s1 := a[1:4]
    s2 := a[:3]
	s3 := a[2:]
	s4 := a[:]
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
}
Answer - ?
```
- Slices are reference types. They refer to an underlying array. Modifying the elements of a slice will modify the corresponding elements in the referenced array. Other slices that refer the same array will also see those modifications.

##*What is the output of below code*
```go
package main
import "fmt"

func main() {
	var a = [5]int{10,20,30,40,50}
	s1 := a[1:]
    s1[0] = 100
	fmt.Println(a)
	fmt.Println(s1)
}
Answer - ?
```

#**Length and Capacity of a Slice*

- The length of the slice is the number of elements in the slice.
- The capacity is the number of elements in the underlying array starting from the first element in the slice.
- example
```go
package main
import "fmt"
func main() {
	a := [6]int{10, 20, 30, 40, 50, 60}
	s := a[1:4]
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))
}
  Answer : s = [20 30 40], len = 3, cap = 5
  ```
- golang provides the library function make() to create slice. below are the syntax
   func make([]T, len, cap) []T
    - he make function takes a type, a length, and an optional capacity. It allocates an underlying array with size equal to the given capacity, and returns a slice that refers to that array.

- The zero value of a slice is nil. A nil slice doesn’t have any underlying array, and has a length and capacity of 0.
  

* Slice Functions
  
  1) copy()
        - The copy() function copies elements from one slice to another.
        - Its signature looks like this - func copy(dst, src []T) int
        - It takes two slices - a destination slice, and a source slice. It then copies elements from the source to      the destination and returns the number of elements that are copied.
        - The number of elements copied will be the minimum of len(src) and len(dst).
  
  2) append())
        - The append() function appends new elements at the end of a given slice.
        - Its signature looks like this - func append(s []T, x ...T) []T
        - It takes a slice and a variable number of arguments x …T. It then returns a new slice containing all the       elements from the given slice as well as the new elements.
        - If the given slice doesn’t have sufficient capacity to append new elements then a new underlying array is      allocated with bigger capacity. All the elements from the underlying array of the existing slice are copied    to this new array, and then the new elements are appended.
        - if the slice has enough capacity to accommodate new elements, then the append() function re-uses its           underlying array and appends new elements to the same array.
        - When you append values to a nil slice, it allocates a new slice and returns the reference of the new slice.
        - You can directly append one slice to another using the ... operator. (Ex: append(slice1, slice2...))
  
#**Slice append eample*
```go
package main
import "fmt"
func main() {
	slice1 := []int{10,20,30}
	slice2 := append(slice1, 40,50,60)
	slice1[0] = 90
	fmt.Println("slice1 = ", slice1)
	fmt.Println("slice2 = ", slice2)
}

Answer - 
slice1 =  [90 20 30]
slice2 =  [10,20,30 40 50 60]
```