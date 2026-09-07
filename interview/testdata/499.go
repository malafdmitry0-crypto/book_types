package main

import (
	"fmt"
)

type Pair[A, B any] struct {
	A A
	B B
}

func Zip[A, B any](a []A, b []B) (out []Pair[A, B]) {
	for i := 0; i < len(a) && i < len(b); i++ {
		out = append(out, Pair[A, B]{a[i], b[i]})
	}
	return
}
func main() {
	fmt.Println(Zip([]int{1, 2}, []string{"a"}))
}
