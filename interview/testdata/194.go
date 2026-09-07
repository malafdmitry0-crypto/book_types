package main

import (
	"fmt"
)

func f() []int {
	a := [2]int{1, 2}
	return a[:]
}
func main() {
	s := f()
	s[0] = 9
	fmt.Println(s)
}
