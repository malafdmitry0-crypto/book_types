// Фрагмент 10: book/chapters/gp-shared-task.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type Search struct {
	Values []fraction.Fraction
	Test   func(fraction.Fraction) (bool, error)
}

func (s Search) Len() int                  { return len(s.Values) }
func (s Search) Match(i int) (bool, error) { return s.Test(s.Values[i]) }

type Fold struct {
	Values []fraction.Fraction
	Total  fraction.Fraction
	Step   func(fraction.Fraction, fraction.Fraction) (fraction.Fraction, error)
}

func (f *Fold) Len() int { return len(f.Values) }
func (f *Fold) Add(i int) error {
	next, err := f.Step(f.Total, f.Values[i])
	if err != nil {
		return err
	}
	f.Total = next
	return nil
}
