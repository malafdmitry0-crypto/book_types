package main

import (
	"fmt"
)

type N int

func isInt[T any](v T) bool {
	_, ok := any(v).(int)
	return ok
}
func main() {
	fmt.Println(isInt(1), isInt(N(1)))
}
