package typeinterfaces

import (
	"fmt"
	"io"
)

// import "io"
func inspect(rw io.ReadWriter) {
	var reader io.Reader = rw        // все нужные методы гарантированы статически
	writer, ok := reader.(io.Writer) // проверка динамического значения
	_, _ = writer, ok
}

// import "fmt"
func describe(x any) string {
	switch v := x.(type) {
	case nil:
		return "nil interface"
	case int:
		return fmt.Sprintf("int: %d", v)
	case string:
		return "string: " + v
	case float32, float64:
		return fmt.Sprintf("float: %v", v) // здесь v имеет тип any
	default:
		return fmt.Sprintf("other: %T", v)
	}
}

type Worker interface{ Work() }
type Job struct{}

func (*Job) Work() {}

var _ Worker = (*Job)(nil) // проверка реализации на этапе компиляции
// var _ Worker = Job{}   // у Job нет метода с pointer receiver

func nilExample() bool {
	var p *Job = nil
	var w Worker = p
	return w == nil // false: динамический тип *Job присутствует
}
