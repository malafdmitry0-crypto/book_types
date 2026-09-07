package main

import (
	"fmt"
)

func f[T, R any](v T) T { return v }
func main() {
	fmt.Println(f(1))
}
