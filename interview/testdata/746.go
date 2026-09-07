package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 2)
	ch <- 2
	ch <- 5
	close(ch)
	sum := 0
	for x := range ch {
		sum += x
	}
	fmt.Println(sum)
}
