// Package intvalue описывает отдельный тип на основе целого числа.
package intvalue

import (
	"cmp"
	"strconv"

	"gotypes/internal/checkedint"
)

// IntValue — определённый числовой тип, а не алиас и не структура.
// В отличие от boxint.BoxInt, поддерживает встроенные числовые операции.
type IntValue int64

var ErrOverflow = checkedint.ErrOverflow
var ErrDivisionByZero = checkedint.ErrDivisionByZero

// Validate always succeeds: every int64 is a valid value.
func (v IntValue) Validate() error           { return nil }
func (v IntValue) String() string            { return strconv.FormatInt(int64(v), 10) }
func (v IntValue) IsZero() bool              { return int64(v) == 0 }
func (v IntValue) Equal(other IntValue) bool { return v == other }
func (v IntValue) Compare(other IntValue) (int, error) {
	return cmp.Compare(int64(v), int64(other)), nil
}

// Add returns a new value; overflow is an error.
func (v IntValue) Add(other IntValue) (IntValue, error) {
	result, err := checkedint.Add(int64(v), int64(other))
	return IntValue(result), err
}

// Sub returns a new value; overflow is an error.
func (v IntValue) Sub(other IntValue) (IntValue, error) {
	result, err := checkedint.Sub(int64(v), int64(other))
	return IntValue(result), err
}

// Mul returns a new value; overflow is an error.
func (v IntValue) Mul(other IntValue) (IntValue, error) {
	result, err := checkedint.Mul(int64(v), int64(other))
	return IntValue(result), err
}

// Div returns a new value; overflow and division by zero is an error. Division truncates toward zero.
func (v IntValue) Div(other IntValue) (IntValue, error) {
	result, err := checkedint.Div(int64(v), int64(other))
	return IntValue(result), err
}

func (v IntValue) Neg() (IntValue, error) {
	result, err := checkedint.Neg(int64(v))
	return IntValue(result), err
}

func (v IntValue) Abs() (IntValue, error) {
	result, err := checkedint.Abs(int64(v))
	return IntValue(result), err
}
