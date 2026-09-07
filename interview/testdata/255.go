package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	var a any = c
	_, ok := a.(<-chan int)
	fmt.Println(ok)
}
