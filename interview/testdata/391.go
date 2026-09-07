package main

import (
	"fmt"
)

type N int

func (n *N) Inc() { *n++ }
func build[T any, P interface {
	*T
	Inc()
}]() T {
	p := P(new(T))
	p.Inc()
	return *p
}
func main() {
	fmt.Println(build[N, *N]())
}
