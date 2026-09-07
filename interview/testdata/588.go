package main

import (
	"fmt"
	"iter"
	"slices"
	"strconv"
)

func Parse(s iter.Seq[string]) iter.Seq2[int, error] {
	return func(yield func(int, error) bool) {
		for v := range s {
			n, e := strconv.Atoi(v)
			if !yield(n, e) {
				return
			}
		}
	}
}
func main() {
	seq := Parse(slices.Values([]string{"1", "x", "3"}))
	var out []int
	for n, e := range seq {
		if e != nil {
			continue
		}
		out = append(out, n)
	}
	fmt.Println(out)
}
