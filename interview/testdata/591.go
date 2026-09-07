package main

import (
	"fmt"
	"iter"
)

func main() {
	seq := func(yield func(int) bool) { yield(1) }
	next, stop := iter.Pull(seq)
	next()
	stop()
	stop()
	v, ok := next()
	fmt.Println(v, ok)
}
