# 7. Интерфейсы: исходная модель полиморфизма

Код главы: [typeinterfaces.go](../../lessons/reference/typeinterfaces/typeinterfaces.go) · [example_test.go](../../lessons/reference/typeinterfaces/example_test.go).


> Цель главы: Работать с динамическим типом интерфейсного значения.

Здесь собраны точные правила, на которых основаны главы об адаптерах. Следите за двумя типами: объявленным типом переменной и конкретным типом значения внутри интерфейса.

## Упаковка конкретного значения

```go
var x any = int(42) // any доступен с Go 1.18
var old interface{} = int(42) // историческая запись того же типа
```

Интерфейс хранит динамический тип и значение. Присваивание `int` в `any` не превращает его в «универсальное число». Пара `тип + значение` — семантическая модель, не обещание конкретной раскладки runtime. См. [The Laws of Reflection](https://go.dev/blog/laws-of-reflection).

## Type assertion: `x.(T)`

```go
var x any = int(42)
n := x.(int)              // 42
wide, ok := x.(int64)      // 0, false
// panicValue := x.(int64) // panic

if n, ok := x.(int); ok {
    converted := int64(n)  // assertion, затем числовая конверсия
    _ = converted
}
```

Для конкретного `T` требуется совпадение динамического типа, а не просто возможность конверсии. Для интерфейсного `T` проверяется реализация интерфейса. Исходный операнд должен иметь интерфейсный тип; к обычной переменной `int` assertion применить нельзя. См. [Effective Go: interface conversions](https://go.dev/doc/effective_go#interface_conversions).

## Интерфейс → интерфейс

```go
// import "io"
func inspect(rw io.ReadWriter) {
    var reader io.Reader = rw // все нужные методы гарантированы статически
    writer, ok := reader.(io.Writer) // проверка динамического значения
    _, _ = writer, ok
}
```

## Type switch

```go
// import "fmt"
func describe(x any) string {
    switch v := x.(type) {
    case nil:
        return "nil interface"
    case int:
        return fmt.Sprintf("int: %d", v)
    case string:
        return "string: " + v
    case float32, float64:
        return fmt.Sprintf("float: %v", v) // здесь v имеет тип any
    default:
        return fmt.Sprintf("other: %T", v)
    }
}
```

Это проверка вариантов во время исполнения. Нетипизированного `x.(type)` вне type switch не существует. См. [Type switches](https://go.dev/ref/spec#Type_switches).

## Методы, указатели и typed nil

```go
type Worker interface { Work() }
type Job struct{}
func (*Job) Work() {}

var _ Worker = (*Job)(nil) // проверка реализации на этапе компиляции
// var _ Worker = Job{}   // у Job нет метода с pointer receiver

func nilExample() bool {
    var p *Job = nil
    var w Worker = p
    return w == nil // false: динамический тип *Job присутствует
}
```

Автоматическое взятие адреса при вызове метода у адресуемой переменной не добавляет pointer-методы в method set типа значения. Интерфейс равен `nil`, когда отсутствуют и динамический тип, и значение. Это также объясняет ошибку возврата typed nil как `error`. См. [Go FAQ: nil error](https://go.dev/doc/faq#nil_error) и [method sets](https://go.dev/ref/spec#Method_sets).

## Самопроверка

Почему интерфейс с nil-указателем не равен nil?

---

[← Определённые типы, структуры, функции и контейнеры](06-defined-types.md) · [Оглавление](../README.md) · [Reflection: когда тип приходит во время исполнения →](08-reflection.md)
