package gotypes_test

import (
	"errors"
	"math"
	"reflect"
	"strings"
	"testing"

	"gotypes/boxint"
	"gotypes/fraction"
	"gotypes/intvalue"
	"gotypes/money"
	"gotypes/vector2"
	"gotypes/version"
)

type arithmetic[T any] interface {
	Validate() error
	String() string
	IsZero() bool
	Equal(T) bool
	Compare(T) (int, error)
	Add(T) (T, error)
	Sub(T) (T, error)
	Mul(T) (T, error)
	Div(T) (T, error)
	Neg() (T, error)
	Abs() (T, error)
}

func checkArithmetic[T arithmetic[T]](t *testing.T, value func(int64) T) {
	t.Helper()
	six, three := value(6), value(3)
	operations := []struct {
		name string
		run  func() (T, error)
		want T
	}{
		{"add", func() (T, error) { return six.Add(three) }, value(9)},
		{"subtract", func() (T, error) { return six.Sub(three) }, value(3)},
		{"multiply", func() (T, error) { return six.Mul(three) }, value(18)},
		{"divide", func() (T, error) { return six.Div(three) }, value(2)},
		{"negate", six.Neg, value(-6)},
		{"absolute", value(-6).Abs, value(6)},
	}
	for _, op := range operations {
		t.Run(op.name, func(t *testing.T) {
			got, err := op.run()
			if err != nil || !got.Equal(op.want) {
				t.Fatalf("got %v, %v; want %v", got, err, op.want)
			}
		})
	}
	if !six.Equal(value(6)) || six.Equal(three) || !value(0).IsZero() || six.IsZero() {
		t.Fatal("equality/zero contract")
	}
	if got, err := six.Compare(three); got != 1 || err != nil {
		t.Fatalf("Compare: %d, %v", got, err)
	}
	if got, err := three.Compare(six); got != -1 || err != nil {
		t.Fatalf("Compare: %d, %v", got, err)
	}
	if got, err := six.Compare(six); got != 0 || err != nil {
		t.Fatalf("Compare: %d, %v", got, err)
	}
	if six.Validate() != nil || six.String() != "6" {
		t.Fatal("validation/string")
	}
	if !six.Equal(value(6)) {
		t.Fatal("receiver mutated")
	}
	if _, err := value(math.MaxInt64).Add(value(1)); err == nil {
		t.Fatal("overflow accepted")
	}
	if _, err := value(math.MinInt64).Neg(); err == nil {
		t.Fatal("negation overflow accepted")
	}
	if _, err := six.Div(value(0)); err == nil {
		t.Fatal("division by zero accepted")
	}
}
func TestCommonArithmetic(t *testing.T) {
	t.Run("fraction", func(t *testing.T) {
		checkArithmetic(t, func(n int64) fraction.Fraction { return fraction.Fraction{Numerator: n, Denominator: 1} })
	})
	t.Run("boxint", func(t *testing.T) { checkArithmetic(t, func(n int64) boxint.BoxInt { return boxint.BoxInt{Value: n} }) })
	t.Run("intvalue", func(t *testing.T) {
		checkArithmetic(t, func(n int64) intvalue.IntValue { return intvalue.IntValue(n) })
	})
}

func TestFractionMathematics(t *testing.T) {
	half := fraction.Fraction{Numerator: 1, Denominator: 2}
	equivalent := fraction.Fraction{Numerator: -2, Denominator: -4}
	if !half.Equal(equivalent) {
		t.Fatal("equal values in different representations")
	}
	normalized, err := fraction.New(-2, -4)
	if err != nil || normalized != half {
		t.Fatalf("New: %v, %v", normalized, err)
	}
	sum, err := half.Add(fraction.Fraction{Numerator: 1, Denominator: 3})
	if err != nil || sum != (fraction.Fraction{Numerator: 5, Denominator: 6}) {
		t.Fatalf("Add: %v, %v", sum, err)
	}
	quotient, err := half.Div(fraction.Fraction{Numerator: 3, Denominator: 4})
	if err != nil || quotient != (fraction.Fraction{Numerator: 2, Denominator: 3}) {
		t.Fatalf("Div: %v, %v", quotient, err)
	}
	// Intermediate products exceed int64, but the reduced result fits.
	a := fraction.Fraction{Numerator: math.MaxInt64, Denominator: 2}
	b := fraction.Fraction{Numerator: 2, Denominator: math.MaxInt64}
	product, err := a.Mul(b)
	if err != nil || product.String() != "1" {
		t.Fatalf("cancellation: %v, %v", product, err)
	}
	negative := fraction.Fraction{Numerator: 1, Denominator: math.MinInt64}
	if c, err := negative.Compare(half); err != nil || c != -1 {
		t.Fatalf("large denominator compare: %d, %v", c, err)
	}
	if _, err := fraction.New(1, math.MinInt64); !errors.Is(err, fraction.ErrOverflow) {
		t.Fatalf("canonical denominator overflow: %v", err)
	}
	zero, err := fraction.New(0, math.MinInt64)
	if err != nil || zero.Denominator != 1 || !zero.IsZero() {
		t.Fatalf("normalized zero: %v, %v", zero, err)
	}
	invalid := fraction.Fraction{}
	if !errors.Is(invalid.Validate(), fraction.ErrZeroDenominator) || invalid.Equal(invalid) || invalid.IsZero() {
		t.Fatal("invalid zero struct")
	}
	if _, err := half.Add(invalid); !errors.Is(err, fraction.ErrZeroDenominator) {
		t.Fatalf("invalid operand: %v", err)
	}
	if _, err := invalid.Div(fraction.Fraction{Denominator: 1}); !errors.Is(err, fraction.ErrZeroDenominator) {
		t.Fatalf("invalid receiver: %v", err)
	}
	if !strings.Contains(invalid.String(), "invalid") {
		t.Fatal("invalid string")
	}
}

