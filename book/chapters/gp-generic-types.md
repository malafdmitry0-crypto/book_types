# 16. Обобщённые типы данных: Stack[T], Set[K] и методы

Код главы: [containers.go](../../lessons/containers/containers.go) · [example_test.go](../../lessons/containers/example_test.go).


> Параметры типов у структуры входят в сигнатуру каждого её метода; ограничение выбирается один раз при объявлении типа, и метод не может его сузить.

До сих пор мы обобщали алгоритмы: `Find`, `Map`, `Clone` принимали `[]E` и возвращали результат, а сам контейнер оставался срезом. Стек, множество или очередь — не функция, а тип: его нужно объявить, хранить в поле, возвращать из конструктора. Go 1.18 разрешает параметры типов не только у функций, но и у объявлений типов. Посмотрим, что при этом переносится с функций без изменений, а что устроено иначе. Источник: [Go 1.18](https://go.dev/doc/go1.18#language).

## Stack[T]: тип элемента фиксируется при объявлении

```go
type Stack[T any] struct {
	items []T
}
```

`Stack` — не тип, а описание семейства типов. Конкретный тип появляется при инстанциации: `Stack[int]`, `Stack[fraction.Fraction]`. Методы объявляются для всего семейства сразу:

```go
func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}
```

```go
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	top := s.items[len(s.items)-1]
	var zero T
	s.items[len(s.items)-1] = zero
	s.items = s.items[:len(s.items)-1]
	return top, true
}
```

Получатель `*Stack[T]` повторяет параметры типа, но только по имени: ограничение `any` здесь не пишется, оно уже задано в объявлении типа. Имя можно выбрать другое, `func (s *Stack[E])`, но число и порядок параметров должны совпадать с объявлением. Добавить в получателе новое ограничение или новый параметр нельзя.

Перед сокращением среза `Pop` очищает освобождённую ячейку. Если T содержит указатели, оставшийся массив иначе продолжал бы удерживать снятое значение. Само копирование Stack по-прежнему копирует только заголовок среза: независимые стеки нельзя получить простым присваиванием после добавления элементов.

В `Pop` два решения, знакомые по итераторам. Значение «стек пуст» нельзя обозначить литералом: для `T` неизвестно, что такое «ноль», поэтому пишем `var zero T`. И нельзя использовать это нулевое значение как признак отсутствия: для `int` ноль — обычный элемент, а для `Fraction` нулевая структура невалидна, и `Pop` из пустого стека вернул бы `invalid fraction(0/0)`. Пара `(T, bool)` отделяет результат от его отсутствия.

```go
var s containers.Stack[fraction.Fraction] // нулевое значение пригодно
s.Push(fraction.Fraction{Numerator: 1, Denominator: 2})
s.Push(fraction.Fraction{Numerator: 1, Denominator: 3})
fmt.Println("Len:", s.Len(), "Stack:", s)
top, ok := s.Pop()
fmt.Println("Pop:", top, ok)
```

## Инстанциации — разные типы, нулевые значения — разные судьбы

`Stack[fraction.Fraction]` и `Stack[*fraction.Fraction]` — два несвязанных типа, как `[]Fraction` и `[]*Fraction`. Значение одного нельзя присвоить другому, метод `Push` первого не примет указатель:

```go
var values containers.Stack[fraction.Fraction]
var pointers containers.Stack[*fraction.Fraction]
half := fraction.Fraction{Numerator: 1, Denominator: 2}
values.Push(half)
pointers.Push(&half)
// values.Push(&half) // ошибка компиляции: *Fraction не является Fraction
// values = pointers  // ошибка компиляции: Stack[*Fraction] — другой тип
fmt.Printf("%T %T\n", values, pointers)
```

Нулевое значение `Stack[T]` пригодно к работе: внутри лежит nil-срез, а `append` к nil-срезу разрешён. Для множества это не так:

```go
type Set[K comparable] struct {
	m map[K]struct{}
}
```

```go
func NewSet[K comparable]() Set[K] {
	return Set[K]{m: make(map[K]struct{})}
}
```

Чтение из nil-карты возвращает нулевое значение, а запись вызывает panic. Поэтому `var s Set[string]` отвечает на `Has` и `Len`, но `Add` завершится ошибкой `assignment to entry in nil map`; см. `ExampleSet_zeroValue`. Это та же тема, что и с `SumValues`: нулевое значение не всегда нейтрально. Параметры типов не изменили правила для карт, они лишь спрятали карту за именем `Set[K]`, и договор «сначала NewSet» должен быть написан в документации типа.

## Set[K]: comparable недостаточно для сортировки

```go
func (s Set[K]) Add(v K) { s.m[v] = struct{}{} }
```

```go
func (s Set[K]) Has(v K) bool {
	_, ok := s.m[v]
	return ok
}
```

Ключ карты обязан быть сравнимым, поэтому у `Set` ограничение `comparable`. Захотим вернуть элементы по возрастанию. Метод `func (s Set[K]) Sorted() []K` написать нельзя: внутри метода о `K` известно только то, что разрешено ограничением типа, а `comparable` не даёт `<`. Компилятор отвечает `K does not satisfy cmp.Ordered`; этот отказ зафиксирован в [method_constraint.go](../../lessons/failures/testdata/method_constraint.go). Выход — функция пакета с более сильным ограничением:

```go
func SortedKeys[K cmp.Ordered](s Set[K]) []K {
	keys := make([]K, 0, len(s.m))
	for k := range s.m {
		keys = append(keys, k)
	}
	slices.Sort(keys)
	return keys
}
```

`SortedKeys` принимает не любое `Set[K]`, а только такое, где `K` упорядочен; `Set[[2]int]` в неё передать нельзя, хотя множество из массивов существует. Так устроена и стандартная библиотека: у срезов нет метода `Sort`, есть функция `slices.Sort[S ~[]E, E cmp.Ordered]`. Методы generic-типа описывают операции, доступные при любом допустимом аргументе типа; всё, что требует большего, живёт рядом в функциях пакета. Пакеты `cmp` и `slices` появились в Go 1.21. Источник: [Go 1.21](https://go.dev/doc/go1.21#library).

## Методы и параметры типов до и после Go 1.27

Ограничения — не единственное, чего метод не может добавить. До Go 1.27 метод не может объявить собственных параметров типов: запись `func (s Stack[T]) Map[R any](f func(T) R) Stack[R]` отвергается на этапе разбора с сообщением `method must have no type parameters`; см. [generic_method.go](../../lessons/failures/testdata/generic_method.go). Преобразование `T → R` приходится оформлять функцией пакета:

```go
func MapStack[T, R any](s Stack[T], f func(T) R) Stack[R] {
	out := Stack[R]{items: make([]R, len(s.items))}
	for i, item := range s.items {
		out.items[i] = f(item)
	}
	return out
}
```

```go
texts := containers.MapStack(s, fraction.Fraction.String)
fmt.Printf("%T %v\n", texts, texts) // containers.Stack[string] [1/2 1/2]
```

Go 1.27 разрешает методам собственные параметры типов, но методы интерфейсов по-прежнему объявлять их не могут, и такой метод не реализует метод интерфейса. Модуль учебника требует Go 1.23, поэтому `containers` использует переносимую форму `MapStack`; вариант для нового компилятора вынесен в [method_go127.go](../../lessons/reference/modern/method_go127.go) под условием сборки. Источник: [Go 1.27](https://go.dev/doc/go1.27#language).

## Обобщённый тип и интерфейсы

Инстанцированный тип — обычный тип, и он реализует обычные интерфейсы обычным способом:

```go
func (s Stack[T]) String() string { return fmt.Sprint(s.items) }
```

`fmt.Sprint` принимает любое значение, поэтому `Stack[T]` при любом `T` удовлетворяет `fmt.Stringer`. Интерфейс тоже может иметь параметры типов:

```go
type Container[T any] interface {
	Len() int
	Push(T)
}
```

```go
func Fill[T any](c Container[T], values ...T) {
	for _, v := range values {
		c.Push(v)
	}
}
```

Здесь важны множества методов. `Push` объявлен на `*Stack[T]`, поэтому `Container[int]` реализует только `*Stack[int]`; `String` объявлен на значении, и `fmt.Stringer` реализуют оба типа. Правило то же, что и для необобщённых типов; параметр типа его не отменяет.

```go
var s containers.Stack[int]
containers.Fill(&s, 1, 2, 3)
// containers.Fill(s, 4) // ошибка компиляции: Stack[int] не реализует Container[int]
var _ fmt.Stringer = s // String объявлен на значении, подходят и Stack[int], и *Stack[int]
```

`Container[T]` — интерфейс с параметром, а не ограничение: значение `*Stack[int]` действительно кладётся в интерфейсную переменную, и вызов `Push` внутри `Fill` идёт через таблицу методов. Для `Fill` этого достаточно; если бы алгоритму был нужен конкретный тип контейнера, мы написали бы `Fill[C Container[T], T any](c C, ...)`, и тогда `Container[T]` стало бы ограничением.

## Проверка понимания

Почему `SortedKeys` нельзя сделать методом `Set[K]`, и какие два способа получить упорядоченные элементы остаются: функция пакета с ограничением `cmp.Ordered` или отдельный тип `OrderedSet[K cmp.Ordered]`? В чём различие для вызывающего, у которого уже есть `Set[[2]int]`?

---

[← Ограничения: выражаем ровно те операции, которые нужны](gp-09-constraints.md) · [Оглавление](../README.md) · [Вывод типов: когда аргументы типов можно не писать →](gp-type-inference.md)
