package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2}
	for _, v := range a {
		v *= 10
	}
	fmt.Println(a)
}
