package main

import (
	"fmt"
	"reflect"
)

func main() {
	t := reflect.StructOf([]reflect.StructField{{Name: "X", Type: reflect.TypeFor[int]()}})
	fmt.Println(t == reflect.TypeFor[struct{ X int }]())
}
