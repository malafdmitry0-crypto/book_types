// Фрагмент 8: book/chapters/gp-value-interfaces.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
text, err := DescribeUnknown(42) // Компилируется, но вернёт ошибку выполнения.
// DescribeStringer(42)          // Не компилируется: у int нет String() string.
