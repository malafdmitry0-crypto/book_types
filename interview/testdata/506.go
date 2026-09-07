package main

import (
	"fmt"
)

type Ascending struct{}

func (Ascending) Less(a, b int) bool { return a < b }
func Pick[E any, O interface{ Less(E, E) bool }](a, b E) E {
	var o O
	if o.Less(b, a) {
		return b
	}
	return a
}
func main() {
	fmt.Println(Pick[int, Ascending](3, 2))
}
