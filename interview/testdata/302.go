package main

import (
	"fmt"
)

func same[T any](a, b T) T { return a }
func main() {
	var a int32 = 1
	var b int64 = 2
	fmt.Println(same(a, b))
}
