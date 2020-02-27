// find the length of the longest substring without repeating characters.
// aaaaaabbbbbbbbbsssssssdddddddeeeeee
// absde that is 5

package main

import (
	"bytes"
	"fmt"
)

func main() {
	str := "aaaaaabbbbbbbbbsssssssdddddddeeeeeeff"
	length, key := findLength(str)
	fmt.Println("length of the longest substring without repeating characters =  ", length, " and substring is ", key)
}

func findLength(str string) (int, string) {
	m := make(map[rune]int)
	var s bytes.Buffer
	for i, c := range str {
		if _, ok := m[c]; !ok {
			s.WriteRune(c)
		}
		m[c] = i
	}
	return len(m), s.String()
}
