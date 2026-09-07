package main

import (
	"fmt"
)

type S struct{ A []int }

func main() {
	a := S{A: []int{1, 2}}
	b := a
	b.A[0] = 9
	b.A = b.A[:1]
	fmt.Println(a.A, b.A)
}
