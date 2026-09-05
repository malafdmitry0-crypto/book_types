// Фрагмент 2: book/chapters/05-strings.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
// import "strconv"
n, err := strconv.Atoi("42")
if err != nil {
    // Обработать ошибку формата или диапазона.
}
decimal := strconv.Itoa(42)
hex := strconv.FormatInt(255, 16) // "ff"
parsed, parseErr := strconv.ParseInt("ff", 16, 64)
flag, boolErr := strconv.ParseBool("true")
number, floatErr := strconv.ParseFloat("3.14", 64)
