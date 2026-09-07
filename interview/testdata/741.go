package main

import (
	"fmt"
)

func main() {
	var ch chan int
	select {
	case ch <- 1:
		fmt.Println("sent")
	default:
		fmt.Println("default")
	}
}
