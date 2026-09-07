package main

import (
	"fmt"
)

func main() {
	a := make(chan int)
	b := a
	close(b)
	_, ok := <-a
	fmt.Println(ok, a == b)
}
