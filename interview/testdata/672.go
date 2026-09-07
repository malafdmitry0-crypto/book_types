package main

import (
	"fmt"
)

func Id[T any](x T) T { return x }
func main() {
	var f func(string) string = Id
	fmt.Println(f("go"))
}
