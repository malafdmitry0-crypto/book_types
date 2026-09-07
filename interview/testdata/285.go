package main

import (
	"fmt"
)

func twice[T ~int | ~string](v T) T { return v + v }
func main() {
	fmt.Println(twice(3), twice("go"))
}
