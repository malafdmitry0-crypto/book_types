package main

import (
	"fmt"
)

func main() {
	n := 0
	f := func() *[3]int {
		n++
		return nil
	}
	fmt.Println(len(f()), n)
}
