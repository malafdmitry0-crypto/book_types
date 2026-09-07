package main

import (
	"fmt"
)

type N int

func main() {
	var n N = 3
	for i := range n {
		fmt.Printf("%T:%v ", i, i)
	}
	fmt.Println()
}
