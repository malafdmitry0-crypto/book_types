package main

import (
	"fmt"
)

func apply(f func() ([]int, error)) ([]int, error) { return f() }
func main() {
	r, e := apply(func() ([]int, error) { return nil, nil })
	fmt.Println(r == nil, e == nil)
}
