package main

import (
	"fmt"
	"reflect"
)

type N struct{}

func (N) A()  {}
func (*N) B() {}
func main() {
	fmt.Println(reflect.TypeOf(N{}).NumMethod(), reflect.TypeOf(&N{}).NumMethod())
}
