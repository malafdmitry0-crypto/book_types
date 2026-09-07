package main

import (
	"fmt"
)

func id[T any](v T) T { return v }
func main() {
	f := id
	fmt.Println(f(1))
}
