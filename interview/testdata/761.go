package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	s := a[:1:1]
	s = append(s, 9)
	s[0] = 7
	fmt.Println(a, s)
}
