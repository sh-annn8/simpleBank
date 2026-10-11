package util

const (
	USD = "USD"
	EUR = "EUR"
	CAD = "CAD"
)

// 检查 currency 是否受支持，return true or false
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, CAD:
		return true
	}
	return false
}
