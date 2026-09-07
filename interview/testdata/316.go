package main

import (
	"fmt"
)

func id[T any](v T) T                 { return v }
func apply[T any](v T, f func(T) T) T { return f(v) }
func main() {
	fmt.Println(apply(3, id))
}
