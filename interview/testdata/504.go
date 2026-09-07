package main

import (
	"fmt"
)

type Order[E any] interface{ Less(E, E) bool }
type LenOrder struct{}

func (LenOrder) Less(a, b string) bool { return len(a) < len(b) }
func Pick[E any](s []E, o Order[E]) E {
	v := s[0]
	for _, x := range s[1:] {
		if o.Less(x, v) {
			v = x
		}
	}
	return v
}
func main() {
	s := []string{"aaa", "b"}
	fmt.Println(Pick(s, LenOrder{}))
}
