package main

import (
	"fmt"
	"reflect"
)

type T int

func (T) F()      {}
func (T) hidden() {}
func main() {
	fmt.Println(reflect.TypeFor[T]().NumMethod())
}
