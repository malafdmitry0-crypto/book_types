package main

import (
	"fmt"
	"reflect"
)

func main() {
	a, b := reflect.ArrayOf(2, reflect.TypeFor[int]()), reflect.TypeFor[[2]int]()
	fmt.Println(a == b, a.Len())
}
