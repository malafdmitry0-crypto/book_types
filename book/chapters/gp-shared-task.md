# 19. Одна задача на дробях: восемь реализаций

> Сравним подходы на одинаковых данных, с одинаковыми результатами и одинаковым договором об ошибках.

Код главы: [полный сценарий](../../lessons/journey/example_test.go) · [все варианты вызова](../../lessons/journey/approaches.go) · [общие тесты](../../lessons/journey/contracts_test.go).

## Зачем нужен ещё один сквозной пример

В основном курсе мы использовали дроби и деньги для предметных ошибок, а User — для простого сравнения сигнатур. Здесь убираем это переключение: каждый подход решает одну задачу над Fraction. Теперь видно, какой код относится к алгоритму, а какой нужен для передачи данных через выбранный API.

Данные — три дроби: 1/3, 2/4 и −1/6. Нужно найти первую половину, сложить все значения и выбрать положительные. Исходный срез не меняется.

| Операция | Ожидаемый результат |
|---|---|
| Find с условием Equal(1/2) | Индекс 1: дробь 2/4 математически равна 1/2 |
| Sum с начальным 0/1 | 2/3 |
| Filter с условием Compare(0/1) > 0 | Две исходные дроби: 1/3 и 2/4 |

String печатает вторую выбранную дробь как 1/2. Filter не нормализует её поля: он сохраняет исходное значение 2/4, а сокращение относится к строковому представлению.

## Договор, который не меняется между реализациями

Find обходит слева направо, возвращает первый индекс или −1 и останавливается при первой ошибке. Sum выполняет левую свёртку с явно переданным начальным значением. При ошибке возвращается последняя успешная сумма. Ошибка имеет приоритет над одновременно возвращённым true или новым значением.

Входные значения и длина не изменяются во время обхода. Callback может вести внешний счётчик вызовов, но не должен менять исходный срез. Это существенное условие: interface{}-вариант предварительно копирует значения при упаковке, а другие варианты могут читать исходное хранилище непосредственно.

Общий договор рассчитан на подходящие ненулевые callback. Невалидная дробь обнаруживается только при её посещении. Проверка совместимости динамических аргументов reflection выполняется до обхода, даже для пустого среза; это дополнительная проверка его собственного API.

## Восемь способов выразить один алгоритм

| Подход | Что получает ядро | Где остаются дроби и сумма |
|---|---|---|
| Конкретный | []Fraction и типизированный callback | В параметрах и локальных переменных функции |
| Замыкание | Длина и функция от индекса | В окружении замыкания |
| Интерфейс | Len/Match либо Len/Add | В Search и Fold |
| Пустой интерфейс | []interface{} и интерфейсные аргументы callback | В упакованных значениях; нужны проверки и извлечение |
| Reflection | Значения и callback как interface{} | В reflect.Value; типы связывает проверка сигнатуры |
| Генерация | Обычная функция для Fraction | В сгенерированной специализации |
| Дженерики | []E и func(E), либо func(A,E) | В параметрах типов E и A |
| Итератор | iter.Seq[E] | Источник доставляет значения, свёртка хранит A |

Это не восемь вызовов одной generic-функции. Ядра находятся в отдельных файлах: concrete.go, closures.go, interfaces.go, any.go, reflection.go, generated.go, generic.go и iterator.go. Ранний интерфейсный Find и обход Sum переиспользуют уже написанный пакет algorithms.

## Сначала прочитайте вызов

В [Approaches](../../lessons/journey/approaches.go) есть единый стенд с полями Find, Filter и Reduce. Он нужен тестам и примерам, чтобы запускать варианты одинаковым способом. Этот удобный интерфейс стенда не подменяет сигнатуры сравниваемых ядер.

Например, вариант с замыканием передаёт в FindIndex длину и функцию, обращающуюся к s[i]. Вариант с интерфейсом создаёт Search{Values: s, Test: p}. Вариант с пустым интерфейсом сначала вызывает Box(s), затем AnyFind, причём каждая проверка извлекает Fraction. Генерация и generics обходятся без этого преобразования контейнера.

Sum в стенде определён через Reduce с операцией Add. Поэтому порядок, обработка ошибки и частичная сумма наследуются от соответствующей реализации Reduce, а не от отдельного общего арифметического цикла.

## Восемь реализаций целиком: как читать дальше

