// Фрагмент 4: book/chapters/gp-early-search.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func (q fractionQuery) Match(i int) (bool, error) {
    value := q.Values[i]           // Получаем конкретную Fraction.
    matched, err := q.Test(value)  // Вызываем сохранённую функцию.
    return matched, err           // Возвращаем ответ алгоритму.
}
