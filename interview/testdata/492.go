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
	groups := Group([]string{"a", "go", "b"}, func(s string) int { return len(s) })
	fmt.Printf("%T %v %v\n", groups, groups[1], groups[2])
}
