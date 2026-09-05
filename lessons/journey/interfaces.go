package journey

import (
	"gotypes/fraction"
)

// Search связывает срез дробей с Predicate из нашего раннего algorithms.
type Search struct {
	Values []fraction.Fraction
	Test   func(fraction.Fraction) (bool, error)
}

func (s Search) Len() int                  { return len(s.Values) }
func (s Search) Match(i int) (bool, error) { return s.Test(s.Values[i]) }

// Selection хранит отдельный результат; ошибка не добавляет текущий элемент.
type Selection struct {
	Values []fraction.Fraction
	Test   func(fraction.Fraction) (bool, error)
	Result []fraction.Fraction
}

func (s *Selection) Len() int { return len(s.Values) }
func (s *Selection) Select(i int) error {
	matched, err := s.Test(s.Values[i])
	if err != nil {
		return err
	}
	if matched {
		s.Result = append(s.Result, s.Values[i])
	}
	return nil
}

type Selectable interface {
	Len() int
	Select(int) error
}

func InterfaceFilter(src Selectable) error {
	n := src.Len()
	for i := 0; i < n; i++ {
		if err := src.Select(i); err != nil {
			return err
		}
	}
	return nil
}

// Fold заменяет аккумулятор только после успешного шага.
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

// Fold удовлетворяет algorithms.Summable: общий Reduce может использовать тот же обход.
