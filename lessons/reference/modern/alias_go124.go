//go:build go1.24

package modern

type Set[T comparable] = map[T]bool // Go 1.24+

func aliasSetExample() {
	raw := map[string]bool{"go": true}
	var names Set[string] = raw // идентичный тип
	_ = names
}
