package main

import (
	"fmt"
)

func makePtr[T any]() *T { return new(T) }
func main() {
	p := makePtr[int]()
	fmt.Printf("%T %v\n", p, *p)
}
