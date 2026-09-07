package main

import (
	"fmt"
)

func main() {
	c := make(chan any, 1)
	c <- nil
	close(c)
	a, ok := <-c
	fmt.Println(a == nil, ok)
}
