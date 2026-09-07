package main

import (
	"fmt"
)

func main() {
	const z = 2 + 0i
	var x float64 = z
	fmt.Printf("%T %v\n", x, x)
}
