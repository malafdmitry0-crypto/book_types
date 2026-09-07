package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := func(x int) {}
	b := func(y int) {}
	fmt.Println(reflect.TypeOf(a) == reflect.TypeOf(b))
}
