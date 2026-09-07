package main

import (
	"fmt"
)

func f() int {
	n := 1
	defer func() { n++ }()
	return n
}
func main() {
	fmt.Println(f())
}
