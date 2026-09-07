package main

import (
	"fmt"
)

type N int

func (n N) Add(m N) *N {
	r := n + m
	return &r
}
func Sum[E interface{ Add(E) E }](s []E, a E) E { return a }
func main() {
	fmt.Println(Sum([]N{1}, N(0)))
}
