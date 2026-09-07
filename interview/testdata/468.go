package main

import (
	"fmt"
)

func memo(f func(int) int) func(int) int {
	cache := map[int]int{}
	return func(n int) int {
		if v, ok := cache[n]; ok {
			return v
		}
		v := f(n)
		cache[n] = v
		return v
	}
}
func main() {
	calls := 0
	f := memo(func(int) int {
		calls++
		return 0
	})
	fmt.Println(f(1), f(1), calls)
}
