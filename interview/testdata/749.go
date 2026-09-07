package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	var s chan<- int = ch
	close(s)
	_, ok := <-ch
	fmt.Println(ok)
}
