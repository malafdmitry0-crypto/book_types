# 9. Алиасы: Go 1.9

Код главы: [aliases.go](../../lessons/reference/aliases/aliases.go) · [example_test.go](../../lessons/reference/aliases/example_test.go).


> Цель главы: Отличать новое имя типа от нового типа.

Сравним новое имя существующего типа и объявление нового типа. Внешне строки похожи, но необходимость конверсии зависит от единственного знака =.

```go
type UserID = int64 // тот же тип
type OrderID int64 // отдельный тип

func aliasExample() {
    var n int64 = 42
    var user UserID = n // конверсия не нужна
    order := OrderID(n)
    _, _ = user, order
}
```

Алиасы помогают переносить API между пакетами без создания нового несовместимого типа. Они не преобразуют значение и не добавляют предметную валидацию. Предопределённые `byte` и `rune` существовали до пользовательского синтаксиса алиасов. Источник: [Go 1.9](https://go.dev/doc/go1.9#language).

## Самопроверка

Чем type ID = int64 отличается от type ID int64?

---

[← Reflection: когда тип приходит во время исполнения](08-reflection.md) · [Оглавление](../README.md) · [Срез → указатель на массив: Go 1.17 →](10-slice-array-pointer.md)
