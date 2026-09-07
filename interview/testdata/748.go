package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	var r <-chan int = ch
	close(r)
	fmt.Println("done")
}
