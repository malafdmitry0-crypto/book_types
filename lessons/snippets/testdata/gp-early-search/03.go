// Фрагмент 3: book/chapters/gp-early-search.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type fractionQuery struct {
    Values []fraction.Fraction
    Test   func(fraction.Fraction) (bool, error)
}

func (q fractionQuery) Len() int { return len(q.Values) }
func (q fractionQuery) Match(i int) (bool, error) {
    return q.Test(q.Values[i])
}
