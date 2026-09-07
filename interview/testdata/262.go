package main

import (
	"fmt"
	"reflect"
)

type A = int
type Box[T any] struct{ V T }

func main() {
	fmt.Println(reflect.TypeOf(Box[int]{}) == reflect.TypeOf(Box[A]{}))
}
