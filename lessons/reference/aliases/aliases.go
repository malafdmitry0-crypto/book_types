package aliases

type UserID = int64 // тот же тип
type OrderID int64  // отдельный тип

func aliasExample() {
	var n int64 = 42
	var user UserID = n // конверсия не нужна
	order := OrderID(n)
	_, _ = user, order
}

// describe различает типы в type switch: алиас UserID совпадает
// с case int64, а определённый OrderID требует собственной ветки.
func describe(x any) string {
	switch x.(type) {
	case int64: // сюда же попадает UserID
		return "int64"
	case OrderID:
		return "OrderID"
	default:
		return "other"
	}
}
