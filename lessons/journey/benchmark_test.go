package journey_test

import (
	"gotypes/fraction"
	"gotypes/lessons/journey"
	"testing"
)

var sinkIndex int
var sinkFraction fraction.Fraction
var sinkSlice []fraction.Fraction

func benchmarkValues() []fraction.Fraction {
	values := make([]fraction.Fraction, 64)
	for i := range values {
		values[i] = fraction.Fraction{Numerator: 1, Denominator: 3}
	}
	return values
}
func cheapMissing(value fraction.Fraction) (bool, error) {
	if err := value.Validate(); err != nil {
		return false, err
	}
	return value.Numerator == 7, nil
}

// Подготовка исходного []Fraction не входит в измерение. Упаковка, проверка
// reflection-сигнатур, распаковка и материализация результата входят.
func BenchmarkJourney(b *testing.B) {
	values := benchmarkValues()
	zero := fraction.Fraction{Denominator: 1}
	for _, a := range journey.Approaches() {
		b.Run("FindCheap/"+a.Name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				index, err := a.Find(values, cheapMissing)
				if err != nil {
					b.Fatal(err)
				}
				sinkIndex = index
			}
		})
		b.Run("FindMath/"+a.Name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				index, err := a.Find(values, journey.IsHalf)
				if err != nil {
					b.Fatal(err)
				}
				sinkIndex = index
			}
		})
		b.Run("Sum/"+a.Name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				result, err := a.Sum(values, zero)
				if err != nil {
					b.Fatal(err)
				}
				sinkFraction = result
			}
		})
		b.Run("Filter/"+a.Name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				result, err := a.Filter(values, journey.Positive)
				if err != nil {
					b.Fatal(err)
				}
				sinkSlice = result
			}
		})
	}
}

// Отдельный вопрос: стоимость AnyFind с уже подготовленным boxed-срезом.
// Его нельзя называть полной стоимостью вызова для исходного []Fraction.
func BenchmarkAnyPrepared(b *testing.B) {
	boxed := journey.Box(benchmarkValues())
	test := journey.AnyPredicate(cheapMissing)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index, err := journey.AnyFind(boxed, test)
		if err != nil {
			b.Fatal(err)
		}
		sinkIndex = index
	}
}
