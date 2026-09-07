package main

import (
	"fmt"
)

type C chan int
type R <-chan int

func main() {
	c := make(C, 1)
	var r R = (<-chan int)(c)
	c <- 8
	fmt.Println(<-r)
}
