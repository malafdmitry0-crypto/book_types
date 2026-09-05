package closures_test

import (
	"fmt"
	"gotypes/algorithms/closures"
	"gotypes/fraction"
	"sort"
)

func ExampleMapIndex() {
	values := []fraction.Fraction{{Numerator: 2, Denominator: 4}, {Numerator: 6, Denominator: 3}}
	texts := make([]string, len(values))
	err := closures.MapIndex(len(values), func(i int) error {
		if err := values[i].Validate(); err != nil {
			return err
		}
		texts[i] = values[i].String()
		return nil
	})
	fmt.Println(texts, err)
	numbers := []int{2, 5, 9, 12}
	index := sort.Search(len(numbers), func(i int) bool { return numbers[i] >= 8 })
	fmt.Println("Бинарный поиск:", index)
	// Output:
	// [1/2 2] <nil>
	// Бинарный поиск: 2
}
