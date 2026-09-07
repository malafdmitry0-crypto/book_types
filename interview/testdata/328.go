package main

import (
	"fmt"
	"reflect"
)

type T int

func (*T) F() {}
func main() {
	i := reflect.TypeFor[interface{ F() }]()
	fmt.Println(reflect.TypeFor[T]().Implements(i), reflect.TypeFor[*T]().Implements(i))
}
