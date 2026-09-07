package main

import (
	"fmt"
)

func main() {
	const (
		a uint8 = iota
		b
	)
	fmt.Printf("%T %v\n", b, b)
}
