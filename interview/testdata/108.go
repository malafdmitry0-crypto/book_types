package main

import (
	"fmt"
)

type S struct {
	N int
	X any
}

func main() {
	a := S{1, []int{}}
	b := S{2, []int{}}
	fmt.Println(a == b)
}
