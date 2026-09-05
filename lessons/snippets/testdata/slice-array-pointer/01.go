// Фрагмент 1: book/chapters/10-slice-array-pointer.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
s := []int{10, 20, 30}
p := (*[2]int)(s)
p[0] = 99 // s[0] тоже стал 99

// _ = (*[4]int)(s) // panic: len(s) < 4
