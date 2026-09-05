// Фрагмент 1: book/chapters/gp-interface-limits.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type Addable interface {
    Add(Addable) (Addable, error)
}

// Не компилируется: у Fraction другая сигнатура Add.
// var _ Addable = fraction.Fraction{}