Ниже у каждого подхода есть собственный код обхода и отдельный прямой вызов. Общий цикл по Approaches остаётся в конце как проверка результатов. Чтобы понять алгоритмы, он больше не нужен.

В этой части подробно показаны Find и сумма через Reduce. Для суммы шаг Reduce — обычный Fraction.Add. Мы показываем тело Reduce в каждом варианте, поэтому за названием Sum не скрывается общий цикл другого подхода. Filter остаётся в тех же исходниках и подробно рассматривается в отдельной главе о Filter/Reduce.

Фрагменты ниже взяты из компилируемых файлов. Объявления относятся к пакету journey, кроме явно отмеченных ядер algorithms и closures. Общие импорты сценариев — fmt, slices, gotypes/fraction, gotypes/algorithms и gotypes/algorithms/closures; reflection-ядро также использует reflect. Полный файл прямых вызовов: [direct_example_test.go](../../lessons/journey/direct_example_test.go).

## Общие данные и предметные операции

Общими являются только исходные дроби, проверка половины и само сложение. Они не знают способа обхода. Ниже также объявлены типы функций Predicate и Step, используемые адаптерами и обёртками.

```go
func exampleFractions() []fraction.Fraction {
	return []fraction.Fraction{
		{Numerator: 1, Denominator: 3},
		{Numerator: 2, Denominator: 4},
		{Numerator: -1, Denominator: 6},
	}
}
```

```go
type Predicate func(fraction.Fraction) (bool, error)
type Step func(fraction.Fraction, fraction.Fraction) (fraction.Fraction, error)
```

```go
func IsHalf(value fraction.Fraction) (bool, error) {
	if err := value.Validate(); err != nil {
		return false, err
	}
	return value.Equal(fraction.Fraction{Numerator: 1, Denominator: 2}), nil
}

func Add(a, b fraction.Fraction) (fraction.Fraction, error) { return a.Add(b) }
```

## Реализация 1. Конкретные функции для Fraction

Алгоритм получает сам срез. В цикле value уже имеет тип Fraction, а total — тип Fraction. Ни интерфейсного контейнера, ни адаптера, ни параметров типов здесь нет.

```go
func ConcreteFind(src []fraction.Fraction, test func(fraction.Fraction) (bool, error)) (int, error) {
	for i, value := range src {
		matched, err := test(value)
		if err != nil {
			return -1, err
		}
		if matched {
			return i, nil
		}
	}
	return -1, nil
}

func ConcreteReduce(src []fraction.Fraction, initial fraction.Fraction, step func(fraction.Fraction, fraction.Fraction) (fraction.Fraction, error)) (fraction.Fraction, error) {
	total := initial
	for _, value := range src {
		next, err := step(total, value)
		if err != nil {
			return total, err
		}
		total = next
	}
	return total, nil
}
```

Прямой вызов. Reduce с шагом Add выполняет сумму слева направо:

```go
func Example_directConcrete() {
	values := exampleFractions()
	zero := fraction.Fraction{Denominator: 1}
	index, findErr := ConcreteFind(values, IsHalf)
	total, sumErr := ConcreteReduce(values, zero, Add)
	fmt.Println(index, findErr, total, sumErr)
	// Output: 1 <nil> 2/3 <nil>
}
```

При переходе на Money пришлось бы написать другую конкретную сигнатуру. Зато весь договор виден в одном обычном цикле. [Исходник](../../lessons/journey/concrete.go).

## Реализация 2. Длина и замыкание

Здесь ядро не получает дробь. FindIndex из пакета closures передаёт индекс callback, а ClosureReduce вызывает шаг по индексу. Типизированная сумма хранится в окружении замыкания вызывающего кода.

```go
func FindIndex(n int, test func(int) (bool, error)) (int, error) {
	for i := 0; i < n; i++ {
		matched, err := test(i)
		if err != nil {
			return -1, err
		}
		if matched {
			return i, nil
		}
	}
	return -1, nil
}
```

```go
func ClosureReduce(n int, stepAt func(int) error) error {
	for i := 0; i < n; i++ {
		if err := stepAt(i); err != nil {
			return err
		}
	}
	return nil
}
```

В вызове видно всю связь с данными: values[i] и total находятся у вызывающего. Новое total присваивается только после успешного Add.

