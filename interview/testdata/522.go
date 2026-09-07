package main

import (
	"fmt"
	"reflect"
)

func main() {
	f := reflect.ValueOf(func(xs ...int) int { return len(xs) })
	out := f.Call([]reflect.Value{reflect.ValueOf(1), reflect.ValueOf(2)})
	fmt.Println(out[0].Int())
}
