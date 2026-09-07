package main

import (
	"fmt"
)

type A int

func main() {
	var a A = 1
	fmt.Printf("%T %v\n", a+2, a+2)
}
