// Фрагмент 1: book/chapters/05-strings.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
text := "Привет"
bytes := []byte(text) // UTF-8 байты
runes := []rune(text) // кодовые точки Unicode
fromBytes := string(bytes)
fromRunes := string(runes)

letter := string(rune(65)) // "A", не "65"
replacement := string(rune(-1)) // "�": недопустимая кодовая точка

type Label string
label := Label(text)
original := string(label)
