package main

import (
	"fmt"
)

type N int

func main() {
	s := func(yield func(N) bool) { yield(7) }
	for v := range s {
		fmt.Printf("%T %v\n", v, v)
	}
}
