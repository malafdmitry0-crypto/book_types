package main

import (
	"fmt"
)

type Item struct{ N int }

func visit(s []Item, f func(*Item)) {
	for _, v := range s {
		f(&v)
	}
}
func main() {
	s := []Item{{1}}
	visit(s, func(v *Item) { v.N = 9 })
	fmt.Println(s[0].N)
}
