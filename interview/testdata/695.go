package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf([]int{1})
	t := reflect.TypeOf((*[2]int)(nil))
	fmt.Println(v.Type().ConvertibleTo(t), v.CanConvert(t))
}
