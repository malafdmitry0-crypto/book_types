package journey_test

import (
	"gotypes/fraction"
	"gotypes/lessons/journey"
	"reflect"
	"testing"
)

// Решение практикума: общий контракт всех подходов на автоматически меняющемся входе.
func FuzzApproachAgreement(f *testing.F) {
	f.Add([]byte{1, 3, 5, 0})
	f.Add([]byte{})
	f.Add([]byte{255, 0, 128})
	f.Fuzz(func(t *testing.T, data []byte) {
		if len(data) > 32 {
			data = data[:32]
		} // Ограничиваем стоимость reflection и арифметики.
		var values []fraction.Fraction
		if data != nil {
			values = make([]fraction.Fraction, len(data))
		}
		for i, n := range data {
			values[i] = fraction.Fraction{Numerator: int64(int(n) - 128), Denominator: 6}
		}
		zero := fraction.Fraction{Denominator: 1}
		approaches := journey.Approaches()
		wantIndex, err := approaches[0].Find(values, journey.IsHalf)
		if err != nil {
			t.Fatal(err)
		}
		wantSum, err := approaches[0].Sum(values, zero)
		if err != nil {
			t.Fatal(err)
		}
		wantFilter, err := approaches[0].Filter(values, journey.Positive)
		if err != nil {
			t.Fatal(err)
		}
		for _, a := range approaches[1:] {
			index, err := a.Find(values, journey.IsHalf)
			if err != nil || index != wantIndex {
				t.Fatal(a.Name, index, err)
			}
			total, err := a.Sum(values, zero)
			if err != nil || !total.Equal(wantSum) {
				t.Fatal(a.Name, total, err)
			}
			selected, err := a.Filter(values, journey.Positive)
			if err != nil || !reflect.DeepEqual(selected, wantFilter) {
				t.Fatal(a.Name, selected, err)
			}
		}
	})
}
