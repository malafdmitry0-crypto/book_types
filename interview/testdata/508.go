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
func Pick[E any, O interface{ Less(E, E) bool }](a, b E, o O) E {
	if o.Less(b, a) {
		return b
	}
	return a
}
func main() {
	fmt.Println(Pick(3, 2, &Order{true}))
}
