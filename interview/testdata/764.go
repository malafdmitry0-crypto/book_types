package main

import (
	"fmt"
)

type S struct{ A [2]int }

func main() {
	a := S{A: [2]int{1, 2}}
	b := a
	s := b.A[:]
	s[0] = 9
	fmt.Println(a.A, b.A)
}
