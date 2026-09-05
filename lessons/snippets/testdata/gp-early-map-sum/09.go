// Фрагмент 9: book/chapters/gp-early-map-sum.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func ExampleSum() {
	values := []money.Money{
		{Amount: 100, Currency: "USD"},
		{Amount: 250, Currency: "USD"},
	}
	initial := money.Money{Currency: "USD"} // Нулевая сумма с заданной валютой.
	sum := moneySum{Values: values, Total: initial}

	// Указатель позволяет адаптеру изменять своё поле Total.
	err := algorithms.Sum(&sum)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Сумма:", sum.Total)
	// Output: Сумма: 350 USD (minor units)
}
