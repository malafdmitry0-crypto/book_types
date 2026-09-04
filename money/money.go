// Package money описывает денежную сумму в заданной валюте.
package money

import (
	"cmp"
	"errors"
	"fmt"

	"gotypes/internal/checkedint"
)

// Money хранит сумму в минимальных денежных единицах валюты.
// Например, Amount = 1050 и Currency = "USD" обозначают 10.50 USD.
// Методы проверяют формат кода валюты и запрещают арифметику между разными валютами.
// Конвертация валют и округление не выполняются.
type Money struct {
	Amount   int64
	Currency string
}

var ErrInvalidCurrency = errors.New("currency must be three uppercase ASCII letters")
var ErrCurrencyMismatch = errors.New("currencies do not match")
var ErrOverflow = checkedint.ErrOverflow

// Validate checks currency syntax, not membership in an external currency registry.
func (m Money) Validate() error {
	if len(m.Currency) != 3 {
		return ErrInvalidCurrency
	}
	for i := 0; i < len(m.Currency); i++ {
		if m.Currency[i] < 'A' || m.Currency[i] > 'Z' {
			return ErrInvalidCurrency
		}
	}
	return nil
}

// String explicitly uses minor units; it does not assume two decimal places.
func (m Money) String() string { return fmt.Sprintf("%d %s (minor units)", m.Amount, m.Currency) }
func (m Money) IsZero() bool   { return m.Validate() == nil && m.Amount == 0 }
func (m Money) Equal(other Money) bool {
	return m.Validate() == nil && other.Validate() == nil && m == other
}
func (m Money) compatible(other Money) error {
	if err := m.Validate(); err != nil {
		return err
	}
	if err := other.Validate(); err != nil {
		return err
	}
	if m.Currency != other.Currency {
		return ErrCurrencyMismatch
	}
	return nil
}
func (m Money) Compare(other Money) (int, error) {
	if err := m.compatible(other); err != nil {
		return 0, err
	}
	return cmp.Compare(m.Amount, other.Amount), nil
}
func (m Money) Add(other Money) (Money, error) {
	if err := m.compatible(other); err != nil {
		return Money{}, err
	}
	amount, err := checkedint.Add(m.Amount, other.Amount)
	if err != nil {
		return Money{}, err
	}
	return Money{amount, m.Currency}, nil
}
func (m Money) Sub(other Money) (Money, error) {
	if err := m.compatible(other); err != nil {
		return Money{}, err
	}
	amount, err := checkedint.Sub(m.Amount, other.Amount)
	if err != nil {
		return Money{}, err
	}
	return Money{amount, m.Currency}, nil
}
func (m Money) Neg() (Money, error) {
	if err := m.Validate(); err != nil {
		return Money{}, err
	}
	amount, err := checkedint.Neg(m.Amount)
	if err != nil {
		return Money{}, err
	}
	return Money{amount, m.Currency}, nil
}
func (m Money) Abs() (Money, error) {
	if err := m.Validate(); err != nil {
		return Money{}, err
	}
	amount, err := checkedint.Abs(m.Amount)
	if err != nil {
		return Money{}, err
	}
	return Money{amount, m.Currency}, nil
}
