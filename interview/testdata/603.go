package main

import (
	"fmt"
)

func f() int {
	x := 3
	defer func() { x += 4 }()
	return x
}
func main() {
	fmt.Println(f())
}
