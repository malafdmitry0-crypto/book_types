package aliases

import "fmt"

func Example() {
	aliasExample()
	var n int64 = 42
	var id UserID = n
	fmt.Printf("%T %T\n", id, OrderID(n))
	// Output: int64 aliases.OrderID
}
