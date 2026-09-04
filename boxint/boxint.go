// Package boxint описывает целое число, обёрнутое в структуру.
package boxint

import (
	"cmp"
	"strconv"

	"gotypes/internal/checkedint"
)

// BoxInt хранит число в поле; арифметические операторы к самой структуре неприменимы.
type BoxInt struct {
	Value int64
}

var ErrOverflow = checkedint.ErrOverflow
var ErrDivisionByZero = checkedint.ErrDivisionByZero

// Validate always succeeds: every int64 is a valid value.
func (v BoxInt) Validate() error                   { return nil }
func (v BoxInt) String() string                    { return strconv.FormatInt(v.Value, 10) }
func (v BoxInt) IsZero() bool                      { return v.Value == 0 }
func (v BoxInt) Equal(other BoxInt) bool           { return v == other }
func (v BoxInt) Compare(other BoxInt) (int, error) { return cmp.Compare(v.Value, other.Value), nil }

// Add returns a new value; overflow is an error.
func (v BoxInt) Add(other BoxInt) (BoxInt, error) {
	result, err := checkedint.Add(v.Value, other.Value)
	return BoxInt{Value: result}, err
}

// Sub returns a new value; overflow is an error.
func (v BoxInt) Sub(other BoxInt) (BoxInt, error) {
	result, err := checkedint.Sub(v.Value, other.Value)
	return BoxInt{Value: result}, err
}

// Mul returns a new value; overflow is an error.
func (v BoxInt) Mul(other BoxInt) (BoxInt, error) {
	result, err := checkedint.Mul(v.Value, other.Value)
	return BoxInt{Value: result}, err
}

// Div returns a new value; overflow and division by zero is an error. Division truncates toward zero.
func (v BoxInt) Div(other BoxInt) (BoxInt, error) {
	result, err := checkedint.Div(v.Value, other.Value)
	return BoxInt{Value: result}, err
}

func (v BoxInt) Neg() (BoxInt, error) {
	result, err := checkedint.Neg(v.Value)
	return BoxInt{Value: result}, err
}

func (v BoxInt) Abs() (BoxInt, error) {
	result, err := checkedint.Abs(v.Value)
	return BoxInt{Value: result}, err
}
