package main

import (
	"fmt"
)

func main() {
	n := 0
Outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if j == 1 {
				continue Outer
			}
			n++
		}
	}
	fmt.Println(n)
}
