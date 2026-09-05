// Фрагмент 2: book/chapters/04-assignment.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type Numbers []int
var named Numbers = []int{1, 2}
var plain []int = named // одинаковые underlying types, один тип неименованный

type OtherNumbers []int
// var other OtherNumbers = named // оба типа определённые
other := OtherNumbers(named)       // явная конверсия допустима
