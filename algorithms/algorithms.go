// Package algorithms demonstrates generic programming in the style of early Go:
// ordinary interfaces describe operations on collection indices. Element and
// result types stay inside caller-owned adapters. No type parameters, reflection,
// or conversion to a slice of empty interfaces is needed.
//
// Adapters must provide a stable, nonnegative length and accept every index in
// [0, Len()). Callbacks must not change the length during traversal. Nil adapters
// and panics inside their methods are not handled by this package.
package algorithms

// Predicate combines a collection with the condition to test at each index.
type Predicate interface {
	Len() int
	Match(i int) (bool, error)
}

// Find returns the first matching index, or -1 if no item matches.
// It visits indices from left to right and stops at the first match or error.
// On an error, the returned index is -1.
func Find(data Predicate) (int, error) {
	n := data.Len()
	for i := 0; i < n; i++ {
		matched, err := data.Match(i)
		if err != nil {
			return -1, err
		}
		if matched {
			return i, nil
		}
	}
	return -1, nil
}

// All reports whether every item matches. Empty collections return true.
// No further items are tested after the first mismatch or error.
func All(data Predicate) (bool, error) {
	n := data.Len()
	for i := 0; i < n; i++ {
		matched, err := data.Match(i)
		if err != nil {
			return false, err
		}
		if !matched {
			return false, nil
		}
	}
	return true, nil
}

// Any reports whether at least one item matches. Empty collections return false.
func Any(data Predicate) (bool, error) {
	index, err := Find(data)
	return index >= 0, err
}

// Mapping owns the input, output storage, and transformation operation.
// Apply transforms source item i and writes it to its destination.
// The caller decides output allocation and nil-versus-empty representation.
type Mapping interface {
	Len() int
	Apply(i int) error
}

// Map applies a transformation once at every index, from left to right.
// The typed result is held by the adapter, not returned as interface{}.
// On an error, previous writes remain; the failing operation may also have
// modified its destination. Later indices are not visited. No rollback occurs.
func Map(data Mapping) error {
	n := data.Len()
	for i := 0; i < n; i++ {
		if err := data.Apply(i); err != nil {
			return err
		}
	}
	return nil
}

// Summable owns an explicitly initialized accumulator and knows how to add
// input item i to it. Different types can require different initial zeros.
type Summable interface {
	Len() int
	Add(i int) error
}

// Sum adds each item, from left to right, to the adapter's existing accumulator.
// The caller must initialize/reset the accumulator before each independent sum.
// An empty input leaves it unchanged. An error stops traversal without rollback.
func Sum(data Summable) error {
	n := data.Len()
	for i := 0; i < n; i++ {
		if err := data.Add(i); err != nil {
			return err
		}
	}
	return nil
}

// Ordered describes comparisons of positions, not a universal element type.
// Compare returns a negative number, zero, or a positive number.
// A consistent ordering is the adapter's responsibility.
type Ordered interface {
	Len() int
	Compare(i, j int) (int, error)
}

// Min returns the index of the first minimum, or -1 for an empty collection.
// On a comparison error it returns -1 and the error. A singleton requires no
// comparison: validation of every item is the caller's responsibility.
func Min(data Ordered) (int, error) {
	n := data.Len()
	if n == 0 {
		return -1, nil
	}
	best := 0
	for i := 1; i < n; i++ {
		comparison, err := data.Compare(i, best)
		if err != nil {
			return -1, err
		}
		if comparison < 0 {
			best = i
		}
	}
	return best, nil
}

// Max is like Min but returns the index of the first maximum.
func Max(data Ordered) (int, error) {
	n := data.Len()
	if n == 0 {
		return -1, nil
	}
	best := 0
	for i := 1; i < n; i++ {
		comparison, err := data.Compare(i, best)
		if err != nil {
			return -1, err
		}
		if comparison > 0 {
			best = i
		}
	}
	return best, nil
}

// Sortable extends comparisons with in-place swaps, following the collection
// adapter idea of sort.Interface. Compare can return an error, unlike its Less.
type Sortable interface {
	Ordered
	Swap(i, j int)
}

// Sort performs stable insertion sort in place: O(n^2) comparisons/swaps in the
// worst case and O(1) auxiliary storage. This deliberately small implementation
// illustrates an interface-based algorithm; it is not Go's historical sort code.
// On an error the collection may already be partially reordered; no rollback
// occurs. Empty and singleton collections make no comparisons.
func Sort(data Sortable) error {
	n := data.Len()
	for i := 1; i < n; i++ {
		for j := i; j > 0; j-- {
			comparison, err := data.Compare(j, j-1)
			if err != nil {
				return err
			}
			if comparison >= 0 {
				break
			}
			data.Swap(j, j-1)
		}
	}
	return nil
}
