// Фрагмент 7: book/chapters/gp-early-map-sum.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type moneySum struct {
    Values []money.Money
    Total  money.Money
}

func (s *moneySum) Len() int { return len(s.Values) }
func (s *moneySum) Add(i int) error {
    next, err := s.Total.Add(s.Values[i])
    if err != nil {
        return err
    }
    s.Total = next
    return nil
}
