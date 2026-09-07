package main

import (
	"fmt"
)

type N int

func IsInt[T any](x T) bool {
	_, ok := any(x).(int)
	return ok
}
func main() {
	fmt.Println(IsInt(N(1)), IsInt(1))
}
