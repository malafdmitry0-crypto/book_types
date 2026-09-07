package main

import (
	"fmt"
)

func zero[T any]() T { return nil }
func main() {
	fmt.Println(zero[int]())
}
