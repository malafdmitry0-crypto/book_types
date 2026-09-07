package main

import (
	"fmt"
	"reflect"
)

type N int

func (n N) Add(x int) int { return int(n) + x }
func main() {
	v := reflect.ValueOf(N(3)).MethodByName("Add")
	fmt.Println(v.Type().NumIn(), v.Call([]reflect.Value{reflect.ValueOf(2)})[0].Int())
}