```go
func Example_directClosure() {
	values := exampleFractions()
	index, findErr := closures.FindIndex(len(values), func(i int) (bool, error) { return IsHalf(values[i]) })
	total := fraction.Fraction{Denominator: 1}
	sumErr := ClosureReduce(len(values), func(i int) error {
		next, err := total.Add(values[i])
		if err != nil {
			return err
		}
		total = next
		return nil
	})
	fmt.Println(index, findErr, total, sumErr)
	// Output: 1 <nil> 2/3 <nil>
}
```

Изменение типа среза меняет замыкание, но не эти индексные циклы. [Исходник обхода](../../lessons/journey/closures.go).

## Реализация 3. Интерфейсы и адаптеры

Ядра ниже находятся в пакете algorithms. Они запрашивают только длину и операции по индексам. Find не знает Fraction; Sum не знает тип аккумулятора.

```go
type Predicate interface {
	Len() int
	Match(i int) (bool, error)
}

func Find(data Predicate) (int, error) {
	n := data.Len()
	for i := 0; i < n; i++ {
		matched, err := data.Match(i)
		if err != nil {
			return -1, err
		}
		if matched {
			return i, nil
		}
	}
	return -1, nil
}

type Summable interface {
	Len() int
	Add(i int) error
}

func Sum(data Summable) error {
	n := data.Len()
	for i := 0; i < n; i++ {
		if err := data.Add(i); err != nil {
			return err
		}
	}
	return nil
}
```

А это конкретные адаптеры из journey. Search хранит срез и условие. Fold хранит срез, сумму и шаг сложения. Их методы переводят индекс в действие над Fraction.

```go
type Search struct {
	Values []fraction.Fraction
	Test   func(fraction.Fraction) (bool, error)
}

func (s Search) Len() int                  { return len(s.Values) }
func (s Search) Match(i int) (bool, error) { return s.Test(s.Values[i]) }

type Fold struct {
	Values []fraction.Fraction
	Total  fraction.Fraction
	Step   func(fraction.Fraction, fraction.Fraction) (fraction.Fraction, error)
}

func (f *Fold) Len() int { return len(f.Values) }
func (f *Fold) Add(i int) error {
	next, err := f.Step(f.Total, f.Values[i])
	if err != nil {
		return err
	}
	f.Total = next
	return nil
}
```

```go
func Example_directInterface() {
	values := exampleFractions()
	query := Search{Values: values, Test: IsHalf}
	index, findErr := algorithms.Find(query)
	sum := Fold{Values: values, Total: fraction.Fraction{Denominator: 1}, Step: Add}
	sumErr := algorithms.Sum(&sum)
	fmt.Println(index, findErr, sum.Total, sumErr)
	// Output: 1 <nil> 2/3 <nil>
}
```

Здесь сумма меняется через указатель &sum. Сам алгоритм только организует вызовы Add(index). [Исходник адаптеров](../../lessons/journey/interfaces.go).

## Реализация 4. Пустой интерфейс

В этом API элементы пересекают границу алгоритма. Ядро принимает []interface{}, а callback должен самостоятельно извлечь конкретное значение. Результат Reduce тоже имеет статический тип interface{}.

```go
func AnyFind(src []interface{}, test func(interface{}) (bool, error)) (int, error) {
	for i, value := range src {
		matched, err := test(value)
		if err != nil {
			return -1, err
		}
		if matched {
			return i, nil
		}
	}
	return -1, nil
}

func AnyReduce(src []interface{}, initial interface{}, step func(interface{}, interface{}) (interface{}, error)) (interface{}, error) {
	total := initial
	for _, value := range src {
		next, err := step(total, value)
		if err != nil {
			return total, err
		}
		total = next
	}
	return total, nil
}
```

Полная подготовка и извлечение данных показаны ниже. Box создаёт интерфейсный срез. unbox проверяет именно Fraction. AnyPredicate и anyStep связывают интерфейсные аргументы с нашими предметными функциями. Здесь нельзя просто вызвать Add у interface{}.

```go
func Box(src []fraction.Fraction) []interface{} {
	if src == nil {
		return nil
	}
	out := make([]interface{}, len(src))
	for i, v := range src {
		out[i] = v
	}
	return out
}

func unbox(value interface{}) (fraction.Fraction, error) {
	f, ok := value.(fraction.Fraction)
	if !ok {
		return fraction.Fraction{}, fmt.Errorf("expected Fraction, got %T", value)
	}
	return f, nil
}

func AnyPredicate(test Predicate) func(interface{}) (bool, error) {
	return func(value interface{}) (bool, error) {
		f, err := unbox(value)
		if err != nil {
			return false, err
		}
		return test(f)
	}
}

func anyStep(step Step) func(interface{}, interface{}) (interface{}, error) {
	return func(a, b interface{}) (interface{}, error) {
		left, err := unbox(a)
		if err != nil {
			return nil, err
		}
		right, err := unbox(b)
		if err != nil {
			return nil, err
		}
		return step(left, right)
	}
}
```

