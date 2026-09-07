package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	var s chan<- int = c
	close(s)
	v, ok := <-c
	fmt.Println(v, ok)
}
