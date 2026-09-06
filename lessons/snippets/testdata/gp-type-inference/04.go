// Фрагмент 4: book/chapters/gp-type-inference.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
n := 42
wide := generic.Convert[int64](n) // To задан явно, From = int выведен
