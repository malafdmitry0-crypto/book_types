package main

import (
	"fmt"
	"iter"
)

func main() {
	done := false
	seq := func(yield func(int) bool) {
		defer func() { done = true }()
		for i := 0; i < 3; i++ {
			if !yield(i) {
				return
			}
		}
	}
	next, stop := iter.Pull[int](seq)
	v, ok := next()
	fmt.Print(v, " ", ok, " ", done, " ")
	stop()
	fmt.Println(done)
}
