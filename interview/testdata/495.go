package main

import (
	"fmt"
)

type User struct {
	ID   int
	Name string
}

func Index[E any, K comparable](s []E, key func(E) K) map[K]E {
	out := map[K]E{}
	for _, v := range s {
		out[key(v)] = v
	}
	return out
}
func main() {
	m := Index([]User{{1, "a"}, {1, "b"}}, func(u User) int { return u.ID })
	fmt.Println(len(m), m[1].Name)
}
