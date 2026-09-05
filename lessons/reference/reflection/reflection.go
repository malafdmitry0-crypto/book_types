package reflection

import (
	"reflect"
)

// import "reflect"
func convertDynamic(x any, target reflect.Type) (any, bool) {
	value := reflect.ValueOf(x)
	if target == nil || !value.IsValid() || !value.CanConvert(target) {
		return nil, false
	}
	return value.Convert(target).Interface(), true
}

// out, ok := convertDynamic(int32(42), reflect.TypeOf(int64(0)))
// out имеет статический тип any и динамический тип int64.
