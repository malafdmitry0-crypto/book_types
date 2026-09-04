// Package fraction описывает обыкновенную дробь.
package fraction

import (
	"errors"
	"fmt"
	"math/big"

	"gotypes/internal/checkedint"
)

// Fraction хранит числитель и знаменатель.
// Литералы могут быть несокращёнными; New и арифметические методы нормализуют результат.
// Нулевое значение структуры недопустимо: знаменатель должен быть ненулевым.
type Fraction struct {
	Numerator   int64
	Denominator int64
}

var ErrZeroDenominator = errors.New("fraction denominator is zero")
var ErrOverflow = checkedint.ErrOverflow
var ErrDivisionByZero = checkedint.ErrDivisionByZero

// New reduces the fraction and makes its denominator positive.
// Its canonical numerator and denominator must fit in int64.
func New(numerator, denominator int64) (Fraction, error) {
	f := Fraction{numerator, denominator}
	if err := f.Validate(); err != nil {
		return Fraction{}, err
	}
	return fromRat(f.rat())
}

// Validate accepts noncanonical fractions too. The zero struct is invalid.
func (f Fraction) Validate() error {
	if f.Denominator == 0 {
		return ErrZeroDenominator
	}
	return nil
}
func (f Fraction) rat() *big.Rat {
	return new(big.Rat).SetFrac(big.NewInt(f.Numerator), big.NewInt(f.Denominator))
}
func fromRat(r *big.Rat) (Fraction, error) {
	if !r.Num().IsInt64() || !r.Denom().IsInt64() {
		return Fraction{}, ErrOverflow
	}
	return Fraction{r.Num().Int64(), r.Denom().Int64()}, nil
}

// String prints the reduced mathematical value, or a diagnostic for invalid data.
func (f Fraction) String() string {
	if f.Validate() != nil {
		return fmt.Sprintf("invalid fraction(%d/%d)", f.Numerator, f.Denominator)
	}
	return f.rat().RatString()
}
func (f Fraction) IsZero() bool { return f.Validate() == nil && f.Numerator == 0 }

// Equal compares mathematical values. Invalid values are never equal.
func (f Fraction) Equal(other Fraction) bool {
	result, err := f.Compare(other)
	return err == nil && result == 0
}
func (f Fraction) Compare(other Fraction) (int, error) {
	if err := f.Validate(); err != nil {
		return 0, err
	}
	if err := other.Validate(); err != nil {
		return 0, err
	}
	return f.rat().Cmp(other.rat()), nil
}
func (f Fraction) binary(other Fraction, op func(*big.Rat, *big.Rat, *big.Rat) *big.Rat) (Fraction, error) {
	if err := f.Validate(); err != nil {
		return Fraction{}, err
	}
	if err := other.Validate(); err != nil {
		return Fraction{}, err
	}
	return fromRat(op(new(big.Rat), f.rat(), other.rat()))
}
func (f Fraction) Add(other Fraction) (Fraction, error) { return f.binary(other, (*big.Rat).Add) }
func (f Fraction) Sub(other Fraction) (Fraction, error) { return f.binary(other, (*big.Rat).Sub) }
func (f Fraction) Mul(other Fraction) (Fraction, error) { return f.binary(other, (*big.Rat).Mul) }
func (f Fraction) Div(other Fraction) (Fraction, error) {
	if err := f.Validate(); err != nil {
		return Fraction{}, err
	}
	if err := other.Validate(); err != nil {
		return Fraction{}, err
	}
	if other.Numerator == 0 {
		return Fraction{}, ErrDivisionByZero
	}
	return f.binary(other, (*big.Rat).Quo)
}
func (f Fraction) Neg() (Fraction, error) {
	if err := f.Validate(); err != nil {
		return Fraction{}, err
	}
	return fromRat(new(big.Rat).Neg(f.rat()))
}
func (f Fraction) Abs() (Fraction, error) {
	if err := f.Validate(); err != nil {
		return Fraction{}, err
	}
	return fromRat(new(big.Rat).Abs(f.rat()))
}
