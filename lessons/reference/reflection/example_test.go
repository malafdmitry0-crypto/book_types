package reflection

import (
	"fmt"
	"reflect"
)

func Example() {
	fmt.Println(convertDynamic(int32(42), reflect.TypeOf(int64(0))))
	fmt.Println(convertDynamic([]int{1}, reflect.TypeOf([2]int{})))
	fmt.Println(convertDynamic(nil, reflect.TypeOf(0)))
	value := reflect.ValueOf(int32(42))
	fmt.Println(value.Int(), value.Type(), value.CanSet())
	n := 1
	reflect.ValueOf(&n).Elem().SetInt(7)
	fmt.Println(n)
	fmt.Println(reflect.TypeFor[interface{}]().Kind())
	// Output:
	// 42 true
	// <nil> false
	// <nil> false
	// 42 int32 false
	// 7
	// interface
}
func Example_kind() {
	typ := reflect.TypeOf(Label("x"))
	fmt.Println(typ, typ.Kind(), typ.Name(), typ == reflect.TypeOf(""))

	var label Label
	direct := reflect.ValueOf(label)
	viaPointer := reflect.ValueOf(&label).Elem()
	fmt.Println(direct.CanSet(), viaPointer.CanSet())
	viaPointer.SetString("changed")
	fmt.Println(label)
	// Output:
	// reflection.Label string Label false
	// false true
	// changed
}

type User struct {
	Age int
}

func Example_interfaceDiagram() {
	u := User{Age: 30}
	var x any = u
	t, v := reflect.TypeOf(x), reflect.ValueOf(x)
	fmt.Println(t.Name(), t.Kind(), t.Field(0).Name, t.Field(0).Type)
	fmt.Println(v.Field(0).Int(), v.CanSet())
	back := v.Interface()
	copy, ok := back.(User)
	fmt.Println(copy.Age, ok)
	// Output:
	// User struct Age int
	// 30 false
	// 30 true
}

func Example_setDiagram() {
	u := User{Age: 30}
	direct := reflect.ValueOf(u)
	field := reflect.ValueOf(&u).Elem().FieldByName("Age")
	fmt.Println(direct.CanSet(), field.CanSet())
	field.SetInt(31)
	fmt.Println(u.Age, direct.Field(0).Int())
	// Output:
	// false true
	// 31 30
}

func Example_interfaceSlot() {
	var x any = int32(42)
	value := reflect.ValueOf(x)
	slot := reflect.ValueOf(&x).Elem()
	inside := slot.Elem()
	fmt.Println(value.Kind(), value.CanSet())
	fmt.Println(slot.Kind(), slot.CanSet())
	fmt.Println(inside.Kind(), inside.CanSet())
	fmt.Println(slot.Type() == reflect.TypeFor[any]())
	slot.Set(reflect.ValueOf("hello"))
	fmt.Printf("%T %v\n", x, x)
	// Output:
	// int32 false
	// interface true
	// int32 false
	// true
	// string hello
}

func Example_nilReflection() {
	fmt.Println(reflect.TypeOf(nil) == nil, reflect.ValueOf(nil).IsValid())
	var pointer *User
	v := reflect.ValueOf(pointer)
	fmt.Println(v.IsValid(), v.Kind() == reflect.Ptr, v.IsNil())
	// Output:
	// true false
	// true true true
}
