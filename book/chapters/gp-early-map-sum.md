# 9. Map и Sum: где в раннем Go хранился типизированный результат

Код главы: [example_test.go](../../algorithms/example_test.go) · [algorithms_test.go](../../algorithms/algorithms_test.go).


> Поиск смог вернуть int. Преобразование и сумма должны сохранить конкретное значение — поэтому адаптер становится важнее.

Поиск возвращал индекс, поэтому тип результата не был проблемой. Map должен получить строки, а Sum — новую денежную сумму. Разберём, кто создаёт и хранит эти результаты, если общий алгоритм не знает их типов.

## Map не может угадать []R

Наша публичная сигнатура `Map(data Mapping) error` не обещает вернуть произвольный срез. Она просит объект, который умеет преобразовать элемент и записать результат:

```go
type Mapping interface {
    Len() int
    Apply(i int) error
}
```

```go
func Map(data Mapping) error {
	n := data.Len()
	for i := 0; i < n; i++ {
		if err := data.Apply(i); err != nil {
			return err
		}
	}
	return nil
}
```

Тип результата не исчез из программы. Он просто отсутствует в сигнатуре общего алгоритма и хранится у вызывающего. Вот конкретный адаптер для `Fraction → string`:

```go
type fractionText struct {
    Source    []fraction.Fraction
    Target    []string
    Transform func(fraction.Fraction) (string, error)
}

func (m fractionText) Len() int { return len(m.Source) }
func (m fractionText) Apply(i int) error {
    text, err := m.Transform(m.Source[i])
    if err != nil {
        return err
    }
    m.Target[i] = text
    return nil
}
```

Связь Fraction → string явно типизирована внутри адаптера. Неправильный callback `func(Money) (string, error)` не подходит полю Transform. Но для связи Money → string потребуется новый адаптер.

## Полный вызов и расположение результата

```go
func ExampleMap() {
	values := []fraction.Fraction{
		{Numerator: 2, Denominator: 4},
		{Numerator: 6, Denominator: 3},
	}
	texts := make([]string, len(values))
	toText := func(value fraction.Fraction) (string, error) {
		if err := value.Validate(); err != nil {
			return "", err
		}
		return value.String(), nil
	}
	job := fractionText{Source: values, Target: texts, Transform: toText}

	err := algorithms.Map(job)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	// Результат записан в texts через адаптер.
	fmt.Println("Строки:", texts)
	// Output: Строки: [1/2 2]
}
```

Этот код берётся из [исполняемого примера](../../algorithms/example_test.go). `texts` создан до Map, остаётся `[]string` после Map и не требует assertion. При nil-источнике сохраняем nil-результат; этот выбор сделал вызывающий, а не общий цикл.

Почему Apply имеет получателя по значению и всё же меняет texts? Адаптер содержит заголовок среза. Копия заголовка продолжает ссылаться на тот же backing array; присваивание Target[i] меняет общий элемент. Но переназначение самого поля Target внутри такой копии не изменило бы заголовок texts у вызывающего. Это принципиальное различие при проектировании адаптеров.

Наша реализация сначала вычисляет строку, проверяет ошибку и лишь потом записывает её. При ошибке на втором элементе первая строка остаётся записанной, второй и последующие слоты не меняются. Алгоритм не умеет откатить произвольные действия Apply.

## Sum: неизвестен не только элемент, но и ноль

У BoxInt корректный ноль — нулевая структура. У Fraction нулевая структура невалидна: нужен знаменатель 1. У Money нулю нужна валюта. Поэтому общий Sum не может просто создать «нулевое значение какого-то типа» и считать вопрос решённым.

```go
type Summable interface {
    Len() int
    Add(i int) error
}
```

```go
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

Начальное значение и текущее состояние суммы находятся в адаптере:

```go
type moneySum struct {
    Values []money.Money
    Total  money.Money
}

func (s *moneySum) Len() int { return len(s.Values) }
func (s *moneySum) Add(i int) error {
    next, err := s.Total.Add(s.Values[i])
    if err != nil {
        return err
    }
    s.Total = next
    return nil
}
```

Здесь два разных Add. Метод адаптера принимает индекс; метод Money принимает Money. Адаптер переводит один протокол в другой и обновляет аккумулятор только после успешной операции.

## Почему получатель — указатель

Total является значением структуры. Если метод адаптера получал бы копию moneySum и присваивал её полю Total, вызывающий не увидел бы нового аккумулятора. Поэтому изменяющий состояние метод определён у `*moneySum`, и в Sum передают `&sum`.

```go
var _ algorithms.Summable = (*moneySum)(nil)
// moneySum{} не реализует Summable: нужные методы имеют pointer receiver.
```

Это не поздняя сложность дженериков: method sets и видимость изменения состояния нужно понимать уже для раннего интерфейсного решения.

## Полный денежный пример

```go
func ExampleSum() {
	values := []money.Money{
		{Amount: 100, Currency: "USD"},
		{Amount: 250, Currency: "USD"},
	}
	initial := money.Money{Currency: "USD"} // Нулевая сумма с заданной валютой.
	sum := moneySum{Values: values, Total: initial}

	// Указатель позволяет адаптеру изменять своё поле Total.
	err := algorithms.Sum(&sum)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Сумма:", sum.Total)
	// Output: Сумма: 350 USD (minor units)
}
```

Если заменить вторую сумму на EUR, вызов вернёт ошибку несовместимой валюты. Total уже содержит первый успешно обработанный элемент; следующие элементы не посещаются. Если вход пуст, Total остаётся исходным. При повторном независимом вызове Sum адаптер нужно инициализировать заново — иначе он продолжит складывать с текущим Total.

Для дробей реализация fractionSum аналогична, но аккумулятор Fraction начинается с 0/1. Один и тот же Sum проверен с обоими типами в [algorithms_test.go](../../algorithms/algorithms_test.go).

## Что выиграли и чего не смогли выразить

Общие Map и Sum не знают конкретных типов и не используют reflection. Результаты сохраняют конкретный тип. Но связь входа и результата, их создание и часть служебного кода приходится повторять в адаптерах.

Это и есть содержательная граница ранней модели. Она позволяла обобщать алгоритм через операции, но не позволяла объявить один Map с произвольными связанными типами входа и выхода. Следующие механизмы будем оценивать по тому, как они перемещают именно эту границу.

## Проверка понимания

Почему fractionText может записывать элементы с value receiver, а moneySum для присваивания Total нужен pointer receiver? Что изменится, если вместо Target[i] адаптер Map станет делать append к своему полю Target?

---

[← Find, All и Any: один цикл для разных предметных условий](gp-early-search.md) · [Оглавление](../README.md) · [Min, Max, Sort и границы обещаний ранней библиотеки →](gp-early-contracts.md)
