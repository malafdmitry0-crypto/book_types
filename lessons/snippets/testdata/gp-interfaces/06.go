// Фрагмент 6: book/chapters/gp-04-interfaces.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type boxes []boxint.BoxInt

func (s boxes) Len() int { return len(s) }
func (s boxes) Compare(i, j int) (int, error) {
    return s[i].Compare(s[j])
}
func (s boxes) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

var _ algorithms.Sortable = boxes(nil)
