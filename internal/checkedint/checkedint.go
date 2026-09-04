// Package checkedint implements checked int64 arithmetic shared by value types.
package checkedint

import (
	"errors"
	"math"
	"math/big"
)

var ErrOverflow = errors.New("integer overflow")
var ErrDivisionByZero = errors.New("division by zero")

func Add(a, b int64) (int64, error) {
	if b > 0 && a > math.MaxInt64-b || b < 0 && a < math.MinInt64-b {
		return 0, ErrOverflow
	}
	return a + b, nil
}
func Sub(a, b int64) (int64, error) {
	if b > 0 && a < math.MinInt64+b || b < 0 && a > math.MaxInt64+b {
		return 0, ErrOverflow
	}
	return a - b, nil
}
func Mul(a, b int64) (int64, error) {
	result := new(big.Int).Mul(big.NewInt(a), big.NewInt(b))
	if !result.IsInt64() {
		return 0, ErrOverflow
	}
	return result.Int64(), nil
}

// Div truncates toward zero.
func Div(a, b int64) (int64, error) {
	if b == 0 {
		return 0, ErrDivisionByZero
	}
	if a == math.MinInt64 && b == -1 {
		return 0, ErrOverflow
	}
	return a / b, nil
}
func Neg(a int64) (int64, error) {
	if a == math.MinInt64 {
		return 0, ErrOverflow
	}
	return -a, nil
}
func Abs(a int64) (int64, error) {
	if a < 0 {
		return Neg(a)
	}
	return a, nil
}
