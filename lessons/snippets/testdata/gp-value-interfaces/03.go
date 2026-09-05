// Фрагмент 3: book/chapters/gp-value-interfaces.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
f := fraction.Fraction{Numerator: 1, Denominator: 2}
var value Stringer = f
text := value.String() // Разрешено контрактом Stringer.
// value.Numerator     // Не компилируется: поле не доступно через Stringer.
// value.Add(f)        // Не компилируется: Stringer не объявляет Add.
