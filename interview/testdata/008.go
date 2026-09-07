package main

import (
	"fmt"
)

func main() {
	const (
		a = 1 << iota
		b
		c = 10
		d
	)
	fmt.Println(a, b, c, d)
}
