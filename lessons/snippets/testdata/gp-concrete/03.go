// Фрагмент 3: book/chapters/gp-02-concrete.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
func SumMoney(values []money.Money, initial money.Money) (money.Money, error) {
    total := initial
    for _, value := range values {
        next, err := total.Add(value)
        if err != nil {
            return total, err // в этом варианте возвращаем уже накопленный префикс
        }
        total = next
    }
    return total, nil
}
