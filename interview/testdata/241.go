package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 1)
	var s chan<- int = c
	var r <-chan int = c
	s <- 7
	fmt.Println(<-r)
}
