package interfaces_test

import (
	"fmt"
	"gotypes/boxint"
	"sort"
)

type boxesByValue []boxint.BoxInt

func (s boxesByValue) Len() int           { return len(s) }
func (s boxesByValue) Less(i, j int) bool { return s[i].Value < s[j].Value }
func (s boxesByValue) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

var _ sort.Interface = boxesByValue(nil)

func Example_historicalSort() {
	values := []boxint.BoxInt{{Value: 3}, {Value: 1}, {Value: 2}}
	sort.Sort(boxesByValue(values))
	fmt.Println(values)
	// Output: [1 2 3]
}