Теперь прямой вызов содержит упаковку, оба алгоритма и извлечение суммы:

```go
func Example_directAny() {
	values := exampleFractions()
	boxed := Box(values)
	index, findErr := AnyFind(boxed, AnyPredicate(IsHalf))
	zero := fraction.Fraction{Denominator: 1}
	raw, sumErr := AnyReduce(boxed, zero, anyStep(Add))
	total, unboxErr := unbox(raw)
	if unboxErr != nil {
		fmt.Println(unboxErr)
		return
	}
	fmt.Println(index, findErr, total, sumErr)
	// Output: 1 <nil> 2/3 <nil>
}
```

Ошибочный конкретный тип обнаруживается при извлечении во время исполнения. В этом сценарии Box сам упаковывает Fraction, поэтому вход согласован с callback. [Исходник ядра](../../lessons/journey/any.go).

## Реализация 5. Reflection

Вызов снова принимает произвольные значения, но теперь библиотека исследует срез и сигнатуры функций. Проверки выполняются до цикла, включая пустой вход. Полный код проверок приведён здесь, а не спрятан за необъяснённым helper.

```go
var errorType = reflect.TypeOf((*error)(nil)).Elem()

func checkPredicate(src, test interface{}) (reflect.Value, reflect.Value, error) {
	s, f := reflect.ValueOf(src), reflect.ValueOf(test)
	if !s.IsValid() || s.Kind() != reflect.Slice {
		return s, f, fmt.Errorf("source must be a slice")
	}
	if !f.IsValid() || f.Kind() != reflect.Func || f.IsNil() {
		return s, f, fmt.Errorf("predicate must be a non-nil function")
	}
	t := f.Type()
	if t.IsVariadic() || t.NumIn() != 1 || t.NumOut() != 2 || t.In(0) != s.Type().Elem() || t.Out(0) != reflect.TypeOf(false) || t.Out(1) != errorType {
		return s, f, fmt.Errorf("predicate must have signature func(E) (bool, error)")
	}
	return s, f, nil
}

func reflectedError(v reflect.Value) error {
	if v.IsNil() {
		return nil
	}
	return v.Interface().(error)
}
```

После проверки Find вызывает предикат через reflect.Value.Call и читает bool и error.

```go
func ReflectFind(src, test interface{}) (int, error) {
	s, f, err := checkPredicate(src, test)
	if err != nil {
		return -1, err
	}
	for i := 0; i < s.Len(); i++ {
		r := f.Call([]reflect.Value{s.Index(i)})
		if err := reflectedError(r[1]); err != nil {
			return -1, err
		}
		if r[0].Bool() {
			return i, nil
		}
	}
	return -1, nil
}
```

Reduce отдельно проверяет связь типов аккумулятора, элемента и результата шага. Затем каждый следующий аккумулятор становится reflect.Value результата успешного вызова.

```go
func ReflectReduce(src, initial, step interface{}) (interface{}, error) {
	s, a, f := reflect.ValueOf(src), reflect.ValueOf(initial), reflect.ValueOf(step)
	if !s.IsValid() || s.Kind() != reflect.Slice || !a.IsValid() {
		return initial, fmt.Errorf("source must be a slice and initial must have a concrete type")
	}
	if !f.IsValid() || f.Kind() != reflect.Func || f.IsNil() {
		return initial, fmt.Errorf("step must be a non-nil function")
	}
	t := f.Type()
	if t.IsVariadic() || t.NumIn() != 2 || t.NumOut() != 2 || t.In(0) != a.Type() || t.In(1) != s.Type().Elem() || t.Out(0) != a.Type() || t.Out(1) != errorType {
		return initial, fmt.Errorf("step must have signature func(A,E) (A,error)")
	}
	for i := 0; i < s.Len(); i++ {
		r := f.Call([]reflect.Value{a, s.Index(i)})
		if err := reflectedError(r[1]); err != nil {
			return a.Interface(), err
		}
		a = r[0]
	}
	return a.Interface(), nil
}
```

