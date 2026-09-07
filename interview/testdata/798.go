package main

import (
	"fmt"
	"iter"
	"slices"
)

func main() {
	next, stop := iter.Pull(slices.Values([]string{"go", "types"}))
	next()
	stop()
	v, ok := next()
	fmt.Printf("%q %v", v, ok)
}
