# **DataTypes**

## *What are aliases in golang*

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

## *What are zero values*

Zero values are the default values assigned while declaration.

0 for numeric types,
false for the boolean type, and
"" (the empty string) for strings.

## *What is Type conversion*

The expression T(v) converts the value v to the type T. But v must be convertible to T.
Ex. We can not convert a string to integer i.e int("45") is not possible but string(45) is possible because it converts to respective ASCII value.

## *What is Type inference*

when performing

```go
var i int
j := i
```

type of j is inferred from type of i implicitly.

## *Which data types can be used as constants in golang*

Constants can be character, string, boolean, or numeric values.

## *Comment on Numeric constants*

Numeric constants are high-precision values.
An untyped constant takes the type needed by its context.

ex.

```go
package main

import "fmt"

const (
    // Create a huge number by shifting a 1 bit left 100 places.
    // In other words, the binary number that is 1 followed by 100 zeroes.
    Big = 1 << 100
    // Shift it right again 99 places, so we end up with 1<<1, or 2.
    Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
    return x * 0.1
}

func main() {
    fmt.Println(needInt(Small))
    fmt.Println(needFloat(Small))
    fmt.Println(needFloat(Big))
}
```

In above code, constant 'Small' can be passed as int as well as float64, depending on context.
