package main

import (
	"fmt"
)

func convert[E, R any](x E) R { return R(x) }
func main() {
	fmt.Println(convert[int, string](1))
}
