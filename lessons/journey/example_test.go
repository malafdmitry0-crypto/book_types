package journey_test

import (
	"fmt"
	"gotypes/fraction"
	"gotypes/lessons/journey"
	"strings"
)

func Example() {
	values := []fraction.Fraction{{Numerator: 1, Denominator: 3}, {Numerator: 2, Denominator: 4}, {Numerator: -1, Denominator: 6}}
	zero := fraction.Fraction{Denominator: 1}
	for _, approach := range journey.Approaches() {
		index, err := approach.Find(values, journey.IsHalf)
		if err != nil {
			fmt.Println(err)
			return
		}
		total, err := approach.Sum(values, zero)
		if err != nil {
			fmt.Println(err)
			return
		}
		selected, err := approach.Filter(values, journey.Positive)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s: Find=%d Sum=%s Filter=%v\n", approach.Name, index, total, selected)
	}
	// Output:
	// concrete: Find=1 Sum=2/3 Filter=[1/3 1/2]
	// closure: Find=1 Sum=2/3 Filter=[1/3 1/2]
	// interface: Find=1 Sum=2/3 Filter=[1/3 1/2]
	// empty-interface: Find=1 Sum=2/3 Filter=[1/3 1/2]
	// reflection: Find=1 Sum=2/3 Filter=[1/3 1/2]
	// generation: Find=1 Sum=2/3 Filter=[1/3 1/2]
	// generic: Find=1 Sum=2/3 Filter=[1/3 1/2]
	// iterator: Find=1 Sum=2/3 Filter=[1/3 1/2]
}

func ExampleGenericReduce() {
	values := []fraction.Fraction{{Numerator: 1, Denominator: 2}, {Numerator: 1, Denominator: 3}}
	text, err := journey.GenericReduce(values, "", func(total string, value fraction.Fraction) (string, error) {
		if err := value.Validate(); err != nil {
			return "", err
		}
		return strings.TrimPrefix(total+", "+value.String(), ", "), nil
	})
	fmt.Println(text, err)
	// Output: 1/2, 1/3 <nil>
}
