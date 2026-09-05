// Package journey compares eight implementations under one Fraction contract.
package journey

import (
	"fmt"
	"gotypes/algorithms"
	"gotypes/algorithms/closures"
	"gotypes/fraction"
	"slices"
)

type Predicate func(fraction.Fraction) (bool, error)
type Step func(fraction.Fraction, fraction.Fraction) (fraction.Fraction, error)

// Approach — только общий стенд вызова; сами алгоритмы реализованы отдельно.
type Approach struct {
	Name   string
	Find   func([]fraction.Fraction, Predicate) (int, error)
	Filter func([]fraction.Fraction, Predicate) ([]fraction.Fraction, error)
	Reduce func([]fraction.Fraction, fraction.Fraction, Step) (fraction.Fraction, error)
}

func (a Approach) Sum(src []fraction.Fraction, initial fraction.Fraction) (fraction.Fraction, error) {
	return a.Reduce(src, initial, Add)
}
func Add(a, b fraction.Fraction) (fraction.Fraction, error) { return a.Add(b) }
func IsHalf(value fraction.Fraction) (bool, error) {
	if err := value.Validate(); err != nil {
		return false, err
	}
	return value.Equal(fraction.Fraction{Numerator: 1, Denominator: 2}), nil
}
func Positive(value fraction.Fraction) (bool, error) {
	order, err := value.Compare(fraction.Fraction{Denominator: 1})
	return order > 0, err
}
func emptyResult(src []fraction.Fraction) []fraction.Fraction {
	if src == nil {
		return nil
	}
	return make([]fraction.Fraction, 0, len(src))
}
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

func Approaches() []Approach {
	return []Approach{
		{Name: "concrete", Find: func(s []fraction.Fraction, p Predicate) (int, error) { return ConcreteFind(s, p) }, Filter: func(s []fraction.Fraction, p Predicate) ([]fraction.Fraction, error) { return ConcreteFilter(s, p) }, Reduce: func(s []fraction.Fraction, a fraction.Fraction, f Step) (fraction.Fraction, error) {
			return ConcreteReduce(s, a, f)
		}},
		{Name: "closure", Find: func(s []fraction.Fraction, p Predicate) (int, error) {
			return closures.FindIndex(len(s), func(i int) (bool, error) { return p(s[i]) })
		},
			Filter: func(s []fraction.Fraction, p Predicate) ([]fraction.Fraction, error) {
				out := emptyResult(s)
				err := ClosureFilter(len(s), func(i int) error {
					ok, err := p(s[i])
					if err != nil {
						return err
					}
					if ok {
						out = append(out, s[i])
					}
					return nil
				})
				return out, err
			},
			Reduce: func(s []fraction.Fraction, a fraction.Fraction, f Step) (fraction.Fraction, error) {
				total := a
				err := ClosureReduce(len(s), func(i int) error {
					next, err := f(total, s[i])
					if err != nil {
						return err
					}
					total = next
					return nil
				})
				return total, err
			}},
		{Name: "interface", Find: func(s []fraction.Fraction, p Predicate) (int, error) {
			return algorithms.Find(Search{Values: s, Test: p})
		},
			Filter: func(s []fraction.Fraction, p Predicate) ([]fraction.Fraction, error) {
				job := Selection{Values: s, Test: p, Result: emptyResult(s)}
				err := InterfaceFilter(&job)
				return job.Result, err
			},
			Reduce: func(s []fraction.Fraction, a fraction.Fraction, f Step) (fraction.Fraction, error) {
				job := Fold{Values: s, Total: a, Step: f}
				err := algorithms.Sum(&job)
				return job.Total, err
			}},
		{Name: "empty-interface", Find: func(s []fraction.Fraction, p Predicate) (int, error) { return AnyFind(Box(s), AnyPredicate(p)) },
			Filter: func(s []fraction.Fraction, p Predicate) ([]fraction.Fraction, error) {
				raw, err := AnyFilter(Box(s), AnyPredicate(p))
				out := emptyResult(s)
				for _, value := range raw {
					v, e := unbox(value)
					if e != nil {
						return out, e
					}
					out = append(out, v)
				}
				return out, err
			},
			Reduce: func(s []fraction.Fraction, a fraction.Fraction, f Step) (fraction.Fraction, error) {
				raw, err := AnyReduce(Box(s), a, anyStep(f))
				result, e := unbox(raw)
				if e != nil {
					return a, e
				}
				return result, err
			}},
		{Name: "reflection", Find: func(s []fraction.Fraction, p Predicate) (int, error) { return ReflectFind(s, p) },
			Filter: func(s []fraction.Fraction, p Predicate) ([]fraction.Fraction, error) {
				raw, err := ReflectFilter(s, p)
				if raw == nil {
					return nil, err
				}
				out, ok := raw.([]fraction.Fraction)
				if !ok {
					return nil, fmt.Errorf("unexpected filter result %T", raw)
				}
				return out, err
			},
			Reduce: func(s []fraction.Fraction, a fraction.Fraction, f Step) (fraction.Fraction, error) {
				raw, err := ReflectReduce(s, a, f)
				result, e := unbox(raw)
				if e != nil {
					return a, e
				}
				return result, err
			}},
		{Name: "generation", Find: func(s []fraction.Fraction, p Predicate) (int, error) { return GeneratedFind(s, p) }, Filter: func(s []fraction.Fraction, p Predicate) ([]fraction.Fraction, error) { return GeneratedFilter(s, p) }, Reduce: func(s []fraction.Fraction, a fraction.Fraction, f Step) (fraction.Fraction, error) {
			return GeneratedReduce(s, a, f)
		}},
		{Name: "generic", Find: func(s []fraction.Fraction, p Predicate) (int, error) { return GenericFind(s, p) }, Filter: func(s []fraction.Fraction, p Predicate) ([]fraction.Fraction, error) { return GenericFilter(s, p) }, Reduce: func(s []fraction.Fraction, a fraction.Fraction, f Step) (fraction.Fraction, error) {
			return GenericReduce(s, a, f)
		}},
		{Name: "iterator", Find: func(s []fraction.Fraction, p Predicate) (int, error) { return IteratorFind(slices.Values(s), p) },
			Filter: func(s []fraction.Fraction, p Predicate) ([]fraction.Fraction, error) {
				out := emptyResult(s)
				for value, err := range IteratorFilter(slices.Values(s), p) {
					if err != nil {
						return out, err
					}
					out = append(out, value)
				}
				return out, nil
			},
			Reduce: func(s []fraction.Fraction, a fraction.Fraction, f Step) (fraction.Fraction, error) {
				return IteratorReduce(slices.Values(s), a, f)
			}},
	}
}
