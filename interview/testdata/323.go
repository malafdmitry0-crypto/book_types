package main

import (
	"fmt"
	"reflect"
)

type N int

func main() {
	a, b := reflect.TypeOf(N(1)), reflect.TypeOf(1)
	fmt.Println(a.Kind() == b.Kind(), a == b, a.Name())
}
