package main

import (
	"fmt"
	"reflect"
)

type N int

func main() {
	a, b := reflect.TypeOf([]N{}), reflect.TypeOf([]int{})
	fmt.Println(a.ConvertibleTo(b))
}
