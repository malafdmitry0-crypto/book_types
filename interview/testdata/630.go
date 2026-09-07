package main

import (
	"fmt"
)

func main() {
	n := -3
	count := 0
	for range n {
		count++
	}
	fmt.Println(count)
}
