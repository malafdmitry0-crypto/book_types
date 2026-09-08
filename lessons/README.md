# Код учебника

Начните с Example в нужном пакете. Каждый пример содержит данные, вызов и ожидаемый вывод в комментарии Output. Реализации находятся рядом либо в algorithms. `go test` проверяет вывод автоматически; при успехе показывает статус, а не stdout.

## Маршрут основного курса

| Тема | Реализация и запуск |
|---|---|
| Одна операция у разных значений | [interfaces](interfaces/example_test.go): `go test ./lessons/interfaces -v` |
| Конкретные Find, Map и Sum | [concrete](concrete/example_test.go): `go test ./lessons/concrete -v` |
| Индексы и замыкания | [closures](../algorithms/closures/example_test.go): `go test ./algorithms/closures -v` |
| Адаптеры и все восемь алгоритмов | [algorithms](../algorithms/example_test.go): `go test ./algorithms -v` |
| Пустой интерфейс, reflection, генерация, generics, iter | [сценарии второй части](../examples/example_test.go): `go -C examples test ./... -v` |
| Ограничения, дроби, деньги и ошибки generic API | [generic](generic/example_test.go): `go test ./lessons/generic -v` |

Во второй части используется дополнительная модель User с callback без ошибок. Она позволяет сравнить сигнатуры восьми подходов без арифметики. Сценарии с Fraction и Money сохраняют ошибки и находятся в основном модуле. Это разные договоры API; их нельзя незаметно смешивать.

## Продолжение: контейнеры и вывод типов

[containers](containers/example_test.go) показывает Stack, Set и методы обобщённых типов; [inference](generic/inference_example_test.go) — явные и выводимые аргументы типов. Запуск: `go test ./lessons/containers ./lessons/generic`.

## Одна задача во всех подходах

[journey](journey/README.md) — одинаковые Fraction, Find, Sum и Filter во всех восьми реализациях; Reduce сохраняет тот же договор об ошибках. [laws](laws/README.md) — свойства равенства, порядка и сложения. Здесь же лежат решения практикума, fuzz-тесты и реальные benchmark-результаты.

```sh
go test ./lessons/journey ./lessons/laws -v
```

## Справочник

| Пакет | Что запускать и наблюдать |
|---|---|
| [typing](reference/typing/example_test.go) | Статический тип, динамический тип интерфейса, assertions и выбор метода |
| [cppmemory](reference/cppmemory/example_test.go) | Go глазами C++: копирование, указатели, интерфейсы и nil |
| [basics](reference/basics/example_test.go) | Числа, константы, строки, парсинг, каналы, массивы, unsafe |
| [defined](reference/defined/example_test.go) | Новые типы, теги структур, функции и общая память |
| [typeinterfaces](reference/typeinterfaces/example_test.go) | Assertions, type switch, методы, typed nil |
| [reflection](reference/reflection/example_test.go) | Динамическая конверсия и её отказ, SetInt, TypeFor |
| [aliases](reference/aliases/example_test.go) | Алиас и определённый тип |
| [conversions](reference/conversions/example_test.go) | Generic-конверсии, assertions, Map и comparable |
| [modern](reference/modern/example_test.go) | Вывод типов и версии новых возможностей |
| [bridges](reference/bridges/example_test.go) | errors.As и http.HandlerFunc без запуска сервера |

```sh
go test ./lessons/reference/... -v
```

## Ошибки — тоже примеры

[failures/testdata](failures/testdata/cases.json) содержит самостоятельные неверные программы. [Тест](failures/failures_test.go) компилирует каждую отдельно и проверяет причину отказа. Он также проверяет ожидаемые паники: assertion, короткий срез, несравнимые интерфейсные значения и nil callback. Текст диагностики проверяется по устойчивому фрагменту, не целиком.

```sh
go test ./lessons/failures -v
```

## Каждый Go-блок из Markdown есть в коде

[Манифест выдержек](snippets/manifest.json) связывает каждый Go-блок с точной копией в `.go` и с полными исходниками соответствующей главы. Выдержки в `snippets/testdata` сохраняют учебный контекст, в том числе неполные и намеренно неверные строки. Их не нужно запускать по одному: для запуска используйте полные сценарии выше. Они не выданы за самостоятельные программы.

После изменения Markdown:

```sh
node scripts/sync-snippets.mjs
node scripts/sync-snippets.mjs --check
node scripts/build.mjs
```

## Версии и общая проверка

Основной модуль требует Go 1.23 или новее. Примеры generic aliases включаются с Go 1.24, самоссылок — с Go 1.26, generic-методов — с Go 1.27. Файлы используют условия `//go:build go1.N`; старый компилятор их исключает. На Go 1.25.4 последние два сценария не выполняются — наличие их исходников не означает прохождения тестов новой версией.

```sh
go test ./...
go vet ./...
go -C examples test ./...
go -C examples vet ./...
node scripts/sync-snippets.mjs --check
```

Отдельный модуль examples не входит в корневой `go test ./...`, поэтому для него нужна вторая команда. Запуск генератора: `go -C examples generate ./...`.
