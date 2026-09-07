package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	x := 0
	go func() {
		x = 9
		ch <- 1
	}()
	<-ch
	fmt.Println(x)
}
