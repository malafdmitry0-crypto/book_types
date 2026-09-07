package main

import (
	"fmt"
)

func main() {
	s := make([]int, 2, 4)
	s[0] = 1
	s[1] = 2
	sum := 0
	for _, v := range s {
		sum += v
		s = append(s, 9)
	}
	fmt.Println(sum, len(s))
}
