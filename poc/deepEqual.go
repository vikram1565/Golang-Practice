package main

import (
	"fmt"
	"reflect"

	"gopkg.in/mgo.v2/bson"
)

func main() {

	a := []byte{}
	b := []byte(nil)
	fmt.Println(reflect.TypeOf(a)) // rune - int32 and byte uint8
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.DeepEqual(a, b))
	aa := bson.ObjectIdHex("5d69013decda69f624553d8c")
	fmt.Println(aa.Hex())
}
