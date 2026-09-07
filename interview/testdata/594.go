package main

import (
	"fmt"
	"iter"
	"slices"
)

type Source[E any] interface{ All() iter.Seq[E] }
type Values[E any] []E

func (v Values[E]) All() iter.Seq[E] { return slices.Values(v) }
func main() {
	var src Source[int] = Values[int]{1, 2}
	fmt.Println(slices.Collect(src.All()))
}
