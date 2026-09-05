package interfaces

import (
	"fmt"
	"gotypes/fraction"
)

type Stringer interface {
	String() string
}

func PrintValue(value Stringer) {
	fmt.Println(value.String())
}
func DescribeUnknown(value interface{}) (string, error) {
	stringer, ok := value.(Stringer)
	if !ok {
		return "", fmt.Errorf("тип %T не предоставляет String() string", value)
	}
	return stringer.String(), nil
}

func DescribeStringer(value Stringer) string {
	return value.String()
}

type Addable interface {
	Add(Addable) (Addable, error)
}

// Не компилируется: у Fraction другая сигнатура Add.
// var _ Addable = fraction.Fraction{}
type FractionAdder interface {
	Add(fraction.Fraction) (fraction.Fraction, error)
}

func AddFraction(left FractionAdder, right fraction.Fraction) (fraction.Fraction, error) {
	return left.Add(right)
}

var _ FractionAdder = fraction.Fraction{}

func PrintAll(values []Stringer) {
	for _, value := range values {
		fmt.Println(value.String())
	}
}
