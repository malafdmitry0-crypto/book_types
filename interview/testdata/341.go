package main

import (
	"fmt"
	"reflect"
)

func main() {
	n := 1
	fmt.Println(reflect.ValueOf(n).CanSet(), reflect.ValueOf(&n).Elem().CanSet())
}
