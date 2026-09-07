package main

import (
	"fmt"
)

func f(a [2]int) { a[0] = 9 }
func main() {
	a := [2]int{1, 2}
	f(a)
	fmt.Println(a)
}
