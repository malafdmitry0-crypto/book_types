// Фрагмент 8: book/chapters/gp-early-map-sum.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
var _ algorithms.Summable = (*moneySum)(nil)
// moneySum{} не реализует Summable: нужные методы имеют pointer receiver.
