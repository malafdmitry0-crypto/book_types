package main

import (
	"fmt"
)

type Order struct{ Reverse bool }

func (o *Order) Less(a, b int) bool {
	if o.Reverse {
		return a > b
	}
	return a < b
}
func Pick[E any, O interface{ Less(E, E) bool }](a, b E) E {
	var o O
	if o.Less(b, a) {
		return b
	}
	return a
}
func main() {
	defer func() { fmt.Println(recover() != nil) }()
	fmt.Println(Pick[int, *Order](3, 2))
}
