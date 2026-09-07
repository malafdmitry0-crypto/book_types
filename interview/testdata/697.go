package main

import (
	"fmt"
	"reflect"
)

func main() {
	f := reflect.ValueOf(func(prefix string, xs ...int) string { return fmt.Sprint(prefix, len(xs)) })
	r := f.CallSlice([]reflect.Value{reflect.ValueOf("n="), reflect.ValueOf([]int{2, 4, 6})})
	fmt.Println(r[0].String())
}
