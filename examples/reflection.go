package algorithms

import (
	"fmt"
	"reflect"
)

// checkTransform intentionally requires an exact element/argument type match.
// Callback panics are allowed to propagate; they are not validation errors.
func checkTransform(src, fn interface{}) (reflect.Value, reflect.Value, error) {
	s, f := reflect.ValueOf(src), reflect.ValueOf(fn)
	if !s.IsValid() || s.Kind() != reflect.Slice {
		return s, f, fmt.Errorf("source must be a slice")
	}
	if !f.IsValid() || f.Kind() != reflect.Func || f.IsNil() {
		return s, f, fmt.Errorf("callback must be a non-nil function")
	}
	t := f.Type()
	if t.IsVariadic() || t.NumIn() != 1 || t.NumOut() != 1 || t.In(0) != s.Type().Elem() {
		return s, f, fmt.Errorf("callback must have signature func(%v) R", s.Type().Elem())
	}
	return s, f, nil
}

func checkPredicate(src, pred interface{}) (reflect.Value, reflect.Value, error) {
	s, f, err := checkTransform(src, pred)
	if err != nil {
		return s, f, err
	}
	if f.Type().Out(0) != reflect.TypeOf(true) {
		return s, f, fmt.Errorf("predicate must return bool")
	}
	return s, f, nil
}

func FindReflect(src, pred interface{}) (int, error) {
	s, f, err := checkPredicate(src, pred)
	if err != nil {
		return -1, err
	}
	for i := 0; i < s.Len(); i++ {
		if f.Call([]reflect.Value{s.Index(i)})[0].Bool() {
			return i, nil
		}
	}
	return -1, nil
}

func AllReflect(src, pred interface{}) (bool, error) {
	s, f, err := checkPredicate(src, pred)
	if err != nil {
		return false, err
	}
	for i := 0; i < s.Len(); i++ {
		if !f.Call([]reflect.Value{s.Index(i)})[0].Bool() {
			return false, nil
		}
	}
	return true, nil
}

func MapReflect(src, transform interface{}) (interface{}, error) {
	s, f, err := checkTransform(src, transform)
	if err != nil {
		return nil, err
	}
	target := reflect.SliceOf(f.Type().Out(0))
	if s.IsNil() {
		return reflect.Zero(target).Interface(), nil
	}
	out := reflect.MakeSlice(target, s.Len(), s.Len())
	for i := 0; i < s.Len(); i++ {
		out.Index(i).Set(f.Call([]reflect.Value{s.Index(i)})[0])
	}
	return out.Interface(), nil
}
