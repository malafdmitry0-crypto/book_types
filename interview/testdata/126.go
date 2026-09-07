package main

import (
	"fmt"
)

type T int

func (t T) Get() T { return t }
func main() {
	t := T(1)
	f := t.Get
	t = 2
	fmt.Println(f(), t.Get())
}
