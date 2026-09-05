package reflection

import (
	"fmt"
	"reflect"
)

func Example() {
	fmt.Println(convertDynamic(int32(42), reflect.TypeOf(int64(0))))
	fmt.Println(convertDynamic([]int{1}, reflect.TypeOf([2]int{})))
	fmt.Println(convertDynamic(nil, reflect.TypeOf(0)))
	value := reflect.ValueOf(int32(42))
	fmt.Println(value.Int(), value.Type(), value.CanSet())
	n := 1
	reflect.ValueOf(&n).Elem().SetInt(7)
	fmt.Println(n)
	fmt.Println(reflect.TypeFor[interface{}]().Kind())
	// Output:
	// 42 true
	// <nil> false
	// <nil> false
	// 42 int32 false
	// 7
	// interface
}
