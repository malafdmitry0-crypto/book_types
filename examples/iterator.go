package algorithms

import "iter"

// FindSeq returns a value, not an index: arbitrary sequences may lack indices.
func FindSeq[E any](src iter.Seq[E], pred func(E) bool) (E, bool) {
	for item := range src {
		if pred(item) {
			return item, true
		}
	}
	var zero E
	return zero, false
}

func AllSeq[E any](src iter.Seq[E], pred func(E) bool) bool {
	for item := range src {
		if !pred(item) {
			return false
		}
	}
	return true
}

// MapSeq is lazy: transform runs only when the returned sequence is consumed.
func MapSeq[E, R any](src iter.Seq[E], transform func(E) R) iter.Seq[R] {
	return func(yield func(R) bool) {
		for item := range src {
			if !yield(transform(item)) {
				return
			}
		}
	}
}
