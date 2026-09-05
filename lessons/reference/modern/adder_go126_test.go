//go:build go1.26

package modern

import "fmt"

type Number int

func (n Number) Add(other Number) Number { return n + other }
func sum[A Adder[A]](a, b A) A           { return a.Add(b) }
func Example_selfConstraint() {
	fmt.Println(sum(Number(1), Number(2)))
	// Output: 3
}
