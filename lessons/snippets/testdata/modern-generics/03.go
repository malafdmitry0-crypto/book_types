// Фрагмент 3: book/chapters/13-modern-generics.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type Adder[A Adder[A]] interface { // Go 1.26+
    Add(A) A
}
