package invalid

import "slices"

type Set[K comparable] struct{ m map[K]struct{} }

func (s Set[K]) Sorted() []K {
	var keys []K
	for k := range s.m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}
