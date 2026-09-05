package generic

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

type Addable[T any] interface {
	Add(T) (T, error)
}

func SumValues[T Addable[T]](values []T, initial T) (T, error) {
	total := initial
	for _, value := range values {
		next, err := total.Add(value)
		if err != nil {
			return total, err
		}
		total = next
	}
	return total, nil
}
func MapValues[E, R any](values []E, transform func(E) (R, error)) ([]R, error) {
	if values == nil {
		return nil, nil
	}
	out := make([]R, len(values))
	for i, value := range values {
		result, err := transform(value)
		if err != nil {
			return out, err
		}
		out[i] = result
	}
	return out, nil
}
func Contains[E comparable](src []E, target E) bool {
	return Find(src, func(item E) bool { return item == target }) >= 0
}
func Clone[S ~[]E, E any](src S) S {
	if src == nil {
		return nil
	}
	out := make(S, len(src))
	copy(out, src)
	return out
}

type Named interface {
	Name() string
}

func Names[E Named](src []E) []string {
	return Map(src, func(item E) string { return item.Name() })
}

type Equaler[E any] interface {
	Equal(E) bool
}

func ContainsEqual[E Equaler[E]](src []E, target E) bool {
	return Find(src, func(item E) bool { return item.Equal(target) }) >= 0
}
