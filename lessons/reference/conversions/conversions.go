package conversions

type Real interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

func Convert[To Real, From Real](x From) To {
	return To(x)
}

// result := Convert[int64](int32(42))
// To задан явно; From выводится из аргумента.
// Не компилируется: any не гарантирует допустимость конверсии.
// func Cast[To, From any](x From) To { return To(x) }

func Assert[T any](x any) (T, bool) {
	value, ok := x.(T)
	return value, ok
}

func IsString[T any](x T) bool {
	_, ok := any(x).(string)
	return ok
}
func Map[S ~[]E, E, R any](src S, convert func(E) R) []R {
	if src == nil {
		return nil
	}
	dst := make([]R, len(src))
	for i, item := range src {
		dst[i] = convert(item)
	}
	return dst
}

// result := Map([]int{1, 2}, func(n int) int64 { return int64(n) })
func Equal[T comparable](a, b T) bool { return a == b }

// Equal[any](42, 42)           // true, Go 1.20+
// Equal[any]([]int{1}, []int{1}) // panic: срезы несравнимы
