// Фрагмент 1: book/chapters/15-adapters.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
// imports: "errors", "io/fs"
func pathFromError(err error) (string, bool) {
    var target *fs.PathError
    if errors.As(err, &target) {
        return target.Path, true
    }
    return "", false
}
