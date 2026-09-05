// Фрагмент 4: book/chapters/gp-09-constraints.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type Users []User
func (s Users) Count() int { return len(s) }

var users Users
copied := Clone(users)
_ = copied.Count() // возвращён Users, не просто []User
