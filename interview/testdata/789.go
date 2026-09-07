package main

import (
	"fmt"
)

type Option[T any] struct {
	v  T
	ok bool
}

func Some[T any](v T) Option[T]    { return Option[T]{v, true} }
func (o Option[T]) Get() (T, bool) { return o.v, o.ok }
func main() {
	o := Some[*int](nil)
	v, ok := o.Get()
	fmt.Println(ok, v == nil)
}
