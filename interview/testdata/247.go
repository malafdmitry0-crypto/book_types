package main

import (
	"fmt"
)

type N int

func main() {
	c := make(chan N)
	close(c)
	v, ok := <-c
	fmt.Printf("%T %v %v\n", v, v, ok)
}
