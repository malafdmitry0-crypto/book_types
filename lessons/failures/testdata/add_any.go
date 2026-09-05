package invalid
import("gotypes/fraction")
type Adder interface {Add(interface{})(interface{},error)}
var _ Adder=fraction.Fraction{}
