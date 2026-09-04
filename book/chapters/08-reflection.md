# 8. Reflection: когда тип приходит во время исполнения

> Цель главы: Проверять возможность динамической конверсии.

```go
// import "reflect"
func convertDynamic(x any, target reflect.Type) (any, bool) {
    value := reflect.ValueOf(x)
    if target == nil || !value.IsValid() || !value.CanConvert(target) {
        return nil, false
    }
    return value.Convert(target).Interface(), true
}

// out, ok := convertDynamic(int32(42), reflect.TypeOf(int64(0)))
// out имеет статический тип any и динамический тип int64.
```

`reflect.Type.ConvertibleTo` проверяет совместимость типов. `Value.CanConvert` учитывает также конкретное значение: например, длину среза перед конверсией в массив. `Convert` без проверки может вызвать panic.

`AssignableTo` относится к присваиванию, `Implements` — к реализации интерфейса. Эти проверки не взаимозаменяемы. `reflect.Value.Int()` читает целочисленное значение как `int64`, но не меняет исходный тип. Запись через `Set`/`SetInt` дополнительно зависит от доступности значения для изменения.

С Go 1.22 `reflect.TypeFor[T]()` возвращает тип аргумента, включая интерфейсный тип; `reflect.TypeOf(x)` возвращает динамический тип значения. Источники: [reflect](https://pkg.go.dev/reflect), [Go 1.22](https://go.dev/doc/go1.22#reflect).

## Самопроверка

Почему ConvertibleTo недостаточно для проверки конкретного среза перед конверсией в массив?

---

[← Интерфейсы: исходная модель полиморфизма](07-interfaces.md) · [Оглавление](../README.md) · [Алиасы: Go 1.9 →](09-aliases.md)
