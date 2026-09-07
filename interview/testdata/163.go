package main

import (
	"fmt"
)

func main() {
	a := make([]int, 1, 3)
	b := append(a, 1)
	c := append(a, 2)
	fmt.Println(b, c)
}
