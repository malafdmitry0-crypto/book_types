package closures_test

import (
	"fmt"

	"gotypes/algorithms/closures"
	"gotypes/fraction"
	"gotypes/money"
)

func ExampleFindIndex() {
	// Данные и предметное условие такие же, как в примере с fractionQuery.
	values := []fraction.Fraction{
		{Numerator: 1, Denominator: 3},
		{Numerator: 2, Denominator: 4},
		{Numerator: 3, Denominator: 4},
	}
	half := fraction.Fraction{Numerator: 1, Denominator: 2}
	isHalf := func(value fraction.Fraction) (bool, error) {
		if err := value.Validate(); err != nil {
			return false, err
		}
		return value.Equal(half), nil
	}

	// Вместо структуры с Len и Match передаём длину и замыкание.
	// Замыкание знает values и переводит индекс в конкретную дробь.
	index, err := closures.FindIndex(len(values), func(i int) (bool, error) {
		return isHalf(values[i])
	})
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	if index == -1 {
		fmt.Println("Совпадений нет")
		return
	}

	fmt.Println("Индекс:", index)
	fmt.Println("Дробь:", values[index])
	// Output:
	// Индекс: 1
	// Дробь: 1/2
}

func ExampleFindIndex_money() {
	values := []money.Money{
		{Amount: 100, Currency: "EUR"},
		{Amount: 250, Currency: "USD"},
	}

	// Тот же FindIndex работает с деньгами. Новый тип адаптера не нужен.
	index, err := closures.FindIndex(len(values), func(i int) (bool, error) {
		value := values[i]
		if err := value.Validate(); err != nil {
			return false, err
		}
		return value.Currency == "USD", nil
	})
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	if index == -1 {
		fmt.Println("Совпадений нет")
		return
	}

	fmt.Println("Индекс:", index)
	fmt.Println("Деньги:", values[index])
	// Output:
	// Индекс: 1
	// Деньги: 250 USD (minor units)
}
