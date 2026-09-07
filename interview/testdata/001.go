package main

import (
	"fmt"
)

func main() {
	const a = 5 / 2
	const b = 5 / 2.0
	fmt.Printf("%T %v | %T %v\n", a, a, b, b)
}
