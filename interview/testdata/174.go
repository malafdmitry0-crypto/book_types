package main

import (
	"fmt"
)

func main() {
	a := []int{1}
	p := &a[0]
	a = append(a, 2)
	a[0] = 9
	fmt.Println(*p, a[0])
}
