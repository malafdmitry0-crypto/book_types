package main

import (
	"fmt"
)

type N int

func (n *N) Inc() { *n++ }
func build[T interface{ Inc() }]() T {
	var t T
	t.Inc()
	return t
}
func main() {
	defer func() { fmt.Println(recover() != nil) }()
	fmt.Println(build[*N]())
}
