package main

import (
	"fmt"
)

type Identity[E any] struct{}

func (Identity[E]) Map(e E) E { return e }
func New[E any, M interface{ Map(E) E }]() M {
	var m M
	return m
}
func main() {
	fmt.Println(New[int, Identity[int]]().Map(4))
}
