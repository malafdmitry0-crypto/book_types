// Фрагмент 5: book/chapters/gp-interface-limits.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
var printable []Stringer
if fractions != nil {
    printable = make([]Stringer, len(fractions))
}
for i, value := range fractions {
    printable[i] = value
}
PrintAll(printable)
