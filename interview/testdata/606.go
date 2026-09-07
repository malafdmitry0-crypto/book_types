package main

import (
	"fmt"
)

func f() []int {
	s := []int{1, 2}
	defer func() { s[0] = 7 }()
	return s
}
func main() {
	fmt.Println(f())
}