func TestMoneyCurrencyAndArithmetic(t *testing.T) {
	usd := money.Money{Amount: 1050, Currency: "USD"}
	fee := money.Money{Amount: 25, Currency: "USD"}
	eur := money.Money{Amount: 1050, Currency: "EUR"}
	sum, err := usd.Add(fee)
	if err != nil || sum.Amount != 1075 || sum.Currency != "USD" {
		t.Fatalf("Add: %v, %v", sum, err)
	}
	difference, err := usd.Sub(fee)
	if err != nil || difference.Amount != 1025 {
		t.Fatalf("Sub: %v, %v", difference, err)
	}
	negative, err := usd.Neg()
	if err != nil || negative.Amount != -1050 {
		t.Fatalf("Neg: %v, %v", negative, err)
	}
	absolute, err := negative.Abs()
	if err != nil || !absolute.Equal(usd) {
		t.Fatalf("Abs: %v, %v", absolute, err)
	}
	if usd.Equal(eur) {
		t.Fatal("different currencies equal")
	}
	if _, err := usd.Add(eur); !errors.Is(err, money.ErrCurrencyMismatch) {
		t.Fatal(err)
	}
	if _, err := usd.Compare(eur); !errors.Is(err, money.ErrCurrencyMismatch) {
		t.Fatal(err)
	}
	if c, err := usd.Compare(fee); c != 1 || err != nil {
		t.Fatalf("Compare: %d, %v", c, err)
	}
	if !((money.Money{Currency: "JPY"}).IsZero()) || (money.Money{}).IsZero() {
		t.Fatal("currency-specific zero")
	}
	if usd.String() != "1050 USD (minor units)" {
		t.Fatal(usd.String())
	}
	for _, code := range []string{"", "usd", "US", "USDD", "€UR"} {
		if !errors.Is((money.Money{Currency: code}).Validate(), money.ErrInvalidCurrency) {
			t.Fatalf("accepted %q", code)
		}
	}
	if _, err := (money.Money{Amount: math.MaxInt64, Currency: "USD"}).Add(fee); !errors.Is(err, money.ErrOverflow) {
		t.Fatal(err)
	}
	if usd.Amount != 1050 {
		t.Fatal("receiver mutated")
	}
}

func TestIntegerBoundaries(t *testing.T) {
	cases := []struct {
		name string
		run  func() (intvalue.IntValue, error)
		want intvalue.IntValue
		err  error
	}{
		{"truncation", func() (intvalue.IntValue, error) { return intvalue.IntValue(-7).Div(3) }, -2, nil},
		{"min divided by minus one", func() (intvalue.IntValue, error) { return intvalue.IntValue(math.MinInt64).Div(-1) }, 0, intvalue.ErrOverflow},
		{"subtract overflow", func() (intvalue.IntValue, error) { return intvalue.IntValue(math.MinInt64).Sub(1) }, 0, intvalue.ErrOverflow},
		{"subtract min", func() (intvalue.IntValue, error) { return intvalue.IntValue(math.MinInt64).Sub(math.MinInt64) }, 0, nil},
		{"negative addition overflow", func() (intvalue.IntValue, error) { return intvalue.IntValue(math.MinInt64).Add(-1) }, 0, intvalue.ErrOverflow},
		{"multiply overflow", func() (intvalue.IntValue, error) { return intvalue.IntValue(math.MaxInt64).Mul(2) }, 0, intvalue.ErrOverflow},
		{"absolute overflow", func() (intvalue.IntValue, error) { return intvalue.IntValue(math.MinInt64).Abs() }, 0, intvalue.ErrOverflow},
		{"division by zero", func() (intvalue.IntValue, error) { return intvalue.IntValue(5).Div(0) }, 0, intvalue.ErrDivisionByZero},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := tc.run()
			if got != tc.want || !errors.Is(err, tc.err) {
				t.Fatalf("got %v,%v want %v,%v", got, err, tc.want, tc.err)
			}
		})
	}
}

