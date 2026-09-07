package main

import (
	"fmt"
)

func memo(f func(int) int) func(int) int {
	cache := map[int]int{}
	return func(n int) int {
		if cache[n] != 0 {
			return cache[n]
		}
		cache[n] = f(n)
		return cache[n]
	}
}
func main() {
	calls := 0
	f := memo(func(int) int {
		calls++
		return 0
	})
	f(1)
	f(1)
	fmt.Println(calls)
}
