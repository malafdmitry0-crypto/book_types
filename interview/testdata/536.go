package main

import (
	"fmt"
	"reflect"
)

type N int

func (n N) Add(x int) int { return int(n) + x }
func main() {
	m, _ := reflect.TypeFor[N]().MethodByName("Add")
	out := m.Func.Call([]reflect.Value{reflect.ValueOf(N(3)), reflect.ValueOf(2)})
	fmt.Println(m.Type.NumIn(), out[0].Int())
}
