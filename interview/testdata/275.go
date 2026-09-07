package main

import (
	"fmt"
)

type Box[T any] struct{ V T }

func (b Box[E]) Get() E { return b.V }
func main() {
	fmt.Println(Box[int]{3}.Get())
}
