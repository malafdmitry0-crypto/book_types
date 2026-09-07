package main

import (
	"fmt"
)

func each(s []int, f func(int)) {
	for _, v := range s {
		f(v)
	}
}
func main() {
	calls := []int{}
	var f func(int)
	f = func(n int) {
		calls = append(calls, n)
		f = func(n int) { calls = append(calls, n*10) }
	}
	each([]int{1, 2}, f)
	fmt.Println(calls)
}
