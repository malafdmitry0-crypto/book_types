# 6. Определённые типы, структуры, функции и контейнеры

Код главы: [defined.go](../../lessons/reference/defined/defined.go) · [example_test.go](../../lessons/reference/defined/example_test.go) · [example_test.go](../../lessons/reference/basics/example_test.go).


> Цель главы: Понимать совместимость составных и определённых типов.

Конверсия нового типа не означает преобразование всех его элементов и не переносит методы. Начнём с ID, затем проверим структуры и контейнеры, обращая внимание на общую память.

## Новый тип не наследует методы

```go
// import "fmt"
type UserID int64
func (id UserID) String() string { return fmt.Sprintf("user:%d", id) }

type OrderID UserID

func exampleIDs() {
    user := UserID(42)
    order := OrderID(user)
    _ = user.String()
    // _ = order.String() // у OrderID нет этого метода
    _ = order
}
```

Конверсия позволяет перейти между такими типами, но не проверяет предметный смысл. Для запрета отрицательного ID нужен конструктор с валидацией.

## Структуры; послабление Go 1.8 для тегов

```go
type APIUser struct {
    ID int64 `json:"id"`
}
type DBUser struct {
    ID int64 `db:"user_id"`
}

func convertUser(api APIUser) DBUser {
    return DBUser(api) // Go 1.8+: различия тегов не мешают
}
```

Это не автоматический маппинг по названиям. Порядок полей и их типы имеют значение. Например, `ID int` и `ID int64` потребуют явного построения результата. Копирование структуры неглубокое: поля-срезы и указатели могут продолжить ссылаться на общие данные. См. [Go 1.8](https://go.dev/doc/go1.8#language).

## Указатели, функции, массивы, срезы и map

```go
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
```

Одинаковое представление памяти само по себе не разрешает конверсию. Например, `*int64 → *float64` требует иных механизмов. Для обычных указателей существенны условия идентичности underlying types базовых типов; для функций — сигнатуры. См. [Conversions](https://go.dev/ref/spec#Conversions).

Нет поэлементной конверсии контейнеров:

```go
src := []int{1, 2, 3}
// dst := []int64(src) // запрещено
dst := make([]int64, len(src))
for i, n := range src {
    dst[i] = int64(n)
}

// Не работают также []string → []any и map[string]int → map[string]int64.
```

## Направление канала

```go
ch := make(chan int)
var recv <-chan int = ch
var send chan<- int = ch
explicit := (<-chan int)(ch)
// both := (chan int)(recv) // восстановить оба направления нельзя
```

Это ограничивает доступ через значение, а не создаёт новый канал. Для двух определённых типов каналов может потребоваться промежуточный неименованный тип. Правила приведены в [Assignability](https://go.dev/ref/spec#Assignability).

## Самопроверка

Какие данные остаются общими после конверсии среза, map и структуры?

---

[← Строки, байты, руны и числа](05-strings.md) · [Оглавление](../README.md) · [Интерфейсы: исходная модель полиморфизма →](07-interfaces.md)
