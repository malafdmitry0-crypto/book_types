package main

import (
	"fmt"
)

func f() []int {
	s := make([]int, 1, 2)
	s[0] = 1
	defer func() { s = append(s, 2) }()
	return s
}
func main() {
	fmt.Println(f())
}
