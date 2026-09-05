// Фрагмент 1: book/chapters/14-unsafe.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
// import "unsafe"
bits := uint32(0x3f800000)
number := *(*float32)(unsafe.Pointer(&bits)) // 1 на соответствующем представлении
