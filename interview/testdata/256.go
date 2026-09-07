package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	var r <-chan int = c
	var a any = r
	_, ok := a.(<-chan int)
	fmt.Println(ok)
}
