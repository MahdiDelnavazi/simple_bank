package util

//constants for all currency we support
const (
	USD = "USD"
	CAD = "CAD"
	IR  = "IR"
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case CAD, USD, IR:
		return true
	}
	return false

}
