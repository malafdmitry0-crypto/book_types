package main

import (
	"fmt"
)

func main() {
	n := 0
	next := func() int {
		n++
		return n
	}
	fmt.Println(next(), next(), next())
}
