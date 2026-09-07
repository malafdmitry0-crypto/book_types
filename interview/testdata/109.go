package main

import (
	"fmt"
)

type S struct {
	X any
	N int
}

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	a := S{[]int{}, 1}
	b := S{[]int{}, 2}
	fmt.Println(a == b)
}
