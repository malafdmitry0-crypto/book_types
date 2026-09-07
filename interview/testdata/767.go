package main

import (
	"fmt"
)

func main() {
	a := [3]int{1, 2, 3}
	sum := 0
	for i, v := range &a {
		if i == 0 {
			a[1] = 9
		}
		sum += v
	}
	fmt.Println(sum, a)
}
