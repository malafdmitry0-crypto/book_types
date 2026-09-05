// Фрагмент 3: book/chapters/06-defined-types.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type A int
type B int

type F func(int) int
type G func(int) int

type Vector [2]int
type SliceA []int
type SliceB []int
type MapA map[string]int
type MapB map[string]int

func exampleComposite() {
    a := A(1)
    b := (*B)(&a) // допустимая конверсия неименованных указателей
    *b = 7       // изменяет a: это та же память

    f := F(func(n int) int { return n + 1 })
    g := G(f)
    vector := Vector([2]int{1, 2})

    s := SliceA{1, 2}
    t := SliceB(s) // тот же backing array
    t[0] = 9

    m := MapA{"x": 1}
    convertedMap := MapB(m) // та же map
    convertedMap["x"] = 2

    _, _, _ = g, vector, b
}
