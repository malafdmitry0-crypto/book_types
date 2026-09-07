package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	for i, v := range a {
		fmt.Print(v)
		if i == 0 {
			a[1] = 9
		}
	}
	fmt.Println()
}
