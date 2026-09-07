package main

import (
	"fmt"
	"slices"
)

func main() {
	s := make([]int, 2, 3)
	s[0] = 1
	s[1] = 2
	seq := slices.Values(s)
	s[0] = 9
	s = append(s, 3)
	fmt.Println(slices.Collect(seq))
}
