package main

import (
	"fmt"
	"reflect"
)

func main() {
	v := reflect.ValueOf([1]int{1})
	fmt.Println(v.Index(0).CanSet())
}
