package main

import (
	"fmt"
)

func mapInts(s []int, f func(int) int) []int {
	out := make([]int, len(s))
	for i, v := range s {
		out[i] = f(v)
	}
	return out
}
func main() {
	s := []int{1, 2}
	out := mapInts(s, func(n int) int {
		s[1] = 9
		return n * 2
	})
	fmt.Println(out, s)
}
