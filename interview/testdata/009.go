package main

import (
	"fmt"
)

func main() {
	const (
		a, b = iota, iota + 10
		c, d
	)
	fmt.Println(a, b, c, d)
}
