package main

import (
	"fmt"
)

func main() {
	var a []int
	b := append(a, []int{}...)
	fmt.Println(b == nil)
}
