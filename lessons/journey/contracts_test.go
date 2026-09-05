package journey_test

import (
	"errors"
	"gotypes/fraction"
	"gotypes/lessons/journey"
	"reflect"
	"slices"
	"testing"
)

func TestAllApproaches(t *testing.T) {
	zero := fraction.Fraction{Denominator: 1}
	for _, a := range journey.Approaches() {
		t.Run(a.Name, func(t *testing.T) {
			for _, src := range [][]fraction.Fraction{nil, {}, {{Numerator: 1, Denominator: 3}, {Numerator: 2, Denominator: 4}, {Numerator: -1, Denominator: 6}}} {
				index, err := a.Find(src, journey.IsHalf)
				want := -1
				if len(src) > 0 {
					want = 1
				}
				if index != want || err != nil {
					t.Fatal(index, err)
				}
				out, err := a.Filter(src, journey.Positive)
				if err != nil || (out == nil) != (src == nil) {
					t.Fatal(out, err)
				}
				if len(src) > 0 {
					if len(out) != 2 || out[0] != src[0] || out[1] != src[1] {
						t.Fatal(out)
					}
					out[0] = zero
					if src[0] == zero {
						t.Fatal("aliased result")
					}
				}
				total, err := a.Sum(src, zero)
				if err != nil {
					t.Fatal(err)
				}
				expected := zero
				if len(src) > 0 {
					expected = fraction.Fraction{Numerator: 2, Denominator: 3}
				}
				if !total.Equal(expected) {
					t.Fatal(total)
				}
			}
			// Пустой результат для непустого входа остаётся ненулевым срезом.
			out, err := a.Filter([]fraction.Fraction{zero}, func(fraction.Fraction) (bool, error) { return false, nil })
			if err != nil || out == nil || len(out) != 0 {
				t.Fatal(out, err)
			}
			stop := errors.New("stop")
			src := []fraction.Fraction{{Numerator: 1, Denominator: 3}, {Numerator: 2, Denominator: 4}, {Numerator: 3, Denominator: 4}}
			var visited []int64
			test := func(v fraction.Fraction) (bool, error) {
				visited = append(visited, v.Numerator)
				if v.Numerator == 2 {
					return true, stop
				}
				return true, nil
			}
			index, err := a.Find(src, test)
			if index != 0 || err != nil || !reflect.DeepEqual(visited, []int64{1}) {
				t.Fatal(index, err, visited)
			}
			visited = nil
			out, err = a.Filter(src, test)
			if !errors.Is(err, stop) || len(out) != 1 || !reflect.DeepEqual(visited, []int64{1, 2}) {
				t.Fatal(out, err, visited)
			}
			visited = nil
			total, err := a.Reduce(src, zero, func(acc, v fraction.Fraction) (fraction.Fraction, error) {
				visited = append(visited, v.Numerator)
				if v.Numerator == 2 {
					return zero, stop
				}
				return acc.Add(v)
			})
			if !errors.Is(err, stop) || !total.Equal(src[0]) || !reflect.DeepEqual(visited, []int64{1, 2}) {
				t.Fatal(total, err, visited)
			}
			index, err = a.Find(src, func(v fraction.Fraction) (bool, error) {
				if v.Numerator == 2 {
					return true, stop
				}
				return false, nil
			})
			if index != -1 || !errors.Is(err, stop) {
				t.Fatal(index, err)
			}
		})
	}
}

func TestDynamicValidation(t *testing.T) {
	for _, source := range []interface{}{nil, 42, []fraction.Fraction{}} {
		if _, err := journey.ReflectFind(source, func(int) (bool, error) { return true, nil }); err == nil {
			t.Fatal("expected invalid signature")
		}
	}
	var predicate func(fraction.Fraction) (bool, error)
	if _, err := journey.ReflectFilter([]fraction.Fraction{}, predicate); err == nil {
		t.Fatal("nil predicate")
	}
	if _, err := journey.ReflectReduce([]fraction.Fraction{}, 0, func(int, int) (int, error) { return 0, nil }); err == nil {
		t.Fatal("step mismatch")
	}
	if _, err := journey.AnyFind([]interface{}{42}, journey.AnyPredicate(journey.IsHalf)); err == nil {
		t.Fatal("wrong boxed type")
	}
}

func TestIteratorStopsAtConsumer(t *testing.T) {
	visited := 0
	source := func(yield func(int) bool) {
		for i := 0; i < 4; i++ {
			visited++
			if !yield(i) {
				return
			}
		}
	}
	for range journey.IteratorFilter(source, func(int) (bool, error) { return true, nil }) {
		break
	}
	if visited != 1 {
		t.Fatal(visited)
	}
	// Источник прекращается и при ошибке предиката.
	calls := 0
	stop := errors.New("stop")
	for _, err := range journey.IteratorFilter(slices.Values([]int{1, 2, 3}), func(int) (bool, error) { calls++; return true, stop }) {
		if !errors.Is(err, stop) {
			t.Fatal(err)
		}
	}
	if calls != 1 {
		t.Fatal(calls)
	}
}

func TestArithmeticErrorAgreement(t *testing.T) {
	zero := fraction.Fraction{Denominator: 1}
	half := fraction.Fraction{Numerator: 1, Denominator: 2}
	max := fraction.Fraction{Numerator: 1<<63 - 1, Denominator: 1}
	one := fraction.Fraction{Numerator: 1, Denominator: 1}
	for _, a := range journey.Approaches() {
		t.Run(a.Name, func(t *testing.T) {
			total, err := a.Sum([]fraction.Fraction{max, one}, zero)
			if !errors.Is(err, fraction.ErrOverflow) || !total.Equal(max) {
				t.Fatal(total, err)
			}
			total, err = a.Sum([]fraction.Fraction{half, {}}, zero)
			if !errors.Is(err, fraction.ErrZeroDenominator) || !total.Equal(half) {
				t.Fatal(total, err)
			}
			index, err := a.Find([]fraction.Fraction{half, {}}, journey.IsHalf)
			if err != nil || index != 0 {
				t.Fatal(index, err)
			}
			out, err := a.Filter([]fraction.Fraction{half, {}}, journey.Positive)
			if !errors.Is(err, fraction.ErrZeroDenominator) || len(out) != 1 || !out[0].Equal(half) {
				t.Fatal(out, err)
			}
		})
	}
}
