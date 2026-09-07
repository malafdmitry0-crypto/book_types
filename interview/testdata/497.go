package main

import (
	"fmt"
)

type Pred func(int) bool

func Any[E any](s []E, p func(E) bool) bool {
	for _, v := range s {
		if p(v) {
			return true
		}
	}
	return false
}
func main() {
	var f Pred = func(n int) bool { return n > 0 }
	fmt.Println(Any([]int{-1, 2}, f))
}
