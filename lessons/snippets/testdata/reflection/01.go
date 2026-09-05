// Фрагмент 1: book/chapters/08-reflection.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
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
