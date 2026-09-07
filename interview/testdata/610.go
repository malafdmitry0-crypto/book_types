package main

import (
	"fmt"
)

func f() (n int) {
	n = 5
	defer func() {
		if recover() != nil {
			n = 12
		}
	}()
	panic("x")
}
func main() {
	fmt.Println(f())
}
