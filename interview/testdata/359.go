package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := make([]int, 1, 3)
	reflect.ValueOf(&s).Elem().SetLen(2)
	fmt.Println(s)
}
