package main

import (
	"fmt"
)

func Group[E any, K comparable](s []E, key func(E) K) map[K][]E {
	out := map[K][]E{}
	for _, v := range s {
		k := key(v)
		out[k] = append(out[k], v)
	}
	return out
}
func main() {
	defer func() { fmt.Println(recover() != nil) }()
	_ = Group([]int{1}, func(n int) any { return []int{n} })
}
