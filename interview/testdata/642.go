package main

import (
	"fmt"
)

type N struct{ V int }

func (n N) Inc() int {
	n.V++
	return n.V
}

type Counter interface{ Inc() int }

func main() {
	n := N{3}
	var c Counter = n
	fmt.Println(c.Inc(), n.V)
}
