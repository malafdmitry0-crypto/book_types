package journey

import (
	"fmt"
	"slices"

	"gotypes/algorithms"
	"gotypes/algorithms/closures"
	"gotypes/fraction"
)

// Общие только данные и предметные операции. Каждый Example вызывает своё ядро напрямую.
func exampleFractions() []fraction.Fraction {
	return []fraction.Fraction{
		{Numerator: 1, Denominator: 3},
		{Numerator: 2, Denominator: 4},
		{Numerator: -1, Denominator: 6},
	}
}

func Example_directConcrete() {
	values := exampleFractions()
	zero := fraction.Fraction{Denominator: 1}
	index, findErr := ConcreteFind(values, IsHalf)
	total, sumErr := ConcreteReduce(values, zero, Add)
	fmt.Println(index, findErr, total, sumErr)
	// Output: 1 <nil> 2/3 <nil>
}

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

func Example_directInterface() {
	values := exampleFractions()
	query := Search{Values: values, Test: IsHalf}
	index, findErr := algorithms.Find(query)
	sum := Fold{Values: values, Total: fraction.Fraction{Denominator: 1}, Step: Add}
	sumErr := algorithms.Sum(&sum)
	fmt.Println(index, findErr, sum.Total, sumErr)
	// Output: 1 <nil> 2/3 <nil>
}

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

func Example_directGeneration() {
	values := exampleFractions()
	zero := fraction.Fraction{Denominator: 1}
	index, findErr := GeneratedFind(values, IsHalf)
	total, sumErr := GeneratedReduce(values, zero, Add)
	fmt.Println(index, findErr, total, sumErr)
	// Output: 1 <nil> 2/3 <nil>
}

func Example_directGeneric() {
	values := exampleFractions()
	zero := fraction.Fraction{Denominator: 1}
	index, findErr := GenericFind(values, IsHalf)
	total, sumErr := GenericReduce(values, zero, Add)
	fmt.Println(index, findErr, total, sumErr)
	// Output: 1 <nil> 2/3 <nil>
}

func Example_directIterator() {
	values := exampleFractions()
	zero := fraction.Fraction{Denominator: 1}
	index, findErr := IteratorFind(slices.Values(values), IsHalf)
	total, sumErr := IteratorReduce(slices.Values(values), zero, Add)
	fmt.Println(index, findErr, total, sumErr)
	// Output: 1 <nil> 2/3 <nil>
}
