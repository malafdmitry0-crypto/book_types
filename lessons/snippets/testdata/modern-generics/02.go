// Фрагмент 2: book/chapters/13-modern-generics.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type Set[T comparable] = map[T]bool // Go 1.24+

func aliasSetExample() {
    raw := map[string]bool{"go": true}
    var names Set[string] = raw // идентичный тип
    _ = names
}
