// Фрагмент 8: book/chapters/gp-type-inference.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
var f func(int) int = generic.Identity // Go 1.21+: T = int из типа переменной
doubled := generic.Map([]int{1, 2}, generic.Double[int])
inferred := generic.Map([]int{1, 2}, generic.Double) // Go 1.21+: T выведен из []int
