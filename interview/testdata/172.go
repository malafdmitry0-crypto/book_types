package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2}
	for _, v := range a {
		a = append(a, v)
	}
	fmt.Println(a)
}
