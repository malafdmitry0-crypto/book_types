package main

import (
	"fmt"
)

type User struct {
	N    int
	Name string
}

func MinBy[E any](s []E, less func(E, E) bool) E {
	m := s[0]
	for _, v := range s[1:] {
		if less(v, m) {
			m = v
		}
	}
	return m
}
func main() {
	v := MinBy([]User{{2, "a"}, {2, "b"}}, func(a, b User) bool { return a.N < b.N })
	fmt.Println(v.Name)
}
