package main

import (
	"fmt"
)

type T int

func (t *T) Nil() bool { return t == nil }

type O struct{ *T }

func main() {
	var o O
	fmt.Println(o.Nil())
}
