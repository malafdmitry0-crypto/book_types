// Фрагмент 5: book/chapters/20-cpp-memory.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type Counter struct { N int }

func (c *Counter) Inc() { c.N++ }

type Incrementer interface { Inc() }
