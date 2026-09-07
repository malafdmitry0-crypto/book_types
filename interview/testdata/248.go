package main

import (
	"fmt"
)

func main() {
	var c chan int
	select {
	case <-c:
		fmt.Println("receive")
	default:
		fmt.Println("default")
	}
}
