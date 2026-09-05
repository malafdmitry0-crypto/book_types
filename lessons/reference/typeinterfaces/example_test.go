package typeinterfaces

import (
	"bytes"
	"fmt"
)

func Example_boxing() {
	var x any = int(42)           // any доступен с Go 1.18
	var old interface{} = int(42) // историческая запись того же типа
	fmt.Println(x, old)
	// Output: 42 42
}
func Example_assertion() {
	var x any = int(42)
	n := x.(int)          // 42
	wide, ok := x.(int64) // 0, false
	// panicValue := x.(int64) // panic

	if n, ok := x.(int); ok {
		converted := int64(n) // assertion, затем числовая конверсия
		_ = converted
	}
	fmt.Println(n, wide, ok)
	// Output: 42 0 false
}
func Example_methods() {
	inspect(new(bytes.Buffer))
	fmt.Println(describe(nil), describe(42), describe("go"), describe(1.5), nilExample())
	// Output: nil interface int: 42 string: go float: 1.5 false
}
