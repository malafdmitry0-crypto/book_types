package containers_test

import (
	"fmt"
	"gotypes/fraction"
	"gotypes/lessons/containers"
)

func ExampleStack() {
	var s containers.Stack[fraction.Fraction] // нулевое значение пригодно
	s.Push(fraction.Fraction{Numerator: 1, Denominator: 2})
	s.Push(fraction.Fraction{Numerator: 1, Denominator: 3})
	fmt.Println("Len:", s.Len(), "Stack:", s)
	top, ok := s.Pop()
	fmt.Println("Pop:", top, ok)
	s.Pop()
	empty, ok := s.Pop()
	fmt.Println("Пустой стек:", empty, ok)
	// Output:
	// Len: 2 Stack: [1/2 1/3]
	// Pop: 1/3 true
	// Пустой стек: invalid fraction(0/0) false
}

func ExampleStack_pointers() {
	var values containers.Stack[fraction.Fraction]
	var pointers containers.Stack[*fraction.Fraction]
	half := fraction.Fraction{Numerator: 1, Denominator: 2}
	values.Push(half)
	pointers.Push(&half)
	// values.Push(&half) // ошибка компиляции: *Fraction не является Fraction
	// values = pointers  // ошибка компиляции: Stack[*Fraction] — другой тип
	fmt.Printf("%T %T\n", values, pointers)
	// Output:
	// containers.Stack[gotypes/fraction.Fraction] containers.Stack[*gotypes/fraction.Fraction]
}

func ExampleMapStack() {
	var s containers.Stack[fraction.Fraction]
	s.Push(fraction.Fraction{Numerator: 1, Denominator: 2})
	s.Push(fraction.Fraction{Numerator: 2, Denominator: 4})
	texts := containers.MapStack(s, fraction.Fraction.String)
	fmt.Printf("%T %v\n", texts, texts)
	// Output:
	// containers.Stack[string] [1/2 1/2]
}

func ExampleFill() {
	var s containers.Stack[int]
	containers.Fill(&s, 1, 2, 3)
	// containers.Fill(s, 4) // ошибка компиляции: Stack[int] не реализует Container[int]
	var _ fmt.Stringer = s // String объявлен на значении, подходят и Stack[int], и *Stack[int]
	fmt.Println(s.Len(), s)
	// Output:
	// 3 [1 2 3]
}

func ExampleSet() {
	s := containers.NewSet[string]()
	s.Add("go")
	s.Add("generic")
	s.Add("go")
	fmt.Println(s.Len(), s.Has("go"), s.Has("rust"))
	fmt.Println(containers.SortedKeys(s))
	// Output:
	// 2 true false
	// [generic go]
}

func ExampleSet_zeroValue() {
	var s containers.Set[string] // карта внутри равна nil
	fmt.Println("Чтение:", s.Len(), s.Has("go"))
	defer func() { fmt.Println("Запись:", recover()) }()
	s.Add("go")
	// Output:
	// Чтение: 0 false
	// Запись: assignment to entry in nil map
}
