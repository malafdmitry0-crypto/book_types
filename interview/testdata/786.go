package main

import (
	"fmt"
)

func Memo[K comparable, V any](f func(K) V) func(K) V {
	cache := map[K]V{}
	return func(k K) V {
		if v, ok := cache[k]; ok {
			return v
		}
		v := f(k)
		cache[k] = v
		return v
	}
}
func main() {
	n := 0
	f := Memo(func(x int) int {
		n++
		return 0
	})
	fmt.Println(f(2), f(2), n)
}
