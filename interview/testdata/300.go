package main

type Set[K comparable] struct{}

func ordered[T ~int](T) {}
func (s Set[K]) F(k K)  { ordered(k) }
func main() {
	_ = Set[int]{}
}
