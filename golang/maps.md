# **Maps in golang**

## *What is hashmap?*

## *Is map in go a hashmap?*

## *How hashing function works for maps in go?*

## *Does a map preserve order of keys in go?*

## *What will you do if you want get entries in an order they are stored?*

## *What will be the output?*

```go
package main

var m = make(map[string]int, 5)

func main() {
    m["1"] = 1
    m["2"] = 2
    m["3"] = 3
    m["4"] = 4
    m["5"] = 5

    for key, val := range m {
        fmt.Println("Key - ", key, "Value - ", val)
    }
}

```

Answer - Keys and values will be printed in random order, each time we run the program.

## *What will happen if I add entries beyond specified capacity to the map?*
