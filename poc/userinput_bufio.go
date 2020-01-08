package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter number : ")
	scan.Scan()
	t := scan.Text()
	n, _ := strconv.Atoi(t)
	str := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Print("Enter Name : ")
		scan.Scan()
		str[i] = scan.Text()
	}

	fmt.Print("String Array Using Bufio : ", str)
}
