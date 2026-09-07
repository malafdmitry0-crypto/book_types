package main

import (
	"fmt"
	"reflect"
)

type T int

func (T) F() {}
func main() {
	m, _ := reflect.TypeFor[T]().MethodByName("F")
	fmt.Println(m.Type.NumIn(), m.Type.In(0))
}
