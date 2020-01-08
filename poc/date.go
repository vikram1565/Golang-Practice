package main

import (
	"fmt"
	"time"
)

func main() {
	ctime := time.Now()
	fmt.Println(ctime.Format("02_01_2006"))
}
