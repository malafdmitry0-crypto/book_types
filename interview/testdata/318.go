package main

import (
	"fmt"
)

type C chan int

func read[T any](c <-chan T) T { return <-c }
func main() {
	c := make(C, 1)
	c <- 7
	fmt.Println(read(c))
}
