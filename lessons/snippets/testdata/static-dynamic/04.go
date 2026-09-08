// Фрагмент 4: book/chapters/21-static-dynamic.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type Speaker interface {
    Speak() string
}

type Cat struct{}
func (Cat) Speak() string { return "meow" }

type Robot struct{}
func (Robot) Speak() string { return "beep" }

func say(s Speaker) {
    fmt.Println(s.Speak())
}
