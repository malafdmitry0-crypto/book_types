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
	defer func() { fmt.Println(recover() != nil) }()
	f := Memo[any](func(x any) int { return 1 })
	fmt.Println(f([]int{1}))
}
