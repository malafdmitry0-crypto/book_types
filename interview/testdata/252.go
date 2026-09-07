package main

import (
	"fmt"
)

func main() {
	c := make(chan *int, 1)
	c <- nil
	close(c)
	p, ok := <-c
	fmt.Println(p == nil, ok)
}
