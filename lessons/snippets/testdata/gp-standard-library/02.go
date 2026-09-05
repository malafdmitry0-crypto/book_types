// Фрагмент 2: book/chapters/gp-10-standard-library.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
// import "slices"
index := slices.IndexFunc(users, adult)
all := !slices.ContainsFunc(users, func(u User) bool {
    return !active(u)
})
