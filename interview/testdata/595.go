package main

import (
	"iter"
)

type Source[E any] interface{ All() iter.Seq[E] }

func main() {
	var src Source[int]
	var dst Source[any] = src
	_ = dst
}
