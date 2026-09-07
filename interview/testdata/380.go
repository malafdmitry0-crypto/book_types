package main

import (
	"fmt"
)

type S string

func main() {
	var s S = "go"
	fmt.Printf("%T %v\n", s+"!", s+"!")
}
