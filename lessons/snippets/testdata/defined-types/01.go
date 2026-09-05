// Фрагмент 1: book/chapters/06-defined-types.md
// Контекст и запускаемые сценарии: lessons/README.md и book/code-map.json.
// Это точная выдержка, не самостоятельная единица компиляции.
// import "fmt"
type UserID int64
func (id UserID) String() string { return fmt.Sprintf("user:%d", id) }

type OrderID UserID

func exampleIDs() {
    user := UserID(42)
    order := OrderID(user)
    _ = user.String()
    // _ = order.String() // у OrderID нет этого метода
    _ = order
}
