//go:build go1.27

package modern

import "fmt"

func Example_genericMethod() {
	fmt.Println((Converter{}).ToInt64(int32(42)))
	// Output: 42
}
