package main

import (
	"fmt"
)

type Query struct {
	N     int
	Match func(int) bool
}

func find(q Query) int {
	for i := 0; i < q.N; i++ {
		if q.Match(i) {
			return i
		}
	}
	return -1
}
func main() {
	s := []int{1}
	q := Query{len(s), func(i int) bool {
		s = append(s, 9)
		return s[i] == 9
	}}
	fmt.Println(find(q), len(s))
}
