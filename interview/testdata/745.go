package main

import (
	"fmt"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	ch := make(chan int)
	close(ch)
	select {
	case ch <- 1:
		fmt.Println("sent")
	default:
		fmt.Println("default")
	}
}
