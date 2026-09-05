package journey

import (
	"fmt"
	"reflect"
)

var errorType = reflect.TypeOf((*error)(nil)).Elem()

// Сначала проверяется весь договор callback, даже для пустого среза.
func checkPredicate(src, test interface{}) (reflect.Value, reflect.Value, error) {
	s, f := reflect.ValueOf(src), reflect.ValueOf(test)
	if !s.IsValid() || s.Kind() != reflect.Slice {
		return s, f, fmt.Errorf("source must be a slice")
	}
	if !f.IsValid() || f.Kind() != reflect.Func || f.IsNil() {
		return s, f, fmt.Errorf("predicate must be a non-nil function")
	}
	t := f.Type()
	if t.IsVariadic() || t.NumIn() != 1 || t.NumOut() != 2 || t.In(0) != s.Type().Elem() || t.Out(0) != reflect.TypeOf(false) || t.Out(1) != errorType {
		return s, f, fmt.Errorf("predicate must have signature func(E) (bool, error)")
	}
	return s, f, nil
}
func reflectedError(v reflect.Value) error {
	if v.IsNil() {
		return nil
	}
	return v.Interface().(error)
}
func ReflectFind(src, test interface{}) (int, error) {
	s, f, err := checkPredicate(src, test)
	if err != nil {
		return -1, err
	}
	for i := 0; i < s.Len(); i++ {
		r := f.Call([]reflect.Value{s.Index(i)})
		if err := reflectedError(r[1]); err != nil {
			return -1, err
		}
		if r[0].Bool() {
			return i, nil
		}
	}
	return -1, nil
}
func ReflectFilter(src, test interface{}) (interface{}, error) {
	s, f, err := checkPredicate(src, test)
	if err != nil {
		return nil, err
	}
	if s.IsNil() {
		return reflect.Zero(s.Type()).Interface(), nil
	}
	out := reflect.MakeSlice(s.Type(), 0, s.Len())
	for i := 0; i < s.Len(); i++ {
		r := f.Call([]reflect.Value{s.Index(i)})
		if err := reflectedError(r[1]); err != nil {
			return out.Interface(), err
		}
		if r[0].Bool() {
			out = reflect.Append(out, s.Index(i))
		}
	}
	return out.Interface(), nil
}
func ReflectReduce(src, initial, step interface{}) (interface{}, error) {
	s, a, f := reflect.ValueOf(src), reflect.ValueOf(initial), reflect.ValueOf(step)
	if !s.IsValid() || s.Kind() != reflect.Slice || !a.IsValid() {
		return initial, fmt.Errorf("source must be a slice and initial must have a concrete type")
	}
	if !f.IsValid() || f.Kind() != reflect.Func || f.IsNil() {
		return initial, fmt.Errorf("step must be a non-nil function")
	}
	t := f.Type()
	if t.IsVariadic() || t.NumIn() != 2 || t.NumOut() != 2 || t.In(0) != a.Type() || t.In(1) != s.Type().Elem() || t.Out(0) != a.Type() || t.Out(1) != errorType {
		return initial, fmt.Errorf("step must have signature func(A,E) (A,error)")
	}
	for i := 0; i < s.Len(); i++ {
		r := f.Call([]reflect.Value{a, s.Index(i)})
		if err := reflectedError(r[1]); err != nil {
			return a.Interface(), err
		}
		a = r[0]
	}
	return a.Interface(), nil
}
