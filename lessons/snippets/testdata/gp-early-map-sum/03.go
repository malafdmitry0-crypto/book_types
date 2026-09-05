// Фрагмент 3: book/chapters/gp-early-map-sum.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type fractionText struct {
    Source    []fraction.Fraction
    Target    []string
    Transform func(fraction.Fraction) (string, error)
}

func (m fractionText) Len() int { return len(m.Source) }
func (m fractionText) Apply(i int) error {
    text, err := m.Transform(m.Source[i])
    if err != nil {
        return err
    }
    m.Target[i] = text
    return nil
}
