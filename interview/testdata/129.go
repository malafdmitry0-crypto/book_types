package main

import (
	"fmt"
)

type T int

func (t T) Get() T { return t }
func main() {
	f := T.Get
	fmt.Printf("%T %v\n", f, f(7))
}
