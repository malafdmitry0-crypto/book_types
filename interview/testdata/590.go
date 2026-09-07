package main

import (
	"fmt"
	"iter"
)

func main() {
	closed := false
	seq := func(yield func(int) bool) {
		defer func() { closed = true }()
		if !yield(1) {
			return
		}
		yield(2)
	}
	next, stop := iter.Pull(seq)
	v, ok := next()
	fmt.Println(v, ok, closed)
	stop()
	fmt.Println(closed)
}
