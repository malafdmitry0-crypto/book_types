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

// FilterSeq is lazy too: keep runs only for elements the consumer requests.
func FilterSeq[E any](src iter.Seq[E], keep func(E) bool) iter.Seq[E] {
	return func(yield func(E) bool) {
		for item := range src {
			if keep(item) && !yield(item) {
				return
			}
		}
	}
}

// MapSeqErr reports each failure as a pair and moves on to the next element.
// Stopping after the first error is the consumer's decision, not the source's.
func MapSeqErr[E, R any](src iter.Seq[E], transform func(E) (R, error)) iter.Seq2[R, error] {
	return func(yield func(R, error) bool) {
		for item := range src {
			if !yield(transform(item)) {
				return
			}
		}
	}
}

// Count yields 0, 1, ..., n-1 and runs cleanup exactly once when the walk
// ends, whether the sequence is exhausted or the consumer stops early.
func Count(n int, cleanup func()) iter.Seq[int] {
	return func(yield func(int) bool) {
		defer cleanup()
		for i := 0; i < n; i++ {
			if !yield(i) {
				return
			}
		}
	}
}

// Interleave alternates elements of a and b: a0, b0, a1, b1, ...
// A push source cannot pause itself, so both inputs are pulled.
func Interleave[E any](a, b iter.Seq[E]) iter.Seq[E] {
	return func(yield func(E) bool) {
		nextA, stopA := iter.Pull(a)
		defer stopA()
		nextB, stopB := iter.Pull(b)
		defer stopB()
		for {
			itemA, okA := nextA()
			if okA && !yield(itemA) {
				return
			}
			itemB, okB := nextB()
			if okB && !yield(itemB) {
				return
			}
			if !okA && !okB {
				return
			}
		}
	}
}
