package laws_test

import (
	"errors"
	"gotypes/fraction"
	"math"
	"testing"
)

// Ограниченная область гарантирует представимость промежуточных сумм.
func bounded(n, d int64) fraction.Fraction {
	denominator := d % 97
	if denominator < 0 {
		denominator = -denominator
	}
	return fraction.Fraction{Numerator: n % 1000, Denominator: denominator + 1}
}

func FuzzFractionLaws(f *testing.F) {
	f.Add(int64(1), int64(2), int64(2), int64(4), int64(-1), int64(6))
	f.Add(int64(math.MinInt64), int64(math.MaxInt64), int64(0), int64(0), int64(3), int64(-4))
	f.Fuzz(func(t *testing.T, an, ad, bn, bd, cn, cd int64) {
		a, b, c := bounded(an, ad), bounded(bn, bd), bounded(cn, cd)
		if !a.Equal(a) || a.Equal(b) != b.Equal(a) || (a.Equal(b) && b.Equal(c) && !a.Equal(c)) {
			t.Fatal("equality law")
		}
		ab, e1 := a.Compare(b)
		ba, e2 := b.Compare(a)
		bc, e3 := b.Compare(c)
		ac, e4 := a.Compare(c)
		if e1 != nil || e2 != nil || e3 != nil || e4 != nil {
			t.Fatal("bounded inputs must be valid")
		}
		if ab != -ba || (ab <= 0 && bc <= 0 && ac > 0) {
			t.Fatal("order law")
		}
		s1, e1 := a.Add(b)
		s2, e2 := b.Add(a)
		if e1 != nil || e2 != nil || !s1.Equal(s2) {
			t.Fatal("commutativity")
		}
		left, e1 := s1.Add(c)
		middle, e2 := b.Add(c)
		right, e3 := a.Add(middle)
		if e1 != nil || e2 != nil || e3 != nil || !left.Equal(right) {
			t.Fatal("bounded associativity")
		}
		same, e1 := a.Add(fraction.Fraction{Denominator: 1})
		if e1 != nil || !same.Equal(a) {
			t.Fatal("identity")
		}
	})
}

func TestOverflowChangesGrouping(t *testing.T) {
	a := fraction.Fraction{Numerator: math.MaxInt64, Denominator: 1}
	one := fraction.Fraction{Numerator: 1, Denominator: 1}
	minus := fraction.Fraction{Numerator: -1, Denominator: 1}
	if _, err := a.Add(one); !errors.Is(err, fraction.ErrOverflow) {
		t.Fatal(err)
	}
	inner, err := one.Add(minus)
	if err != nil {
		t.Fatal(err)
	}
	result, err := a.Add(inner)
	if err != nil || !result.Equal(a) {
		t.Fatal(result, err)
	}
}
