package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.Zero(reflect.TypeFor[any]())
	fmt.Println(v.IsValid(), v.IsNil(), v.Kind())
}
