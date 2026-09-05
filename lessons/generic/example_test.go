package generic_test

import (
	"fmt"
	"gotypes/fraction"
	"gotypes/lessons/generic"
	"gotypes/money"
)

type fractions []fraction.Fraction

func (s fractions) Count() int { return len(s) }

type person struct{ name string }

func (p person) Name() string { return p.name }

func Example() {
	values := fractions{{Numerator: 1, Denominator: 2}, {Numerator: 1, Denominator: 3}}
	half := fraction.Fraction{Numerator: 2, Denominator: 4}
	fmt.Println("Равенство полей:", generic.Contains(values, half))
	fmt.Println("Равенство дробей:", generic.ContainsEqual(values, half))
	fmt.Println("Тип среза сохранён:", generic.Clone(values).Count())
	fmt.Println("Имена:", generic.Names([]person{{"Аня"}, {"Борис"}}))
	total, err := generic.SumValues(values, fraction.Fraction{Denominator: 1})
	fmt.Println("Дроби:", total, err)
	cash, err := generic.SumValues([]money.Money{{Amount: 100, Currency: "USD"}, {Amount: 250, Currency: "USD"}}, money.Money{Currency: "USD"})
	fmt.Println("Деньги:", cash, err)
	texts, err := generic.MapValues(values, func(f fraction.Fraction) (string, error) { return f.String(), f.Validate() })
	fmt.Println("Строки:", texts, err)
	fmt.Println("Find:", generic.Find([]string{"go", "generic"}, func(s string) bool { return len(s) > 2 }))
	fmt.Println("All:", generic.All(values, func(f fraction.Fraction) bool { return !f.IsZero() }))
	fmt.Println("Map:", generic.Map([]int{1, 2}, func(n int) int { return n * 2 }))
	// Output:
	// Равенство полей: false
	// Равенство дробей: true
	// Тип среза сохранён: 2
	// Имена: [Аня Борис]
	// Дроби: 5/6 <nil>
	// Деньги: 350 USD (minor units) <nil>
	// Строки: [1/2 1/3] <nil>
	// Find: 1
	// All: true
	// Map: [2 4]
}
