package main

import (
	"fmt"
)

type A int
type B = A

func main() {
	var x B
	fmt.Printf("%T\n", x)
}
