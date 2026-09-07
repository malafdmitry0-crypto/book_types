package main

import (
	"fmt"
)

func IsInt[T any](x T) bool {
	_, ok := x.(int)
	return ok
}
func main() {
	fmt.Println(IsInt(1))
}
