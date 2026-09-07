package main

import (
	"fmt"
)

type S struct{}

func (S) Read() int                         { return 7 }
func read[T any](s interface{ Read() T }) T { return s.Read() }
func main() {
	fmt.Printf("%T %v\n", read(S{}), read(S{}))
}
