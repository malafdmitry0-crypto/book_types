package interfaces_test

import (
	"fmt"
	"gotypes/fraction"
	"gotypes/lessons/interfaces"
	"gotypes/money"
)

func Example() {
	f := fraction.Fraction{Numerator: 1, Denominator: 2}
	m := money.Money{Amount: 250, Currency: "USD"}
	interfaces.PrintValue(f)
	interfaces.PrintValue(m)
	fmt.Println(interfaces.DescribeStringer(f))
	text, err := interfaces.DescribeUnknown(f)
	fmt.Println(text, err)
	_, err = interfaces.DescribeUnknown(42)
	fmt.Println("У int нет String:", err != nil)
	var unknown interface{} = f
	stringer, ok := unknown.(interfaces.Stringer)
	fmt.Println("Подходит Stringer:", ok, interfaces.DescribeStringer(stringer))
	restored, ok := unknown.(fraction.Fraction)
	fmt.Println("Это Fraction:", ok, restored.Numerator, restored.Denominator)
	total, err := interfaces.AddFraction(f, f)
	fmt.Println("Сложение дробей:", total, err)
	// Output:
	// 1/2
	// 250 USD (minor units)
	// 1/2
	// 1/2 <nil>
	// У int нет String: true
	// Подходит Stringer: true 1/2
	// Это Fraction: true 1 2
	// Сложение дробей: 1 <nil>
}

func Example_sliceAndPointer() {
	fractions := []fraction.Fraction{{Numerator: 1, Denominator: 2}}
	printable := make([]interfaces.Stringer, len(fractions))
	for i, value := range fractions {
		printable[i] = value
	}
	interfaces.PrintAll(printable)
	printable[0] = money.Money{Currency: "USD"}
	fmt.Println("Исходная дробь:", fractions[0])
	var pointer *fraction.Fraction
	var value interfaces.Stringer = pointer
	fmt.Println("nil-указатель:", pointer == nil, "nil-интерфейс:", value == nil)
	_, ok := value.(fraction.Fraction)
	fmt.Println("Указатель не значение:", ok)
	// Output:
	// 1/2
	// Исходная дробь: 1/2
	// nil-указатель: true nil-интерфейс: false
	// Указатель не значение: false
}
