package main

import (
	"fmt"
)

func main() {
	s := (&[2]int{1, 2})[:]
	fmt.Println(s)
}
