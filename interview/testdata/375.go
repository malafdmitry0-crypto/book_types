package main

import (
	"fmt"
)

type B byte

func main() {
	b := []B("go")
	fmt.Printf("%T %v\n", b, b)
}
