package main

import (
	"fmt"
)

func main() {
	c := make(chan []int, 1)
	s := []int{1}
	c <- s
	s[0] = 9
	fmt.Println(<-c)
}
