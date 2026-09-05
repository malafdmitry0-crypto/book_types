# 13. Развитие дженериков после Go 1.20

Код главы: [alias_go124.go](../../lessons/reference/modern/alias_go124.go) · [adder_go126.go](../../lessons/reference/modern/adder_go126.go) · [method_go127.go](../../lessons/reference/modern/method_go127.go).


> Цель главы: Читать современный generic-код с учётом версии языка.

Не все возможности современного Go доступны одной версии компилятора. У каждого примера ниже есть минимальная версия; новые файлы кода отделены условиями сборки.

## Вывод типов: Go 1.21

```go
func Identity[T any](x T) T { return x }

var intIdentity func(int) int = Identity // Go 1.21+
// До этого можно было явно написать Identity[int].
```

Вывод аргументов типа не преобразует значения. Он помогает компилятору определить нужную инстанциацию. Источник: [Go 1.21](https://go.dev/doc/go1.21#language).

## Generic aliases: Go 1.24

```go
type Set[T comparable] = map[T]bool // Go 1.24+

func aliasSetExample() {
    raw := map[string]bool{"go": true}
    var names Set[string] = raw // идентичный тип
    _ = names
}
```

Запись `type Set[T comparable] map[T]bool` без `=` создаёт определённый generic-тип. Алиас сохраняет идентичность типа справа. Эксперимент появился в Go 1.23, полноценная поддержка — в [Go 1.24](https://go.dev/doc/go1.24#language).

## Самоссылки в ограничениях: Go 1.26

```go
type Adder[A Adder[A]] interface { // Go 1.26+
    Add(A) A
}
```

Ограничение может ссылаться на объявляемый generic-тип. Это расширяет описание допустимых типов, но не даёт дополнительной runtime-конверсии. Источник: [Go 1.26](https://go.dev/doc/go1.26#language).

## Методы с собственными параметрами типов: Go 1.27

```go
type Converter struct{}

func (Converter) ToInt64[T ~int | ~int32 | ~int64](x T) int64 {
    return int64(x)
}

// n := (Converter{}).ToInt64(int32(42)) // Go 1.27+
```

До Go 1.27 такой алгоритм размещали в отдельной generic-функции пакета. Методы generic-типов, использующие параметры получателя, существовали с Go 1.18; новинка — собственные параметры метода.

Методы интерфейсов по-прежнему не могут объявлять параметры типов, и generic-метод не реализует метод интерфейса. В Go 1.27 также обобщён вывод типов при присваивании generic-функции или её конверсии к подходящему функциональному типу. Источник: [Go 1.27](https://go.dev/doc/go1.27#language).

## Самопроверка

В каких версиях появились generic aliases и собственные параметры типов у методов?

---

[← Срез → значение массива: Go 1.20](12-slice-array.md) · [Оглавление](../README.md) · [`unsafe`: интерпретация памяти →](14-unsafe.md)
