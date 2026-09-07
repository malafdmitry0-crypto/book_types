package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	ch <- 4
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