func TestVectorOperations(t *testing.T) {
	v := vector2.Vector2{X: 3, Y: 4}
	w := vector2.Vector2{X: 1, Y: -2}
	sum, err := v.Add(w)
	if err != nil || sum != (vector2.Vector2{X: 4, Y: 2}) {
		t.Fatal(sum, err)
	}
	diff, err := v.Sub(w)
	if err != nil || diff != (vector2.Vector2{X: 2, Y: 6}) {
		t.Fatal(diff, err)
	}
	neg, err := v.Neg()
	if err != nil || neg != (vector2.Vector2{X: -3, Y: -4}) {
		t.Fatal(neg, err)
	}
	scaled, err := v.Scale(2)
	if err != nil || scaled != (vector2.Vector2{X: 6, Y: 8}) {
		t.Fatal(scaled, err)
	}
	dot, err := v.Dot(w)
	if err != nil || dot != -5 {
		t.Fatal(dot, err)
	}
	norm, err := v.Norm()
	if err != nil || norm != 5 {
		t.Fatal(norm, err)
	}
	if !v.Equal(v) || v.IsZero() || !(vector2.Vector2{}).IsZero() || v.String() != "(3, 4)" {
		t.Fatal("value contract")
	}
	for _, bad := range []float64{math.NaN(), math.Inf(1), math.Inf(-1)} {
		invalid := vector2.Vector2{X: bad}
		if invalid.Validate() == nil || invalid.Equal(invalid) {
			t.Fatal("invalid vector accepted")
		}
		if _, err := v.Scale(bad); !errors.Is(err, vector2.ErrNonFinite) {
			t.Fatal(err)
		}
	}
	if _, err := (vector2.Vector2{X: math.MaxFloat64}).Scale(2); !errors.Is(err, vector2.ErrNonFinite) {
		t.Fatal(err)
	}
	if _, err := (vector2.Vector2{X: math.MaxFloat64, Y: math.MaxFloat64}).Norm(); !errors.Is(err, vector2.ErrNonFinite) {
		t.Fatal(err)
	}
}

func TestVersionOrderingAndBumps(t *testing.T) {
	v := version.Version{Major: 1, Minor: 9, Patch: 7}
	for _, other := range []version.Version{{Major: 2}, {Major: 1, Minor: 10}, {Major: 1, Minor: 9, Patch: 8}} {
		if c, err := v.Compare(other); c != -1 || err != nil {
			t.Fatal(c, err)
		}
	}
	patch, err := v.NextPatch()
	if err != nil || patch != (version.Version{Major: 1, Minor: 9, Patch: 8}) {
		t.Fatal(patch, err)
	}
	minor, err := v.NextMinor()
	if err != nil || minor != (version.Version{Major: 1, Minor: 10}) {
		t.Fatal(minor, err)
	}
	major, err := v.NextMajor()
	if err != nil || major != (version.Version{Major: 2}) {
		t.Fatal(major, err)
	}
	max := version.Version{Major: math.MaxUint32, Minor: math.MaxUint32, Patch: math.MaxUint32}
	for _, bump := range []func() (version.Version, error){max.NextPatch, max.NextMinor, max.NextMajor} {
		if _, err := bump(); !errors.Is(err, version.ErrOverflow) {
			t.Fatal(err)
		}
	}
	if v.String() != "1.9.7" || !v.Equal(v) || v.IsZero() || !(version.Version{}).IsZero() || v.Validate() != nil {
		t.Fatal("value contract")
	}
}

// Signatures are compared modulo each receiver's concrete type T.
func methodShapes(value interface{}) map[string]string {
	self := reflect.TypeOf(value)
	out := map[string]string{}
	name := func(t reflect.Type) string {
		if t == self {
			return "T"
		}
		return t.String()
	}
	for i := 0; i < self.NumMethod(); i++ {
		m := self.Method(i)
		shape := "("
		for j := 1; j < m.Type.NumIn(); j++ {
			shape += name(m.Type.In(j)) + ","
		}
		shape += ")->("
		for j := 0; j < m.Type.NumOut(); j++ {
			shape += name(m.Type.Out(j)) + ","
		}
		out[m.Name] = shape + ")"
	}
	return out
}
func TestMethodOverlap(t *testing.T) {
	baseline := methodShapes(fraction.Fraction{})
	if len(baseline) != 11 {
		t.Fatalf("fraction method count: %d", len(baseline))
	}
	for _, v := range []interface{}{boxint.BoxInt{}, intvalue.IntValue(0)} {
		if !reflect.DeepEqual(baseline, methodShapes(v)) {
			t.Fatalf("%T does not share all eleven signatures", v)
		}
	}
	monetary := methodShapes(money.Money{})
	if len(monetary) != 9 {
		t.Fatalf("money method count: %d", len(monetary))
	}
	common := 0
	for name, shape := range monetary {
		if baseline[name] == shape {
			common++
		}
	}
	union := len(baseline) + len(monetary) - common
	if common*100 < union*60 {
		t.Fatalf("overlap %d/%d is below 60%%", common, union)
	}
	if common != 9 {
		t.Fatalf("expected nine common methods, got %d", common)
	}
}
