package main

import (
	"fmt"
)

func main() {
	a := make([]int, 0)
	b := a[:0]
	fmt.Println(a == nil, b == nil)
}
