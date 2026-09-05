package invalid
type Stringer interface {String()string}
var _ Stringer=42
