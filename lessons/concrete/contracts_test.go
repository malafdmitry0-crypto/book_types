package concrete_test

import (
	"errors"
	"gotypes/fraction"
	"gotypes/lessons/concrete"
	"gotypes/money"
	"testing"
)

func TestContracts(t *testing.T) {
	half := fraction.Fraction{Numerator: 1, Denominator: 2}
	if i, e := concrete.FindFraction(nil, fraction.Fraction{}); i != -1 || e == nil {
		t.Fatal(i, e)
	}
	if i, e := concrete.FindFraction([]fraction.Fraction{half, {}}, half); i != 0 || e != nil {
		t.Fatal(i, e)
	}
	stop := errors.New("stop")
	if i, e := concrete.FindFractionFunc([]fraction.Fraction{half}, func(fraction.Fraction) (bool, error) { return true, stop }); i != -1 || !errors.Is(e, stop) {
		t.Fatal(i, e)
	}
	texts, e := concrete.MapFractionText([]fraction.Fraction{half, {}}, func(v fraction.Fraction) (string, error) { return v.String(), v.Validate() })
	if e == nil || len(texts) != 2 || texts[0] != "1/2" || texts[1] != "" {
		t.Fatal(texts, e)
	}
	total, e := concrete.SumMoney([]money.Money{{Amount: 100, Currency: "USD"}, {Amount: 50, Currency: "EUR"}}, money.Money{Currency: "USD"})
	if !errors.Is(e, money.ErrCurrencyMismatch) || total.Amount != 100 {
		t.Fatal(total, e)
	}
}
