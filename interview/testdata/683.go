package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x any = 7
	v := reflect.ValueOf(&x).Elem()
	fmt.Println(v.Kind(), v.Elem().Kind(), v.CanSet(), v.Elem().CanSet())
}
