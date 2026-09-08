// Фрагмент 3: book/chapters/08-reflection.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
var x any = int32(42)
value := reflect.ValueOf(x)
slot := reflect.ValueOf(&x).Elem()
inside := slot.Elem()

fmt.Println(value.Kind(), value.CanSet())   // int32 false
fmt.Println(slot.Kind(), slot.CanSet())     // interface true
fmt.Println(inside.Kind(), inside.CanSet()) // int32 false

slot.Set(reflect.ValueOf("hello"))
fmt.Printf("%T %v\n", x, x) // string hello
