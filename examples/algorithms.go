package algorithms

// User is the data model used throughout the textbook.
type User struct {
	Name   string
	Age    int
	Active bool
}

// FindUser returns the index of the first match, or -1.
func FindUser(src []User, pred func(User) bool) int {
	for i, user := range src {
		if pred(user) {
			return i
		}
	}
	return -1
}

func AllUsers(src []User, pred func(User) bool) bool {
	for _, user := range src {
		if !pred(user) {
			return false
		}
	}
	return true
}

func MapUserNames(src []User, transform func(User) string) []string {
	if src == nil {
		return nil
	}
	out := make([]string, len(src))
	for i, user := range src {
		out[i] = transform(user)
	}
	return out
}

// FindIndex is a linear search. Unlike sort.Search, it needs no monotonicity.
// n must be nonnegative.
func FindIndex(n int, pred func(int) bool) int {
	for i := 0; i < n; i++ {
		if pred(i) {
			return i
		}
	}
	return -1
}

func AllIndex(n int, pred func(int) bool) bool {
	return FindIndex(n, func(i int) bool { return !pred(i) }) == -1
}

// MapInto delegates storage and transformation to the caller.
func MapInto(n int, assign func(int)) {
	for i := 0; i < n; i++ {
		assign(i)
	}
}

// MatchSequence hides the element type behind the behavior needed by search.
type MatchSequence interface {
	Len() int
	Match(int) bool
}

type UserQuery struct {
	Values []User
	Pred   func(User) bool
}

func (q UserQuery) Len() int         { return len(q.Values) }
func (q UserQuery) Match(i int) bool { return q.Pred(q.Values[i]) }

func FindInterface(src MatchSequence) int {
	for i := 0; i < src.Len(); i++ {
		if src.Match(i) {
			return i
		}
	}
	return -1
}

func AllInterface(src MatchSequence) bool {
	for i := 0; i < src.Len(); i++ {
		if !src.Match(i) {
			return false
		}
	}
	return true
}

// MappingJob also hides allocation and the output type in the adapter.
type MappingJob interface {
	Len() int
	Apply(int)
}

func RunMap(job MappingJob) {
	for i := 0; i < job.Len(); i++ {
		job.Apply(i)
	}
}

type UserNameJob struct {
	Source []User
	Target []string
	Fn     func(User) string
}

func (j UserNameJob) Len() int    { return len(j.Source) }
func (j UserNameJob) Apply(i int) { j.Target[i] = j.Fn(j.Source[i]) }

// These APIs deliberately use the pre-1.18 spelling.
func FindAny(src []interface{}, pred func(interface{}) bool) int {
	for i, item := range src {
		if pred(item) {
			return i
		}
	}
	return -1
}

func AllAny(src []interface{}, pred func(interface{}) bool) bool {
	for _, item := range src {
		if !pred(item) {
			return false
		}
	}
	return true
}

func MapAny(src []interface{}, transform func(interface{}) interface{}) []interface{} {
	if src == nil {
		return nil
	}
	out := make([]interface{}, len(src))
	for i, item := range src {
		out[i] = transform(item)
	}
	return out
}

func Find[E any](src []E, pred func(E) bool) int {
	for i, item := range src {
		if pred(item) {
			return i
		}
	}
	return -1
}

func All[E any](src []E, pred func(E) bool) bool {
	for _, item := range src {
		if !pred(item) {
			return false
		}
	}
	return true
}

func Map[E, R any](src []E, transform func(E) R) []R {
	if src == nil {
		return nil
	}
	out := make([]R, len(src))
	for i, item := range src {
		out[i] = transform(item)
	}
	return out
}

func Contains[E comparable](src []E, target E) bool {
	return Find(src, func(item E) bool { return item == target }) >= 0
}

// Clone preserves a defined slice type and nilness; the copy is shallow.
func Clone[S ~[]E, E any](src S) S {
	if src == nil {
		return nil
	}
	out := make(S, len(src))
	copy(out, src)
	return out
}
