package main

import (
	"fmt"
)

func filter(s []int, p func(int) bool) (out []int) {
	for _, v := range s {
		if p(v) {
			out = append(out, v)
		}
	}
	return
}
func main() {
	n := 0
	p := func(int) bool {
		n++
		return n%2 == 1
	}
	fmt.Println(filter([]int{1, 2, 3}, p), filter([]int{1, 2, 3}, p))
}
