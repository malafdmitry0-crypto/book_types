package laws_test

import (
	"fmt"
	"gotypes/fraction"
	"math"
)

func Example() {
	half := fraction.Fraction{Numerator: 1, Denominator: 2}
	other := fraction.Fraction{Numerator: 2, Denominator: 4}
	fmt.Println("Поля и значение:", half == other, half.Equal(other))
	invalid := fraction.Fraction{}
	fmt.Println("Невалидная дробь равна себе:", invalid.Equal(invalid))
	zero := fraction.Fraction{Denominator: 1}
	sum, err := half.Add(zero)
	fmt.Println("Нейтральный элемент:", sum.Equal(half), err)
	a := fraction.Fraction{Numerator: math.MaxInt64, Denominator: 1}
	b := fraction.Fraction{Numerator: 1, Denominator: 1}
	c := fraction.Fraction{Numerator: -1, Denominator: 1}
	_, leftErr := a.Add(b)
	bc, err := b.Add(c)
	if err != nil {
		fmt.Println(err)
		return
	}
	right, err := a.Add(bc)
	fmt.Println("(a+b)+c: первый шаг переполнен:", leftErr != nil)
	fmt.Println("a+(b+c):", right, err)
	// Output:
	// Поля и значение: false true
	// Невалидная дробь равна себе: false
	// Нейтральный элемент: true <nil>
	// (a+b)+c: первый шаг переполнен: true
	// a+(b+c): 9223372036854775807 <nil>
}
