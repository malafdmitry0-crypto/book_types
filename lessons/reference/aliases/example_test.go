package aliases

import "fmt"

func Example() {
	aliasExample()
	var n int64 = 42
	var id UserID = n
	fmt.Printf("%T %T\n", id, OrderID(n))
	// Output: int64 aliases.OrderID
}
func Example_typeSwitch() {
	var user UserID = 1
	var order OrderID = 1
	fmt.Println(describe(user), describe(order), describe(int32(1)))
	// Output: int64 OrderID other
}
