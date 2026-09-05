# Одна задача на Fraction

Восемь реализаций Find, Filter и Reduce. Sum — специализация соответствующего Reduce через Fraction.Add. [Восемь прямых вызовов](direct_example_test.go) · [сводный Example](example_test.go), затем откройте [подключение подходов](approaches.go).

| Файл | Что объясняет |
|---|---|
| [concrete.go](concrete.go) | Конкретные сигнатуры для Fraction |
| [closures.go](closures.go) | Обход по индексам, данные и состояние в замыкании |
| [interfaces.go](interfaces.go) | Search, Selection и Fold связывают операции с данными |
| [any.go](any.go) | Срез интерфейсов; упаковка и проверки находятся в approaches.go |
| [reflection.go](reflection.go) | Проверка сигнатур и динамические вызовы |
| [generated.go](generated.go) | Настоящий результат шаблонной генерации |
| [generic.go](generic.go) | Связь типа элемента E и аккумулятора A |
| [iterator.go](iterator.go) | Постепенная доставка, остановка источника и ошибка |

Все варианты сравниваются на подходящих callback и неизменяемых во время обхода данных. Find возвращает позицию либо -1. Filter сохраняет порядок и nilness, создаёт отдельное хранилище элементов. Reduce возвращает последний успешный аккумулятор. При ошибке текущий результат не фиксируется, но побочные эффекты callback не откатываются. Ошибки и пустые входы проверяются [общими тестами](contracts_test.go).

Генерация:

```sh
go generate ./lessons/journey
```

Сценарии, решения и проверки:

```sh
go test ./lessons/journey ./lessons/laws -v
go test ./lessons/journey -run '^$' -fuzz '^FuzzApproachAgreement$' -fuzztime=5s -parallel=2
go test ./lessons/laws -run '^$' -fuzz '^FuzzFractionLaws$' -fuzztime=5s -parallel=2
```

Измерения запускайте отдельно от fuzz-тестов:

```sh
go test ./lessons/journey -run '^$' -bench 'BenchmarkJourney|BenchmarkAnyPrepared' -benchmem -benchtime=100ms -count=3 -cpu=1
```

Сохранённый [сырой замер](results/bench-go1.25.4.txt) относится к Go 1.25.4 на Apple M1 Pro. [Таблицы](results/tables.md) вычисляются скриптом `node scripts/summarize-journey.mjs`. Это короткий учебный замер, не универсальный рейтинг подходов.
