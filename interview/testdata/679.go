package main

import (
	"fmt"
)

type Box[T any] struct{ V T }

func (b Box[T]) Get() T { return b.V }
func main() {
	b := Box[string]{"go"}
	fmt.Printf("%T %v", b.Get(), b.Get())
}
