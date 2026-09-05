// Фрагмент 1: book/chapters/03-numeric.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
var n int = 42
wide := int64(n)
fraction := float64(n)

var f float64 = -3.9
whole := int(f) // -3: дробная часть отбрасывается к нулю

var big uint16 = 300
small := uint8(big) // 44: остаются младшие 8 бит

var negative int = -1
unsigned := uint8(negative) // 255

var z complex128 = 3 + 4i
compact := complex64(z)
realPart := real(z)
imagPart := imag(z)
constructed := complex(float64(n), 0)

// var bad int64 = n // разные типы: нужна int64(n)
// _ = int(true)    // bool → число не предусмотрено
// _ = bool(1)      // число → bool не предусмотрено
// _ = float64(z)   // сначала явно выбираем real(z) или imag(z)
