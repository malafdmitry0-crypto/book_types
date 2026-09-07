package main

import (
	"fmt"
)

func id[T any](v T) T { return v }
func main() {
	fmt.Println(id(nil))
}
