package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	close(c)
	select {
	case v := <-c:
		fmt.Println(v)
	default:
		fmt.Println("default")
	}
}
