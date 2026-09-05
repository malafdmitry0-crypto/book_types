# 5. Строки, байты, руны и числа

Код главы: [example_test.go](../../lessons/reference/basics/example_test.go).


> Цель главы: Выбирать между байтами, рунами и текстовым представлением числа.

Строка может быть текстом числа или последовательностью UTF-8. В первом случае нужен парсер, во втором — выбор между байтами и рунами. Примеры ниже показывают разные результаты этих операций.

```go
text := "Привет"
bytes := []byte(text) // UTF-8 байты
runes := []rune(text) // кодовые точки Unicode
fromBytes := string(bytes)
fromRunes := string(runes)

letter := string(rune(65)) // "A", не "65"
replacement := string(rune(-1)) // "�": недопустимая кодовая точка

type Label string
label := Label(text)
original := string(label)
```

`byte` — алиас `uint8`, `rune` — алиас `int32`. Одна руна не обязательно равна одному видимому символу. `[]byte` сохраняет байты, а `[]rune` декодирует UTF-8; некорректные последовательности при декодировании заменяются на U+FFFD.

Обычные конверсии между строкой и срезом обеспечивают независимость наблюдаемых изменений: изменение полученного `[]byte` не меняет исходную строку. Из этого не следует обязательная heap-аллокация: компилятор может оптимизировать её. См. [Effective Go: strings](https://go.dev/doc/effective_go#strings) и [конверсии строк](https://go.dev/ref/spec#Conversions_to_and_from_a_string_type).

### Текстовое представление: `strconv`

```go
// import "strconv"
n, err := strconv.Atoi("42")
if err != nil {
    // Обработать ошибку формата или диапазона.
}
decimal := strconv.Itoa(42)
hex := strconv.FormatInt(255, 16) // "ff"
parsed, parseErr := strconv.ParseInt("ff", 16, 64)
flag, boolErr := strconv.ParseBool("true")
number, floatErr := strconv.ParseFloat("3.14", 64)
```

`int("42")` запрещено, `string(42)` даёт кодовую точку, а `strconv.Itoa(42)` — десятичную запись. Для форматирования с шаблоном есть `fmt.Sprintf`, для чтения формата — `fmt.Sscanf`. Проверка синтаксиса и диапазона относится к разбору данных, а не к конверсии языка. Справочник: [strconv](https://pkg.go.dev/strconv).

## Самопроверка

Какие результаты дадут string(rune(65)) и strconv.Itoa(65)?

---

[← Константы и присваивание](04-assignment.md) · [Оглавление](../README.md) · [Определённые типы, структуры, функции и контейнеры →](06-defined-types.md)
