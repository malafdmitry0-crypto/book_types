package generic_test

import (
	"errors"
	"gotypes/fraction"
	"gotypes/lessons/generic"
	"gotypes/money"
	"testing"
)

func TestPartialResults(t *testing.T) {
	stop := errors.New("stop")
	calls := 0
	values, err := generic.MapValues([]int{1, 2, 3}, func(n int) (int, error) {
		calls++
		if n == 2 {
			return 99, stop
		}
		return n * 10, nil
	})
	if !errors.Is(err, stop) || calls != 2 || len(values) != 3 || values[0] != 10 || values[1] != 0 || values[2] != 0 {
		t.Fatal(values, err, calls)
	}
	var empty []int
	values, err = generic.MapValues(empty, func(n int) (int, error) { t.Fatal("unexpected callback"); return n, nil })
	if values != nil || err != nil {
		t.Fatal(values, err)
	}
	total, err := generic.SumValues([]money.Money{{Amount: 100, Currency: "USD"}, {Amount: 50, Currency: "EUR"}}, money.Money{Currency: "USD"})
	if !errors.Is(err, money.ErrCurrencyMismatch) || total.Amount != 100 {
		t.Fatal(total, err)
	}
	initial := fraction.Fraction{Numerator: 2, Denominator: 3}
	result, err := generic.SumValues([]fraction.Fraction(nil), initial)
	if err != nil || !result.Equal(initial) {
		t.Fatal(result, err)
	}
}
