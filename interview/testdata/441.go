package main

import (
	"fmt"
)

func all(s []int, p func(int) bool) bool {
	for _, v := range s {
		if !p(v) {
			return false
		}
	}
	return true
}
func main() {
	calls := 0
	ok := all([]int{0, 1}, func(n int) bool {
		calls++
		if n == 1 {
			panic("late")
		}
		return n > 0
	})
	fmt.Println(ok, calls)
}
