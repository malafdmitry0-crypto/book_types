package invalid
import("gotypes/fraction")
type Addable interface {Add(Addable)(Addable,error)}
var _ Addable=fraction.Fraction{}
