package main

import (
	"fmt"
)

type A = int

func (A) F() {}
func main() {
	fmt.Println(1)
}
