// Фрагмент 5: book/chapters/07-interfaces.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type Worker interface { Work() }
type Job struct{}
func (*Job) Work() {}

var _ Worker = (*Job)(nil) // проверка реализации на этапе компиляции
// var _ Worker = Job{}   // у Job нет метода с pointer receiver

func nilExample() bool {
    var p *Job = nil
    var w Worker = p
    return w == nil // false: динамический тип *Job присутствует
}
