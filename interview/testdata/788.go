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
	base := func(x int) int {
		n++
		return x
	}
	a := Memo(base)
	b := Memo(base)
	a(1)
	a(1)
	b(1)
	fmt.Println(n)
}
