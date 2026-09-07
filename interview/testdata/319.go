package main

import (
	"fmt"
)

func first[T any](v ...T) T { return v[0] }
func main() {
	fmt.Printf("%T %v\n", first(1, 2.5), first(1, 2.5))
}
