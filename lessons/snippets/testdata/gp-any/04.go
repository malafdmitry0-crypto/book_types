// Фрагмент 4: book/chapters/gp-05-any.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
var boxed []interface{}
if users != nil {
    boxed = make([]interface{}, len(users))
}
for i, user := range users {
    boxed[i] = user
}

index := FindAny(boxed, func(x interface{}) bool {
    return x.(User).Age >= 18
})
all := AllAny(boxed, func(x interface{}) bool {
    return x.(User).Active
})
raw := MapAny(boxed, func(x interface{}) interface{} {
    return x.(User).Name
})

var names []string
if raw != nil {
    names = make([]string, len(raw))
}
for i, item := range raw {
    names[i] = item.(string)
}
