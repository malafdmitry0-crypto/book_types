package aliases

type UserID = int64 // тот же тип
type OrderID int64  // отдельный тип

func aliasExample() {
	var n int64 = 42
	var user UserID = n // конверсия не нужна
	order := OrderID(n)
	_, _ = user, order
}
