package main

import (
	"fmt"
)

type T int

func (t *T) Get() T { return *t }
func main() {
	a, b := T(1), T(2)
	p := &a
	f := p.Get
	p = &b
	fmt.Println(f(), p.Get())
}
