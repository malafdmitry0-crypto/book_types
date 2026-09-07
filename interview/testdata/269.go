package main

import (
	"fmt"
)

func mapOne[E, R any](x E, f func(E) R) R { return f(x) }
func main() {
	fmt.Println(mapOne(7, func(n int) string { return fmt.Sprint(n) }))
}
