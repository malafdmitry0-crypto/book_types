package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	ch <- 7
	func() { defer close(ch) }()
	a, ok := <-ch
	b, more := <-ch
	fmt.Println(a, ok, b, more)
}
