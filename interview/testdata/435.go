package main

import (
	"fmt"
)

type Job struct{ Values []int }

func main() {
	a := Job{[]int{1, 2}}
	b := a
	b.Values[0] = 9
	fmt.Println(a.Values)
}
