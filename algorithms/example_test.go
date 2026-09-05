package algorithms_test

import (
	"fmt"

	"gotypes/algorithms"
	"gotypes/boxint"
	"gotypes/fraction"
	"gotypes/money"
)

// Каждый Example читается как маленький main: подготовка, вызов, результат.
// Запуск всех сценариев: go test ./algorithms -run Example -v

func ExampleFind() {
	// 1. Данные и условие поиска.
	values := []fraction.Fraction{
		{Numerator: 1, Denominator: 3},
		{Numerator: 2, Denominator: 4},
		{Numerator: 3, Denominator: 4},
	}
	half := fraction.Fraction{Numerator: 1, Denominator: 2}
	isHalf := func(value fraction.Fraction) (bool, error) {
		if err := value.Validate(); err != nil {
			return false, err
		}
		return value.Equal(half), nil
	}

	// 2. Адаптер связывает срез и условие с методами Len и Match.
	query := fractionQuery{Values: values, Test: isHalf}

	// 3. Алгоритм видит только интерфейс адаптера.
	index, err := algorithms.Find(query)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	if index == -1 {
		fmt.Println("Совпадений нет")
		return
	}

	// 4. Найденное значение берём из исходного типизированного среза.
	fmt.Println("Индекс:", index)
	fmt.Println("Дробь:", values[index])
	// Output:
	// Индекс: 1
	// Дробь: 1/2
}

func ExampleAll() {
	values := []fraction.Fraction{
		{Numerator: 1, Denominator: 3},
		{Numerator: 2, Denominator: 4},
	}
	zero := fraction.Fraction{Denominator: 1}
	isPositive := func(value fraction.Fraction) (bool, error) {
		comparison, err := value.Compare(zero)
		return comparison > 0, err
	}
	query := fractionQuery{Values: values, Test: isPositive}

	positive, err := algorithms.All(query)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Все положительные:", positive)
	// Output: Все положительные: true
}

func ExampleAny() {
	values := []fraction.Fraction{
		{Numerator: 1, Denominator: 3},
		{Denominator: 1},
	}
	isZero := func(value fraction.Fraction) (bool, error) {
		if err := value.Validate(); err != nil {
			return false, err
		}
		return value.IsZero(), nil
	}
	query := fractionQuery{Values: values, Test: isZero}

	found, err := algorithms.Any(query)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Есть ноль:", found)
	// Output: Есть ноль: true
}

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

func ExampleMin() {
	values := []boxint.BoxInt{{Value: 3}, {Value: 1}, {Value: 2}, {Value: 1}}
	collection := boxes(values)

	index, err := algorithms.Min(collection)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	if index == -1 {
		fmt.Println("Коллекция пуста")
		return
	}
	fmt.Println("Индекс:", index)
	fmt.Println("Минимум", values[index])
	// Output:
	// Индекс: 1
	// Минимум 1
}

func ExampleMax() {
	values := []boxint.BoxInt{{Value: 3}, {Value: 1}, {Value: 3}}
	collection := boxes(values)

	index, err := algorithms.Max(collection)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	if index == -1 {
		fmt.Println("Коллекция пуста")
		return
	}
	fmt.Println("Индекс:", index)
	fmt.Println("Максимум", values[index])
	// Output:
	// Индекс: 0
	// Максимум 3
}

func ExampleSort() {
	values := []boxint.BoxInt{{Value: 3}, {Value: 1}, {Value: 2}}
	collection := boxes(values) // Добавляем Len, Compare и Swap без копии элементов.

	err := algorithms.Sort(collection)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("После сортировки:", values)
	// Output: После сортировки: [1 2 3]
}

// Адаптер поиска: хранит дроби и условие, предоставляет алгоритму Len и Match.
type fractionQuery struct {
	Values []fraction.Fraction
	Test   func(fraction.Fraction) (bool, error)
}

func (q fractionQuery) Len() int                  { return len(q.Values) }
func (q fractionQuery) Match(i int) (bool, error) { return q.Test(q.Values[i]) }

var _ algorithms.Predicate = fractionQuery{}

// Адаптер преобразования: читает дроби и записывает строки.
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

var _ algorithms.Mapping = fractionText{}

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

var _ algorithms.Summable = (*moneySum)(nil)

// Срезовый адаптер предоставляет операции сравнения и перестановки.
type boxes []boxint.BoxInt

func (s boxes) Len() int                      { return len(s) }
func (s boxes) Compare(i, j int) (int, error) { return s[i].Compare(s[j]) }
func (s boxes) Swap(i, j int)                 { s[i], s[j] = s[j], s[i] }

var _ algorithms.Sortable = boxes(nil)
