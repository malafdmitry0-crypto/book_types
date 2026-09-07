package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	b := a[:1:1]
	b = append(b, 9)
	fmt.Println(a, b)
}