Прямой вызов не упаковывает элементы по одному, но получает сумму как interface{}. Для извлечения используется полностью показанная выше unbox.

```go
func Example_directReflection() {
	values := exampleFractions()
	index, findErr := ReflectFind(values, IsHalf)
	zero := fraction.Fraction{Denominator: 1}
	raw, sumErr := ReflectReduce(values, zero, Add)
	total, unboxErr := unbox(raw)
	if unboxErr != nil {
		fmt.Println(unboxErr)
		return
	}
	fmt.Println(index, findErr, total, sumErr)
	// Output: 1 <nil> 2/3 <nil>
}
```

Reflection проверяет форму функций; корректность дроби и переполнение по-прежнему проверяет Fraction.Add. [Исходник](../../lessons/journey/reflection.go).

## Реализация 6. Сгенерированные функции

Генератор создал следующие обычные функции для Fraction. Их сходство с конкретным вариантом ожидаемо: различается способ получения исходника. Во время выполнения шаблон не интерпретируется.

```go
func GeneratedFind(src []fraction.Fraction, test func(fraction.Fraction) (bool, error)) (int, error) {
	for i, value := range src {
		matched, err := test(value)
		if err != nil {
			return -1, err
		}
		if matched {
			return i, nil
		}
	}
	return -1, nil
}

func GeneratedReduce(src []fraction.Fraction, initial fraction.Fraction, step func(fraction.Fraction, fraction.Fraction) (fraction.Fraction, error)) (fraction.Fraction, error) {
	total := initial
	for _, value := range src {
		next, err := step(total, value)
		if err != nil {
			return total, err
		}
		total = next
	}
	return total, nil
}
```

Прямой вызов использует выпущенные имена функций:

```go
func Example_directGeneration() {
	values := exampleFractions()
	zero := fraction.Fraction{Denominator: 1}
	index, findErr := GeneratedFind(values, IsHalf)
	total, sumErr := GeneratedReduce(values, zero, Add)
	fmt.Println(index, findErr, total, sumErr)
	// Output: 1 <nil> 2/3 <nil>
}
```

[Шаблон](../../lessons/journey/internal/generate/algorithms.tmpl) и [генератор](../../lessons/journey/internal/generate/main.go) находятся в проекте; команда регенерации остаётся в конце главы. Это именно сгенерированный код, а не переименование generic-вызова.

## Реализация 7. Параметры типов

В Find параметр E связывает срез и предикат. В Reduce параметры E и A связывают элемент, аккумулятор и результат шага. В нашем вызове оба выбраны как Fraction, но для другого Reduce они могут различаться.

```go
func GenericFind[E any](src []E, test func(E) (bool, error)) (int, error) {
	for i, value := range src {
		matched, err := test(value)
		if err != nil {
			return -1, err
		}
		if matched {
			return i, nil
		}
	}
	return -1, nil
}

func GenericReduce[E, A any](src []E, initial A, step func(A, E) (A, error)) (A, error) {
	total := initial
	for _, value := range src {
		next, err := step(total, value)
		if err != nil {
			return total, err
		}
		total = next
	}
	return total, nil
}
```

Компилятор выводит типы из values, zero и Add:

```go
func Example_directGeneric() {
	values := exampleFractions()
	zero := fraction.Fraction{Denominator: 1}
	index, findErr := GenericFind(values, IsHalf)
	total, sumErr := GenericReduce(values, zero, Add)
	fmt.Println(index, findErr, total, sumErr)
	// Output: 1 <nil> 2/3 <nil>
}
```

Здесь не нужны ни коллекционный адаптер, ни упаковка, ни assertion результата. При этом error и порядок накопления остались частью явно написанного алгоритма. [Исходник](../../lessons/journey/generic.go).

## Реализация 8. Типизированный итератор

Ядро получает источник iter.Seq, а не срез. Find считает позиции в порядке получения значений. Reduce обрабатывает значения по мере поступления. Выход из range прекращает корректно реализованный источник.

```go
func IteratorFind[E any](src iter.Seq[E], test func(E) (bool, error)) (int, error) {
	index := 0
	for value := range src {
		matched, err := test(value)
		if err != nil {
			return -1, err
		}
		if matched {
			return index, nil
		}
		index++
	}
	return -1, nil
}

func IteratorReduce[E, A any](src iter.Seq[E], initial A, step func(A, E) (A, error)) (A, error) {
	total := initial
	for value := range src {
		next, err := step(total, value)
		if err != nil {
			return total, err
		}
		total = next
	}
	return total, nil
}
```

