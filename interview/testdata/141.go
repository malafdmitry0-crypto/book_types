package main

import (
	"fmt"
)

type T int

func (t *T) Inc() { *t++ }

type O struct{ T }

func main() {
	var o O
	o.Inc()
	fmt.Println(o.T)
}
