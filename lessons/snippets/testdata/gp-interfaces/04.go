// Фрагмент 4: book/chapters/gp-04-interfaces.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
type boxesByValue []boxint.BoxInt

func (s boxesByValue) Len() int { return len(s) }
func (s boxesByValue) Less(i, j int) bool {
    return s[i].Value < s[j].Value
}
func (s boxesByValue) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

// import "sort"
// sort.Sort(boxesByValue(values))
