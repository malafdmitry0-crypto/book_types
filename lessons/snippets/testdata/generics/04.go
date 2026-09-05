// Фрагмент 4: book/chapters/11-generics.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func Equal[T comparable](a, b T) bool { return a == b }

// Equal[any](42, 42)           // true, Go 1.20+
// Equal[any]([]int{1}, []int{1}) // panic: срезы несравнимы
