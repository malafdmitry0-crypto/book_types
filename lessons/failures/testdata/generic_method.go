package invalid
type Stack[T any] struct{items []T}
func (s Stack[T]) Map[R any](f func(T) R) Stack[R] {return Stack[R]{}}
