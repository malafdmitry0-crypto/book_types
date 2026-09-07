package main

import (
	"fmt"
)

type N int

func (n N) Add(m N) N { return n + m }
func Sum[E interface{ Add(E) E }](s []E, a E) E {
	for _, v := range s {
		a = a.Add(v)
	}
	return a
}
func main() {
	fmt.Println(Sum([]N{1, 2}, N(10)))
}
