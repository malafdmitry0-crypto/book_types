package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)
	c <- 7
	close(c)
	a, okA := <-c
	b, okB := <-c
	fmt.Println(a, okA, b, okB)
}
