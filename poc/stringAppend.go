package main

import (
	"fmt"
	"strings"
)

func main() {

	err := make(map[string]interface{})
	err["name"] = "name is required"
	err["phone"] = "phone is required and numeric"
	var s strings.Builder
	for k, v := range err {
		s.WriteString(k)
		s.WriteString(":")
		s.WriteString(v.(string))
		s.WriteString("\n")
	}
	fmt.Println(s.String())
}