В этом сценарии источником служит slices.Values. Каждый алгоритм получает собственный обход тех же дробей:

```go
func Example_directIterator() {
	values := exampleFractions()
	zero := fraction.Fraction{Denominator: 1}
	index, findErr := IteratorFind(slices.Values(values), IsHalf)
	total, sumErr := IteratorReduce(slices.Values(values), zero, Add)
	fmt.Println(index, findErr, total, sumErr)
	// Output: 1 <nil> 2/3 <nil>
}
```

Для произвольного источника порядковая позиция не обязательно является индексом хранилища. Здесь она совпадает с индексом среза, потому что источник — slices.Values. [Исходник](../../lessons/journey/iterator.go).

## Проверить каждый вариант отдельно

Все восемь показанных вызовов находятся в одном файле и запускаются без стенда Approaches:

```sh
go test ./lessons/journey -run '^Example_direct' -v
```

У каждого собственный Output: индекс 1, сумма 2/3, обе ошибки nil. Общие тесты дополнительно проверяют пустой вход, остановку при ошибке и частичную сумму. Следующий общий запуск нужен только для сводного сравнения, в том числе Filter.

## Полный запуск

```go
func Example() {
	values := []fraction.Fraction{{Numerator: 1, Denominator: 3}, {Numerator: 2, Denominator: 4}, {Numerator: -1, Denominator: 6}}
	zero := fraction.Fraction{Denominator: 1}
	for _, approach := range journey.Approaches() {
		index, err := approach.Find(values, journey.IsHalf)
		if err != nil {
			fmt.Println(err)
			return
		}
		total, err := approach.Sum(values, zero)
		if err != nil {
			fmt.Println(err)
			return
		}
		selected, err := approach.Filter(values, journey.Positive)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s: Find=%d Sum=%s Filter=%v\n", approach.Name, index, total, selected)
	}
	// Output:
	// concrete: Find=1 Sum=2/3 Filter=[1/3 1/2]
	// closure: Find=1 Sum=2/3 Filter=[1/3 1/2]
	// interface: Find=1 Sum=2/3 Filter=[1/3 1/2]
	// empty-interface: Find=1 Sum=2/3 Filter=[1/3 1/2]
	// reflection: Find=1 Sum=2/3 Filter=[1/3 1/2]
	// generation: Find=1 Sum=2/3 Filter=[1/3 1/2]
	// generic: Find=1 Sum=2/3 Filter=[1/3 1/2]
	// iterator: Find=1 Sum=2/3 Filter=[1/3 1/2]
}
```

Сначала проследите конкретный вариант, затем откройте реализацию следующего. Для каждого задайте три вопроса: где берётся элемент, кто вызывает Add и кто хранит результат после ошибки?

## Почему итератор тоже возвращает индекс

В основном курсе FindSeq возвращал значение с bool. Здесь IteratorFind возвращает порядковую позицию в обходе, чтобы сравнить результат со срезовыми вариантами. Для slices.Values эта позиция совпадает с индексом исходного среза. Для произвольного источника она не обещает возможность повторного доступа по индексу.

Filter-итератор ленивый, но общий стенд полностью материализует его результат. Поэтому сравнение Filter не выигрывает за счёт того, что вычисление ещё не началось. Отдельный тест показывает раннюю остановку источника при break.

## Генерация должна быть настоящей

[Шаблон](../../lessons/journey/internal/generate/algorithms.tmpl) описывает Find, Filter и Reduce. Генератор подставляет Fraction, форматирует исходник и создаёт [generated.go](../../lessons/journey/generated.go). Получившиеся функции компилируются и проходят те же тесты, что остальные.

```sh
go generate ./lessons/journey
go test ./lessons/journey -run Example -v
go test ./lessons/journey -run Test -v
```

Генератор запускается до обычной компиляции. Файл результата включён в проект, поэтому для чтения и первого запуска регенерация не нужна.

## Проверка понимания

Если заменить задачу на сумму Money, какие варианты потребуют новой специализации или адаптера? Какие смогут принять новый тип, но потребуют динамической проверки? Где тип результата будет проверен компилятором?

---

[← Сравнение подходов и проектирование собственной библиотеки](gp-12-synthesis.md) · [Оглавление](../README.md) · [Законы операций: что не выражает сигнатура →](gp-operation-laws.md)
