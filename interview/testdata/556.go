package main

import (
	"fmt"
	"reflect"
)

func main() {
	s := []int{1, 2}
	a := reflect.ValueOf(s).Convert(reflect.TypeFor[[2]int]()).Interface().([2]int)
	s[0] = 9
	fmt.Println(a)
}
