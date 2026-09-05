package journey_test

import (
	"errors"
	"fmt"
	"gotypes/fraction"
	"gotypes/lessons/generic"
	"gotypes/lessons/journey"
)

// Решение: ошибка имеет приоритет над true, частичный Filter сохраняется.
func Example_partialFilter() {
	values := []fraction.Fraction{{Numerator: 1, Denominator: 2}, {Numerator: 1, Denominator: 3}, {Numerator: 1, Denominator: 4}}
	stop := errors.New("stop")
	for _, approach := range journey.Approaches() {
		calls := 0
		selected, err := approach.Filter(values, func(f fraction.Fraction) (bool, error) {
			calls++
			if calls == 2 {
				return true, stop
			}
			return true, nil
		})
		fmt.Println(approach.Name, selected, errors.Is(err, stop), calls)
	}
	// Output:
	// concrete [1/2] true 2
	// closure [1/2] true 2
	// interface [1/2] true 2
	// empty-interface [1/2] true 2
	// reflection [1/2] true 2
	// generation [1/2] true 2
	// generic [1/2] true 2
	// iterator [1/2] true 2
}

// Решение: одинаковые итоговые значения не означают одинаковые эффекты.
func Example_mapComposition() {
	var events []string
	f := func(n int) int { events = append(events, fmt.Sprintf("f%d", n)); return n + 1 }
	g := func(n int) int { events = append(events, fmt.Sprintf("g%d", n)); return n * 2 }
	first := generic.Map(generic.Map([]int{1, 2}, f), g)
	fmt.Println(first, events)
	events = nil
	second := generic.Map([]int{1, 2}, func(n int) int { return g(f(n)) })
	fmt.Println(second, events)
	// Output:
	// [4 6] [f1 f2 g2 g3]
	// [4 6] [f1 g2 f2 g3]
}

func Example_comparatorCounterexample() {
	a := fraction.Fraction{Numerator: 1, Denominator: 2}
	b := fraction.Fraction{Numerator: 2, Denominator: 4}
	wrongLess := a.Numerator < b.Numerator
	order, err := a.Compare(b)
	fmt.Println("По числителям:", wrongLess)
	fmt.Println("По значениям:", order, err)
	// Output:
	// По числителям: true
	// По значениям: 0 <nil>
}
