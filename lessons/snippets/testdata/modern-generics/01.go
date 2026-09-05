// Фрагмент 1: book/chapters/13-modern-generics.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func Identity[T any](x T) T { return x }

var intIdentity func(int) int = Identity // Go 1.21+
// До этого можно было явно написать Identity[int].
