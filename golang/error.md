# **error**

error is a built-in type in Go and its zero value is nil. 

An idiomatic way to handle an error is to return it as the last return value of a function call and check for the nil condition.

val, err := myFunction( args... );
if err != nil {
  // handle error
} else {
  // success
}

the error is a built-in type but in reality, it is an interface type made available globally and it implements Error method which returns a string error message.

type error interface {
  Error() string
}

```go
package main

import "fmt"

// create a struct
type MyError struct{}

// struct implements `Error` method
func (myErr *MyError) Error() string {
	return "Something unexpected happend!"
}

func main() {

	// create error
	myErr := &MyError{}

	// print error message
	fmt.Println(myErr)
}
```

Go provides the built-in errors package which exports the New function. This function expects an error message and returns an error.

```go
package main

import "fmt"
import "errors"

func main() {

	// create error
	myErr := errors.New("Something unexpected happend!")

	// print error message
	fmt.Println(myErr)
}
```