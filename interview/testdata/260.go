package main

import (
	"fmt"
)

func main() {
	c := make(chan [1]int, 1)
	a := [1]int{1}
	c <- a
	a[0] = 9
	fmt.Println(<-c)
}
