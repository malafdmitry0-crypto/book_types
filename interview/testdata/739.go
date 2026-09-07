package main

import (
	"fmt"
)

func main() {
	n := 0
	f := func() int {
		n++
		return n
	}
	g := f
	fmt.Println(f(), g(), f())
}
