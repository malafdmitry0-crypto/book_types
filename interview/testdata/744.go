package main

import (
	"fmt"
)

func main() {
	var ch chan int
	n := 0
	idx := func() int {
		n++
		return 0
	}
	a := []int{0}
	select {
	case a[idx()] = <-ch:
	default:
	}
	fmt.Println(a, n)
}
