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
	s := []string{"a", "go", "x"}
	job := Query{len(s), func(i int) bool { return len(s[i]) == 2 }}
	fmt.Println(find(job))
}
