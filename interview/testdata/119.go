package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a []int
	b := []int{}
	fmt.Println(reflect.DeepEqual(a, b))
}
