package main

import (
	"fmt"
	"reflect"
)

type N int

func main() {
	a, b := reflect.TypeFor[N](), reflect.TypeFor[int]()
	fmt.Println(a.Kind() == b.Kind(), a.AssignableTo(b), a.ConvertibleTo(b))
}
