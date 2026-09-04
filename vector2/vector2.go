// Package vector2 описывает двумерный вектор.
package vector2

import (
	"errors"
	"fmt"
	"math"
)

// Vector2 хранит две конечные координаты. Общего порядка у векторов нет;
// критерий сортировки должен выбрать вызывающий код.
type Vector2 struct {
	X float64
	Y float64
}

var ErrNonFinite = errors.New("vector operation requires finite coordinates and result")

func finite(v float64) bool { return !math.IsNaN(v) && !math.IsInf(v, 0) }
func (v Vector2) Validate() error {
	if !finite(v.X) || !finite(v.Y) {
		return ErrNonFinite
	}
	return nil
}
func (v Vector2) String() string { return fmt.Sprintf("(%g, %g)", v.X, v.Y) }
func (v Vector2) IsZero() bool   { return v.Validate() == nil && v.X == 0 && v.Y == 0 }

// Equal uses exact coordinate equality. Approximate equality is a separate policy.
func (v Vector2) Equal(other Vector2) bool {
	return v.Validate() == nil && other.Validate() == nil && v == other
}
func checked(v Vector2) (Vector2, error) {
	if err := v.Validate(); err != nil {
		return Vector2{}, err
	}
	return v, nil
}
func (v Vector2) Add(other Vector2) (Vector2, error) {
	if err := v.Validate(); err != nil {
		return Vector2{}, err
	}
	if err := other.Validate(); err != nil {
		return Vector2{}, err
	}
	return checked(Vector2{v.X + other.X, v.Y + other.Y})
}
func (v Vector2) Sub(other Vector2) (Vector2, error) {
	if err := v.Validate(); err != nil {
		return Vector2{}, err
	}
	if err := other.Validate(); err != nil {
		return Vector2{}, err
	}
	return checked(Vector2{v.X - other.X, v.Y - other.Y})
}
func (v Vector2) Neg() (Vector2, error) {
	if err := v.Validate(); err != nil {
		return Vector2{}, err
	}
	return Vector2{-v.X, -v.Y}, nil
}
func (v Vector2) Scale(scalar float64) (Vector2, error) {
	if err := v.Validate(); err != nil {
		return Vector2{}, err
	}
	if !finite(scalar) {
		return Vector2{}, ErrNonFinite
	}
	return checked(Vector2{v.X * scalar, v.Y * scalar})
}
func (v Vector2) Dot(other Vector2) (float64, error) {
	if err := v.Validate(); err != nil {
		return 0, err
	}
	if err := other.Validate(); err != nil {
		return 0, err
	}
	result := v.X*other.X + v.Y*other.Y
	if !finite(result) {
		return 0, ErrNonFinite
	}
	return result, nil
}
func (v Vector2) Norm() (float64, error) {
	if err := v.Validate(); err != nil {
		return 0, err
	}
	result := math.Hypot(v.X, v.Y)
	if !finite(result) {
		return 0, ErrNonFinite
	}
	return result, nil
}
