// Фрагмент 9: book/chapters/gp-value-interfaces.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
var unknown interface{} = fraction.Fraction{Numerator: 1, Denominator: 2}
// DescribeStringer(unknown) // Не компилируется: interface{} не обещает String.

stringer, ok := unknown.(Stringer)
if ok {
    fmt.Println(DescribeStringer(stringer))
}
