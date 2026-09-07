package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := reflect.TypeOf(struct {
		X int `a:"1"`
	}{})
	b := reflect.TypeOf(struct {
		X int `a:"2"`
	}{})
	fmt.Println(a == b, a.AssignableTo(b), a.ConvertibleTo(b))
}
