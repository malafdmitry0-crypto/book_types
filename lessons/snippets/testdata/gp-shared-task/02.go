// Фрагмент 2: book/chapters/gp-shared-task.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type Predicate func(fraction.Fraction) (bool, error)
type Step func(fraction.Fraction, fraction.Fraction) (fraction.Fraction, error)
