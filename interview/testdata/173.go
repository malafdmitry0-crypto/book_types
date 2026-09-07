package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2}
	for i, v := range a {
		if i == 0 {
			a = append(a, 3)
			a[1] = 9
		}
		fmt.Print(v)
	}
	fmt.Println(a)
}
