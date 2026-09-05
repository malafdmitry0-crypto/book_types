// Фрагмент 2: book/chapters/06-defined-types.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type APIUser struct {
    ID int64 `json:"id"`
}
type DBUser struct {
    ID int64 `db:"user_id"`
}

func convertUser(api APIUser) DBUser {
    return DBUser(api) // Go 1.8+: различия тегов не мешают
}
