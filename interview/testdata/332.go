package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf([]int{1})
	t := reflect.TypeFor[*[2]int]()
	fmt.Println(v.Type().ConvertibleTo(t), v.CanConvert(t))
}
