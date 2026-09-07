package main

import (
	"fmt"
)

func isInt[T any](v T) bool {
	_, ok := v.(int)
	return ok
}
func main() {
	fmt.Println(isInt(1))
}
