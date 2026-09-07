package main

import (
	"fmt"
)

func main() {
	s := []int{1, 2}
	a := [2]int(s)
	s[0] = 9
	fmt.Println(a, s)
}
