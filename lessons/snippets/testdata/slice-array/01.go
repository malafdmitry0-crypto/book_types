// Фрагмент 1: book/chapters/12-slice-array.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
s := []int{10, 20, 30}
a := [2]int(s)
a[0] = 99 // s[0] остаётся 10

// _ = [4]int(s) // panic: len(s) < 4
