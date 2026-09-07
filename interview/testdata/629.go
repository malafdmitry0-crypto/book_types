package main

import (
	"fmt"
)

type N int

func main() {
	var end N = 4
	var sum N
	for i := range end {
		sum += i
	}
	fmt.Printf("%T %v", sum, sum)
}
