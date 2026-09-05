package concrete_test

import (
	"fmt"
	"gotypes/fraction"
	"gotypes/lessons/concrete"
	"gotypes/money"
)

func Example() {
	values := []fraction.Fraction{{Numerator: 1, Denominator: 3}, {Numerator: 2, Denominator: 4}}
	half := fraction.Fraction{Numerator: 1, Denominator: 2}
	index, err := concrete.FindFraction(values, half)
	fmt.Println("По значению:", index, err)
	index, err = concrete.FindFractionFunc(values, func(value fraction.Fraction) (bool, error) {
		if err := value.Validate(); err != nil {
			return false, err
		}
		return value.Equal(half), nil
	})
	fmt.Println("По условию:", index, err)
	texts, err := concrete.MapFractionText(values, func(value fraction.Fraction) (string, error) { return value.String(), value.Validate() })
	fmt.Println("Строки:", texts, err)
	total, err := concrete.SumMoney([]money.Money{{Amount: 100, Currency: "USD"}, {Amount: 250, Currency: "USD"}}, money.Money{Currency: "USD"})
	fmt.Println("Сумма:", total, err)
	// Output:
	// По значению: 1 <nil>
	// По условию: 1 <nil>
	// Строки: [1/3 1/2] <nil>
	// Сумма: 350 USD (minor units) <nil>
}
