package main

import (
	"fmt"
)

func main() {
	var s []int
	a := [0]int(s)
	p := (*[0]int)(s)
	fmt.Println(len(a), p == nil)
}
