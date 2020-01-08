package main

import "fmt"

func main() {

	arr := [2][2]int{{1, 2}, {3, 4}}
	fmt.Println("before ---- ", arr)
	i, j := 0, 1
	for a := range arr {
		if arr[i][a] < arr[j][a] {
			temp := arr[i][a]
			arr[i][a] = arr[j][a]
			arr[j][a] = temp
		}
	}
	fmt.Println("after ---- ", arr)
}
