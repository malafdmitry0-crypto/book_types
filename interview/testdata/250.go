package main

import (
	"fmt"
)

func main() {
	defer func() { fmt.Println(recover() != nil) }()
	var c chan int
	close(c)
}
