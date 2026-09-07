package main

import (
	"fmt"
	"reflect"
)

func main() {
	a, b := reflect.TypeOf(int32(1)), reflect.TypeOf(int64(1))
	fmt.Println(a.AssignableTo(b), a.ConvertibleTo(b))
}
