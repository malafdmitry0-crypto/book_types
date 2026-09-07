package main

import (
	"fmt"
)

type Flag bool

func main() {
	var a Flag = true
	fmt.Printf("%T %v\n", !a, !a)
}
