package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a any = 1
	fmt.Println(reflect.TypeOf(a).Kind(), reflect.TypeFor[any]().Kind())
}
