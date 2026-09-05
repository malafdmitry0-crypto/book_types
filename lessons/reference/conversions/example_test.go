package conversions

import "fmt"

func Example() {
	fmt.Println(Convert[int64](int32(42)))
	fmt.Println(Assert[int64](int(42)))
	fmt.Println(IsString("go"), IsString(42))
	fmt.Println(Map([]int{1, 2}, func(n int) int64 { return int64(n) }))
	fmt.Println(Equal[any](42, 42))
	// Output:
	// 42
	// 0 false
	// true false
	// [1 2]
	// true
}
