package main

import (
	"fmt"
)

func main() {
	const (
		_ = iota
		a
		_
		b
	)
	fmt.Println(a, b)
}
