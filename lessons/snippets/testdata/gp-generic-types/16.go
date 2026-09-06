// Фрагмент 16: book/chapters/gp-generic-types.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
var s containers.Stack[int]
containers.Fill(&s, 1, 2, 3)
// containers.Fill(s, 4) // ошибка компиляции: Stack[int] не реализует Container[int]
var _ fmt.Stringer = s // String объявлен на значении, подходят и Stack[int], и *Stack[int]
