package main

import (
	"fmt"
)

type T int

func (t *T) Inc() { *t++ }
func main() {
	var t T
	t.Inc()
	fmt.Println(t)
}
