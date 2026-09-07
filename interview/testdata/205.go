package main

import (
	"fmt"
)

func f(m map[int]int) { m = map[int]int{1: 9} }
func main() {
	m := map[int]int{1: 2}
	f(m)
	fmt.Println(m[1])
}
