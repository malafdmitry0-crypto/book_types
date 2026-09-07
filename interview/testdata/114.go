package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	var r <-chan int = c
	var s chan<- int = c
	fmt.Println(r == s)
}
