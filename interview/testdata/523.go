package main

import (
	"fmt"
	"reflect"
)

func main() {
	f := reflect.ValueOf(func(xs ...int) int { return len(xs) })
	out := f.CallSlice([]reflect.Value{reflect.ValueOf([]int{1, 2})})
	fmt.Println(out[0].Int())
}
