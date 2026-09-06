package journey_test

import (
	"gotypes/algorithms"
	"gotypes/boxint"
	"gotypes/lessons/generic"
	"testing"
)

// Три способа сложить 64 значения BoxInt. Операция Add дешёвая, поэтому
// стоимость самого способа вызова не тонет в предметной работе, как в Sum
// над Fraction с math/big.

func shapeValues() []boxint.BoxInt {
	values := make([]boxint.BoxInt, 64)
	for i := range values {
		values[i] = boxint.BoxInt{Value: 3}
	}
	return values
}

// concreteSum вызывает конкретный метод boxint.BoxInt.Add напрямую.
func concreteSum(values []boxint.BoxInt, initial boxint.BoxInt) (boxint.BoxInt, error) {
	total := initial
	for _, value := range values {
		next, err := total.Add(value)
		if err != nil {
			return total, err
		}
		total = next
	}
	return total, nil
}

// boxFold — адаптер algorithms.Summable для BoxInt по образцу Fold из interfaces.go;
// внутри Add(i) вызывается конкретный метод, снаружи — метод интерфейса.
type boxFold struct {
	Values []boxint.BoxInt
	Total  boxint.BoxInt
}

func (f *boxFold) Len() int { return len(f.Values) }
func (f *boxFold) Add(i int) error {
	next, err := f.Total.Add(f.Values[i])
	if err != nil {
		return err
	}
	f.Total = next
	return nil
}

var _ algorithms.Summable = (*boxFold)(nil)

func interfaceSum(values []boxint.BoxInt, initial boxint.BoxInt) (boxint.BoxInt, error) {
	job := boxFold{Values: values, Total: initial}
	err := algorithms.Sum(&job)
	return job.Total, err
}

var sinkBox boxint.BoxInt

func TestShapeAgreement(t *testing.T) {
	values := shapeValues()
	want := boxint.BoxInt{Value: 64 * 3}
	got, err := concreteSum(values, boxint.BoxInt{})
	if err != nil || got != want {
		t.Fatalf("concrete: %v, %v", got, err)
	}
	got, err = generic.SumValues(values, boxint.BoxInt{})
	if err != nil || got != want {
		t.Fatalf("generic: %v, %v", got, err)
	}
	got, err = interfaceSum(values, boxint.BoxInt{})
	if err != nil || got != want {
		t.Fatalf("interface: %v, %v", got, err)
	}
}

func BenchmarkShape(b *testing.B) {
	values := shapeValues()
	zero := boxint.BoxInt{}
	b.Run("concrete", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			result, err := concreteSum(values, zero)
			if err != nil {
				b.Fatal(err)
			}
			sinkBox = result
		}
	})
	b.Run("generic", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			result, err := generic.SumValues(values, zero)
			if err != nil {
				b.Fatal(err)
			}
			sinkBox = result
		}
	})
	b.Run("interface", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			result, err := interfaceSum(values, zero)
			if err != nil {
				b.Fatal(err)
			}
			sinkBox = result
		}
	})
}
