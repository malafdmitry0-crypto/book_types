package main

import (
	"fmt"
)

type T int

func (t *T) Inc() { *t++ }
func main() {
	t := T(1)
	var i interface{ Inc() } = &t
	i.Inc()
	fmt.Println(t)
}
