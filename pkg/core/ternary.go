package core

func Ternary[T any](condition bool, firstValue, secondValue T) T {
	if condition {
		return firstValue
	} else {
		return secondValue
	}
}
