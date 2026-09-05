// Фрагмент 1: book/chapters/09-aliases.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type UserID = int64 // тот же тип
type OrderID int64 // отдельный тип

func aliasExample() {
    var n int64 = 42
    var user UserID = n // конверсия не нужна
    order := OrderID(n)
    _, _ = user, order
}
