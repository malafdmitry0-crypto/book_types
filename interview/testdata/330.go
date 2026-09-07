package main

import (
	"fmt"
	"reflect"
)

type I interface {
	F()
	hidden()
}

func main() {
	fmt.Println(reflect.TypeFor[I]().NumMethod())
}
