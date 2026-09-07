package main

import (
	"fmt"
)

func main() {
	var ch chan int
	n := 0
	f := func() int {
		n++
		return 7
	}
	select {
	case ch <- f():
	default:
	}
	fmt.Println(n)
}
