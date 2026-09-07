package main

import (
	"fmt"
)

type T int

func (t *T) Inc() { *t++ }

type O struct{ *T }

func main() {
	x := T(1)
	a := O{&x}
	b := a
	b.Inc()
	fmt.Println(a.T == b.T, x)
}
