package main

import (
	"fmt"
)

type OF[E any] func(E, E) bool

func (f OF[E]) Less(a, b E) bool { return f(a, b) }
func main() {
	o := OF[string](func(a, b string) bool { return len(a) > len(b) })
	fmt.Println(o.Less("go", "x"))
}
