package generic

import "cmp"

// Zero возвращает нулевое значение T. Аргументов нет, выводить T не из чего:
// вызывать можно только как Zero[int]().
func Zero[T any]() T {
	var zero T
	return zero
}

// Number объединяет числовые типы, между которыми разрешена конверсия.
type Number interface {
	~int | ~int32 | ~int64 | ~float64
}

// Convert переводит число в другой числовой тип. From выводится из аргумента,
// To указывают явно: Convert[int64](n). Порядок параметров выбран ради этого.
func Convert[To, From Number](x From) To {
	return To(x)
}

// Max возвращает большее из двух значений одного упорядоченного типа.
func Max[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

// Identity возвращает аргумент без изменений.
func Identity[T any](x T) T { return x }

// Double удваивает число; используется как callback для Map.
func Double[T Number](x T) T { return x + x }
