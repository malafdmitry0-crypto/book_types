// Фрагмент 2: book/chapters/gp-early-contracts.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func ExampleMin() {
	values := []boxint.BoxInt{{Value: 3}, {Value: 1}, {Value: 2}, {Value: 1}}
	collection := boxes(values)

	index, err := algorithms.Min(collection)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	if index == -1 {
		fmt.Println("Коллекция пуста")
		return
	}
	fmt.Println("Индекс:", index)
	fmt.Println("Минимум", values[index])
	// Output:
	// Индекс: 1
	// Минимум 1
}
