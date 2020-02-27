/*
Problem -
given arr = [1, 2, 3, 4,5], sum = 6,
so -
1) 2 + 4 = 6
2) 1 + 5 = 6
print  [1 3] [0 4] arr index.
*/

package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	sum := 6

	findIndex(arr, sum)
}

func findIndex(arr []int, sum int) {
	m := make(map[int]int, len(arr))

	for i, n := range arr {
		if j, ok := m[sum-n]; ok {
			fmt.Print([]int{j, i})
		}
		m[n] = i
	}
	return
}
