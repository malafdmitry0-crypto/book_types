# 11. Дженерики: Go 1.18

> Цель главы: Обобщать конверсии с проверяемыми ограничениями.

### 11.1. Что они добавили

До дженериков обобщённый алгоритм обычно выражали через интерфейсы с методами, `interface{}` с assertions, reflection, генерацию кода или отдельные реализации для типов. Дженерики позволили проверять операции над параметрами типов при компиляции и сохранять конкретные типы аргументов и результатов. См. [An Introduction To Generics](https://go.dev/blog/intro-generics).

### 11.2. Ограниченная конверсия

```go
type Real interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
    ~float32 | ~float64
}

func Convert[To Real, From Real](x From) To {
    return To(x)
}

// result := Convert[int64](int32(42))
// To задан явно; From выводится из аргумента.
```

`~int64` допускает типы с underlying type `int64`, например `type UserID int64`. Без `~` член `int64` описывает сам этот тип. `|` объединяет множества типов. Здесь конверсия разрешена для всех сочетаний исходного и целевого типов, но потери точности и переполнение остаются возможными.

Интерфейс с type terms, например `interface{ ~int | ~string }`, применяется как ограничение, а не как тип обычной переменной. `any` означает отсутствие ограничений, а не разрешение любой операции. Источник: [Go 1.18](https://go.dev/doc/go1.18#language).

### 11.3. Почему универсальный cast не получается

```go
// Не компилируется: any не гарантирует допустимость конверсии.
// func Cast[To, From any](x From) To { return To(x) }

func Assert[T any](x any) (T, bool) {
    value, ok := x.(T)
    return value, ok
}

func IsString[T any](x T) bool {
    _, ok := any(x).(string)
    return ok
}
```

`Assert[int64](int(42))` вернёт `0, false`. `Assert` обобщает проверку типа, не числовую конверсию. Прямой `x.(string)` для `x T` запрещён: сначала нужно интерфейсное значение. Так же для type switch используют `switch any(x).(type)`. См. [Type assertions](https://go.dev/ref/spec#Type_assertions).

### 11.4. Поэлементный маппинг

```go
func Map[S ~[]E, E, R any](src S, convert func(E) R) []R {
    if src == nil {
        return nil
    }
    dst := make([]R, len(src))
    for i, item := range src {
        dst[i] = convert(item)
    }
    return dst
}

// result := Map([]int{1, 2}, func(n int) int64 { return int64(n) })
```

Алгоритм создаёт новый срез, а смысл конверсии задаёт вызывающий код. При возможных ошибках нужен отдельный контракт с `(R, error)` и правилом обработки частичного результата. Дженерики сами такой контракт не выбирают.

### 11.5. `comparable`: уточнение Go 1.20

С Go 1.20 `any` может удовлетворять ограничению `comparable`. Это не обещает, что сравнение любых динамических значений безопасно:

```go
func Equal[T comparable](a, b T) bool { return a == b }

// Equal[any](42, 42)           // true, Go 1.20+
// Equal[any]([]int{1}, []int{1}) // panic: срезы несравнимы
```

Это изменение правил удовлетворения ограничения, а не новый способ конверсии. Источник: [Go 1.20](https://go.dev/doc/go1.20#language).

## Самопроверка

Почему функция с параметрами To any и From any не может просто вернуть To(x)?

---

[← Срез → указатель на массив: Go 1.17](10-slice-array-pointer.md) · [Оглавление](../README.md) · [Срез → значение массива: Go 1.20 →](12-slice-array.md)
