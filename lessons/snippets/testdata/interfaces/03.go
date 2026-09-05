// Фрагмент 3: book/chapters/07-interfaces.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
// import "io"
func inspect(rw io.ReadWriter) {
    var reader io.Reader = rw // все нужные методы гарантированы статически
    writer, ok := reader.(io.Writer) // проверка динамического значения
    _, _ = writer, ok
}
